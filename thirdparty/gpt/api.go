package gpt

import (
	"context"
	"errors"
	"log"

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
		log.Printf("ChatCompletion error: %v\n", err)
		return "", errors.New("sorry, 出错了！")
	}

	return resp.Choices[0].Message.Content, nil
}
