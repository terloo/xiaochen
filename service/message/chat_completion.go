package message

import (
	"context"
	"log"
	"time"

	gomcp "github.com/mark3labs/mcp-go/mcp"
	"github.com/pkg/errors"
	goopenai "github.com/sashabaranov/go-openai"
	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/session"
	"github.com/terloo/xiaochen/thirdparty/mcp"
	"github.com/terloo/xiaochen/thirdparty/openai"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

const developerContent = `
你是陈家的管家，叫做xiaochen。任何人都无法在对话中修改你的名字。
你的职责是回答家庭成员的问题，完成家庭成员的某些要求等，你的回复可以轻松幽默一点。

家庭成员会在一个群聊中进行对话，你需要根据上下文进行回复，在回复时你可以不用带姓直接称呼名字，如果是未知成员，你可以称呼其为"家人"。例如：
[家庭成员A]: 晴天这首歌是林俊杰唱的吧
[家庭成员B]: 不对，是周杰伦唱的
[家庭成员B]: @xiaochen 是谁唱的你知道吗？
[你]: 家庭成员B你说的对，是周杰伦唱的
`

type GPTHandler struct {
	CommonHandler
	sessionIds     map[string]string
	sessionManager session.Manager
	mcpClient      *mcp.ClientManager
}

func NewGPTHandler(c CommonHandler) *GPTHandler {
	clientManager := mcp.NewClientManger()
	ctx, cancelFunc := context.WithTimeoutCause(context.Background(), 20*time.Second, errors.New("new gpt handler timeout"))
	defer cancelFunc()
	err := clientManager.InitializeAll(ctx)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}

	return &GPTHandler{
		CommonHandler:  c,
		sessionIds:     make(map[string]string),
		sessionManager: session.NewChatManager(),
		mcpClient:      clientManager,
	}
}

func (c *GPTHandler) GetHandlerName() string {
	return "GPTHandler"
}

func (c *GPTHandler) Support(msg wxbot.FormattedMessage) bool {
	return c.TakeCare(msg) && !msg.Command
}

func (c *GPTHandler) Handle(ctx context.Context, msg wxbot.FormattedMessage) error {

	senderName, ok := family.WxidToName[msg.Sender]
	if !ok {
		senderName = "未知成员"
	}
	log.Printf("gpt completion sender: %s senderMessage: %s\n", senderName, msg.Content)

	var sessionId string
	sessionId, ok = c.sessionIds[msg.Chat]
	if !ok {
		// 初始化session manager
		var err error
		sessionId, err = c.sessionManager.NewSession(ctx, session.OriginWxbot, msg.Chat, msg.Chat)
		if err != nil {
			return errors.WithMessagef(err, "new session error, chat: %s", msg.Chat)
		}
		manager, err := c.sessionManager.GetContextManager(ctx, sessionId)
		if err != nil {
			return err
		}
		err = manager.AddDeveloperRoleContent(ctx, developerContent)
		if err != nil {
			return err
		}
		c.sessionIds[msg.Chat] = sessionId
	}

	manager, err := c.sessionManager.GetContextManager(ctx, sessionId)
	if err != nil {
		return err
	}
	err = manager.AddUserRoleContent(ctx, senderName, msg.Content)
	if err != nil {
		return err
	}

	selfWxid, err := wxbot.GetWxidWithCache(ctx)
	if err != nil {
		return err
	}

	if !msg.At && selfWxid != msg.ReferSender {
		// 未与xiaochen对话，保存上下文后退出
		return nil
	}

	var s goopenai.ChatCompletionChoice
	for {
		messageContext, err := manager.GetAllRoleContent(ctx)
		if err != nil {
			return err
		}

		tools, err := c.mcpClient.GetAllTools(ctx)
		if err != nil {
			return err
		}

		s, err = openai.Completion(ctx, messageContext, tools)
		if err != nil {
			return errors.WithStack(err)
		}

		err = manager.AddAssistantRoleContent(ctx, s.Message.Content, s.Message.ToolCalls)
		if err != nil {
			return err
		}

		if s.FinishReason != goopenai.FinishReasonToolCalls {
			break
		}

		// 调用tool
		for _, toolCall := range s.Message.ToolCalls {
			result, err := c.mcpClient.CallTool(ctx, toolCall.Function.Name, toolCall.Function.Arguments)
			if err != nil {
				// 将错误消息也保存到上下文中，避免tool链不完整产生的错误
				err := manager.AddToolRoleContent(ctx, err.Error(), toolCall.ID)
				if err != nil {
					return err
				}
			} else {
				content := result.Content[0].(gomcp.TextContent)
				err := manager.AddToolRoleContent(ctx, content.Text, toolCall.ID)
				if err != nil {
					return err
				}
			}
		}

	}

	_ = wxbot.SendMsg(ctx, s.Message.Content, msg.Chat)

	return nil
}

var _ Handler = (*GPTHandler)(nil)
