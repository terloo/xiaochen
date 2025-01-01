package chat

import (
	"context"
	"errors"
	"log"
	"strings"
	"sync"

	"github.com/coocood/freecache"
	"github.com/sashabaranov/go-openai"
)

var cache *freecache.Cache

func init() {
	cache = freecache.NewCache(100 * 1024 * 1024)
}

var tmux sync.Mutex

func GetContext(ctx context.Context, wxid string) ([]openai.ChatCompletionMessage, error) {
	tmux.Lock()
	defer tmux.Unlock()

	value, err := cache.Get([]byte(wxid))
	if errors.Is(err, freecache.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	c := string(value)
	splitContext := strings.Split(c, "#*#")
	if len(splitContext) == 0 {
		return []openai.ChatCompletionMessage{}, nil
	}

	// 减少上下文大小
	for len(c) > 14_000 {
		splitContext = splitContext[1:]
		c = strings.Join(splitContext, "#*#")
	}
	var chatContext []openai.ChatCompletionMessage
	for _, turn := range splitContext {
		turnSplit := strings.Split(turn, "$*$")
		if len(turnSplit) != 2 {
			log.Printf("turnSplit length invalid, turn: %s", turn)
			continue
		}
		chatContext = append(chatContext, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: turnSplit[0],
		})
		chatContext = append(chatContext, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: turnSplit[1],
		})
	}
	return chatContext, nil
}

func SetContext(ctx context.Context, wxid string, userMsg string, aiMsg string) error {
	tmux.Lock()
	defer tmux.Unlock()

	value, err := cache.Get([]byte(wxid))
	if errors.Is(err, freecache.ErrNotFound) {
		value = []byte("")
	} else if err != nil {
		return err
	}
	c := string(value)
	c = c + "#*#" + userMsg + "$*$" + aiMsg
	cache.Set([]byte(wxid), []byte(c), 3600)
	return nil
}
