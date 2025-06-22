package message

import (
	"context"
	"fmt"
	"log"

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
你可以不用带姓直接称呼名字，如果是未知成员，你可以称呼其为"家人"。

家庭成员会在一个群聊中进行对话，所以对话前面会加上说话人的姓名(注意你回复时不用添加自己名字)，你需要根据上下文进行回复，例如：
[家庭成员A]: 晴天这首歌是林俊杰唱的吧
[家庭成员B]: 不对，是周杰伦唱的
[家庭成员B]: @xiaochen 是谁唱的你知道吗？
`

var selfWxid string

func init() {
	wxid, err := wxbot.GetWxid(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	selfWxid = wxid
}

type GPTHandler struct {
	CommonHandler
	sessionIds     map[string]string
	sessionManager session.Manager
	mcpClient      *mcp.ClientManager
}

func NewGPTHandler(c CommonHandler) *GPTHandler {
	clientManager := mcp.NewClientManger()
	err := clientManager.InitializeAll(context.TODO())
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
	senderMessage := fmt.Sprintf("%s: %s", senderName, msg.Content)
	log.Printf("gpt completion senderMessage: %s\n", senderMessage)

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
	err = manager.AddUserRoleContent(ctx, senderMessage)
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
		result, err := c.mcpClient.CallTool(ctx, s.Message.ToolCalls[0].Function.Name, s.Message.ToolCalls[0].Function.Arguments)
		if err != nil {
			return err
		}
		content := result.Content[0].(gomcp.TextContent)
		err = manager.AddToolRoleContent(ctx, content.Text, s.Message.ToolCalls[0].ID)
		if err != nil {
			return err
		}

	}

	_ = wxbot.SendMsg(ctx, s.Message.Content, msg.Chat)

	return nil
}

var _ Handler = (*GPTHandler)(nil)
