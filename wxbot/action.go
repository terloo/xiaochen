package wxbot

import (
	"context"
	"fmt"
	"log"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/thirdparty/gpt"
)

func KeepAlive(ctx context.Context) {
	_ = SendMsg(ctx, "ping", family.TestChatroomWxid)
}

func ResponseWithGPT(ctx context.Context, wxid string, message string) {
	log.Printf("gpt completion message: %s\n", message)
	s, err := gpt.Completion(ctx, message)
	respMessage := s
	if err != nil {
		respMessage = fmt.Sprintf("sorry，出错了：%v", err)
	}
	_ = SendMsg(ctx, respMessage, wxid)
}
