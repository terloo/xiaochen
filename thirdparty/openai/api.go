package openai

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

func Completion(ctx context.Context, completionMessages []openai.ChatCompletionMessage, tools []openai.Tool) (openai.ChatCompletionChoice, error) {
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
			Tools:       tools,
		},
	)

	if err != nil {
		return openai.ChatCompletionChoice{}, errors.Wrap(err, "ChatCompletion error")
	}

	if len(resp.Choices) == 0 {
		return openai.ChatCompletionChoice{}, errors.New("ChatCompletion no content")
	}

	return resp.Choices[0], nil
}
