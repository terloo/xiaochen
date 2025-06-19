package session

import (
	"context"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestSessionManager(t *testing.T) {
	ctx := context.Background()
	var sessionManager Manager = NewChatManager()
	sessionId, err := sessionManager.NewSession(ctx, OriginWxbot, "zhangsan", "zhangsan")
	if err != nil {
		t.Fatal(err)
	}

	contextManager, err := sessionManager.GetContextManager(ctx, sessionId)
	if err != nil {
		t.Fatal(err)
	}

	exceptMessages := []*openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleDeveloper,
			Content: "你是一个ai小助手，名字叫xiaochen",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "你是谁",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "我叫xiaochen，是一个ai小助手",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "重复一下上面的回答",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "我叫xiaochen，是一个ai小助手",
		},
	}

	for _, exceptMessage := range exceptMessages {
		switch exceptMessage.Role {
		case openai.ChatMessageRoleDeveloper:
			err := contextManager.AddDeveloperRoleContent(ctx, exceptMessage.Content)
			if err != nil {
				t.Fatal(err)
			}
		case openai.ChatMessageRoleUser:
			err := contextManager.AddUserRoleContent(ctx, exceptMessage.Content)
			if err != nil {
				t.Fatal(err)
			}
		case openai.ChatMessageRoleAssistant:
			err := contextManager.AddAssistantRoleContent(ctx, exceptMessage.Content)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	roleContent, err := contextManager.GetAllRoleContent(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(exceptMessages) != len(roleContent) {
		t.Errorf("except len %d but got %d\n", len(exceptMessages), len(roleContent))
	}

	for i, content := range roleContent {
		exceptMessage := exceptMessages[i]
		if exceptMessage.Role != content.Role {
			t.Errorf("except role %s but got %s\n", exceptMessage.Role, content.Role)
		}
		if exceptMessage.Content != content.Content {
			t.Errorf("except content %s but got %s\n", exceptMessage.Content, content.Content)
		}
	}

}
