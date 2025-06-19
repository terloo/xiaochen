package message

import (
	"context"

	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/service/message/command"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

type Handler interface {
	GetHandlerName() string
	Support(msg wxbot.FormattedMessage) bool
	Handle(ctx context.Context, msg wxbot.FormattedMessage) error
}

var handlers []Handler

func init() {
	handlers = append(handlers, &CommandHandler{
		Handlers: []command.Handler{
			&command.Birthday{},
			&command.Weather{},
			&command.Music{},
		},
		CommonHandler: CommonHandler{
			CareSender: []string{family.FamilyChatroomWxid, family.TestChatroomWxid},
		},
	})

	handlers = append(handlers, NewGPTHandler(CommonHandler{
		CareSender: []string{family.FamilyChatroomWxid, family.TestChatroomWxid},
	}))
}
