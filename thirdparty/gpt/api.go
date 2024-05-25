package gpt

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

func Completion(ctx context.Context, message string) (string, error) {
	config := openai.DefaultConfig(openKey.Get())
	config.BaseURL = openHost + "v1"

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: ModelName,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
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
