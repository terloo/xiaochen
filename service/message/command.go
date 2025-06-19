package message

import (
	"context"
	"log"

	"github.com/terloo/xiaochen/service/message/command"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

type CommandHandler struct {
	CommonHandler
	Handlers []command.Handler
}

func (c *CommandHandler) GetHandlerName() string {
	return "CommandHandler"
}

func (c *CommandHandler) Support(msg wxbot.FormattedMessage) bool {
	return c.TakeCare(msg) && msg.Command
}

func (c *CommandHandler) Handle(ctx context.Context, msg wxbot.FormattedMessage) error {
	log.Printf("handle command: %s", msg.CommandName)
	for _, handler := range c.Handlers {
		if handler.CommandName() == msg.CommandName {
			return handler.Exec(ctx, msg.Chat, msg.CommandArgs)
		}
	}
	return wxbot.SendMsg(ctx, "未知命令", msg.Chat)
}

var _ Handler = (*CommandHandler)(nil)
