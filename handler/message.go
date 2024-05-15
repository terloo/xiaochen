package handler

import (
	"context"
	"log"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/wxbot"
)

type MessageHandler interface {
	GetHandlerName() string
	Support(msg FormattedMessage) bool
	Handle(ctx context.Context, msg FormattedMessage) error
}

var handlers []MessageHandler

func HandleMessage(ctx context.Context, msg wxbot.WxGeneralMsg) {
	for _, msgData := range msg.Data {
		if msgData.IsSender == "1" {
			// 暂不处理自己发送的消息
			continue
		}
		message, err := FormatMessage(msgData)
		if err != nil {
			continue
		}
		for _, h := range handlers {
			if h.Support(message) {
				err := h.Handle(ctx, message)
				if err != nil {
					log.Printf("handle msg err:%v", err)
				}
			}
		}
	}
}

func init() {
	handlers = append(handlers, &CommandHandler{
		CommonHandler: CommonHandler{
			CareSender: []string{family.FamilyChatroomWxid, family.TestChatroomWxid},
		},
	})
	handlers = append(handlers, &GPTHandler{
		CommonHandler: CommonHandler{
			CareSender: []string{family.FamilyChatroomWxid, family.TestChatroomWxid},
		},
	})
}
