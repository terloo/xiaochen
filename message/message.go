package message

import (
	"context"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/message/command"
	"github.com/terloo/xiaochen/wxbot"
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
			&command.Zhougong{},
			&command.Weather{},
		},
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
