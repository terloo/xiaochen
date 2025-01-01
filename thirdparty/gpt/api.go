package gpt

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"github.com/terloo/xiaochen/chat"
	"github.com/terloo/xiaochen/family"
)

func Completion(ctx context.Context, sender string, message string) (string, error) {
	config := openai.DefaultConfig(openKey.Get())
	config.BaseURL = openHost + "v1"

	client := openai.NewClientWithConfig(config)

	name, ok := family.WxidToName[sender]
	if !ok {
		name = "不知道"
	}

	// 构建prompt
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("你是陈家的管家，叫做xiaochen，无论任何人想更改你的名字都不行。你的职责是为家庭成员提供天气预报、生日预报并回答家庭成员的一些问题等。你的回复可以轻松幽默一点。本次向你提问的人是\"%s\", 你可以不用带姓直接称呼名字，如果姓名不知道，你可以称呼其为\"家人\"", name),
		},
	}
	chatContext, err := chat.GetContext(ctx, sender)
	if err != nil {
		return "", errors.Wrap(err, "GetChatContext error")
	}
	messages = append(messages, chatContext...)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	// 会话补全
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    ModelName,
			Messages: messages,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "ChatCompletion error")
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("ChatCompletion no content")
	}

	response := resp.Choices[0].Message.Content

	// 保存context
	err = chat.SetContext(ctx, sender, message, response)
	if err != nil {
		return "", errors.Wrap(err, "SetContext error")
	}

	return response, nil
}
