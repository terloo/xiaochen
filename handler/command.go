package handler

import (
	"context"
	"log"

	"github.com/terloo/xiaochen/notify/period"
)

type CommandHandler struct {
	CommonHandler
}

func (c *CommandHandler) GetHandlerName() string {
	return "CommandHandler"
}

func (c *CommandHandler) Support(msg FormattedMessage) bool {
	return c.TakeCare(msg) && msg.Command
}

func (c *CommandHandler) Handle(ctx context.Context, msg FormattedMessage) error {
	log.Printf("command: %s", msg.Content)
	weatherNotifier := period.WeatherNotifier{}
	weatherNotifier.Notify(ctx, msg.Chat)
	return nil
}

var _ MessageHandler = (*CommandHandler)(nil)
