package wxbot

import (
	"context"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/thirdparty/gpt"
)

func KeepAlive(ctx context.Context) {
	_ = SendMsg(ctx, family.TestChatroomWxid, "1")
}

func ResponseWithGPT(ctx context.Context, wxid string, message string) {
	s, err := gpt.Completion(ctx, message)
	respMessage := s
	if err != nil {
		respMessage = err.Error()
	}
	_ = SendMsg(ctx, wxid, respMessage)
}
