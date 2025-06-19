package openai

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

func Completion(ctx context.Context, completionMessages []openai.ChatCompletionMessage) (string, error) {
	config := openai.DefaultConfig(openKey.Get())
	config.BaseURL = openHost + "v1"

	client := openai.NewClientWithConfig(config)
	// 会话补全
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       ModelName,
			Messages:    completionMessages,
			Temperature: 0.8,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "ChatCompletion error")
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("ChatCompletion no content")
	}

	return resp.Choices[0].Message.Content, nil
}
