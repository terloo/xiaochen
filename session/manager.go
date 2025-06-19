package session

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sync"
)

// LocalManager 用于管理所有的对话
type LocalManager struct {
	chats map[string]*ChatContextManager
	mux   sync.Mutex
}

func NewChatManager() *LocalManager {
	return &LocalManager{chats: make(map[string]*ChatContextManager)}
}

func (m *LocalManager) NewSession(ctx context.Context, origin Origin, sender string, receiver string) (string, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	s := fmt.Sprintf("%s:%s:%s", string(origin), sender, receiver)
	h := sha1.New()
	h.Write([]byte(s))
	sessionId := hex.EncodeToString(h.Sum(nil))
	m.chats[sessionId] = NewChat(origin, sender, receiver)
	return sessionId, nil
}

func (m *LocalManager) GetContextManager(ctx context.Context, sessionId string) (ContextManger, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	chat, ok := m.chats[sessionId]
	if !ok {
		return nil, ErrSessionIdNotFound
	}
	return chat, nil
}

var _ Manager = (*LocalManager)(nil)
