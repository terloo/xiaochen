package message

import (
	"github.com/terloo/xiaochen/util"
	"github.com/terloo/xiaochen/wxbot"
)

type CommonHandler struct {
	CareSender []string
}

func (c CommonHandler) TakeCare(message wxbot.FormattedMessage) bool {
	return util.Contains(c.CareSender, message.Chat)
}
