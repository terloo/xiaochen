package session

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"github.com/terloo/xiaochen/session/cache"
)

// ChatContextManager 代表一次对话的上下文
type ChatContextManager struct {
	// 对话发起源
	Origin Origin
	// 消息发送人
	Sender string
	// 消息接收人
	Receiver string
	// 消息缓存key顺序排列
	messageKeys []string
	// 当前消息编号
	currentMessageCount int
	// 消息内容缓存
	cache cache.MessageCache

	mux sync.Mutex
}

func NewChat(origin Origin, sender string, receiver string) *ChatContextManager {
	localCache := cache.NewLocalCache(128*1024*1024, 6*3600)
	return &ChatContextManager{
		Origin:      origin,
		Sender:      sender,
		Receiver:    receiver,
		messageKeys: make([]string, 0),
		cache:       localCache,
	}
}

func (c *ChatContextManager) AddDeveloperRoleContent(ctx context.Context, content string) error {
	return c.addRoleContent(ctx, content, openai.ChatMessageRoleSystem)
}

func (c *ChatContextManager) AddUserRoleContent(ctx context.Context, sender string, content string) error {
	key := c.generateMessageCacheKey()
	message := openai.ChatCompletionMessage{
		Name:    sender,
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}
	err = c.cache.SetValue(ctx, key, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChatContextManager) AddAssistantRoleContent(ctx context.Context, content string, toolCalls []openai.ToolCall) error {
	key := c.generateMessageCacheKey()
	message := openai.ChatCompletionMessage{
		Role:      openai.ChatMessageRoleAssistant,
		Content:   content,
		ToolCalls: toolCalls,
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}
	err = c.cache.SetValue(ctx, key, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChatContextManager) AddToolRoleContent(ctx context.Context, content string, toolCallId string) error {
	key := c.generateMessageCacheKey()
	message := openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    content,
		ToolCallID: toolCallId,
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}
	err = c.cache.SetValue(ctx, key, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChatContextManager) addRoleContent(ctx context.Context, content string, role string) error {
	key := c.generateMessageCacheKey()
	message := openai.ChatCompletionMessage{
		Role:    role,
		Content: content,
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}
	err = c.cache.SetValue(ctx, key, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChatContextManager) GetAllRoleContent(ctx context.Context) ([]openai.ChatCompletionMessage, error) {
	var result []openai.ChatCompletionMessage
	for i := 0; i < c.currentMessageCount; i++ {
		messageContent, err := c.cache.GetValue(ctx, fmt.Sprintf("message:%d", i))
		if err != nil {
			log.Printf("%+v\n", errors.Wrapf(err, "get message cache error"))
			continue
		}
		if messageContent == nil {
			continue
		}
		var message openai.ChatCompletionMessage
		err = json.Unmarshal(messageContent, &message)
		if err != nil {
			log.Printf("%+v\n", errors.Wrapf(err, "unmarshal message cache error"))
			continue
		}
		result = append(result, message)
	}
	return result, nil
}

func (c *ChatContextManager) generateMessageCacheKey() string {
	c.mux.Lock()
	defer c.mux.Unlock()

	defer func() { c.currentMessageCount += 1 }()
	return fmt.Sprintf("message:%d", c.currentMessageCount)
}

var _ ContextManger = (*ChatContextManager)(nil)
