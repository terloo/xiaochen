package message

import (
	"github.com/terloo/xiaochen/thirdparty/wxbot"
	"github.com/terloo/xiaochen/util"
)

type CommonHandler struct {
	CareSender []string
}

func (c CommonHandler) TakeCare(message wxbot.FormattedMessage) bool {
	return util.Contains(c.CareSender, message.Chat)
}
