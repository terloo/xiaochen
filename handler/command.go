package handler

import (
	"context"
	"log"
	"strings"

	"github.com/terloo/xiaochen/notify/period"
	"github.com/terloo/xiaochen/thirdparty/juhe"
	"github.com/terloo/xiaochen/wxbot"
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
	log.Printf("command: %s", msg.CommandName)
	switch msg.CommandName {
	case "天气":
		weatherNotifier := period.WeatherNotifier{}
		weatherNotifier.Notify(ctx, msg.Chat)
	case "解梦":
		zhouGongResult, err := juhe.GetZhouGong(ctx, strings.Join(msg.CommandArgs, " "))
		if err != nil {
			wxbot.SendMsg(ctx, err.Error(), msg.Chat)
			return err
		}
		respMsg := strings.Join(zhouGongResult.List, "\n")
		_ = wxbot.SendMsg(ctx, respMsg, msg.Chat)
	default:
		_ = wxbot.SendMsg(ctx, "未知命令", msg.Chat)
	}

	return nil
}

var _ MessageHandler = (*CommandHandler)(nil)
