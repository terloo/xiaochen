package handler

import (
	"github.com/terloo/xiaochen/util"
)

type CommonHandler struct {
	CareSender []string
}

func (c CommonHandler) TakeCare(message FormattedMessage) bool {
	return util.Contains(c.CareSender, message.Chat)
}
