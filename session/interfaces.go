package session

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

var ErrSessionIdNotFound = errors.New("session id not found")

type Origin string

const OriginWxbot Origin = "ChatOriginWxbot"

type ContextManger interface {
	AddDeveloperRoleContent(ctx context.Context, content string) error
	AddUserRoleContent(ctx context.Context, content string) error
	AddAssistantRoleContent(ctx context.Context, content string) error
	GetAllRoleContent(ctx context.Context) ([]openai.ChatCompletionMessage, error)
}

type Manager interface {
	// NewSession 初始化一个session，返回sessionId
	NewSession(ctx context.Context, origin Origin, sender string, receiver string) (sessionId string, err error)
	GetContextManager(ctx context.Context, sessionId string) (ContextManger, error)
}
