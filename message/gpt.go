package message

import (
	"context"
	"log"

	"github.com/terloo/xiaochen/wxbot"
)

var selfWxid string

func init() {
	wxid, err := wxbot.GetWxid(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	selfWxid = wxid
}

type GPTHandler struct {
	CommonHandler
}

func (c *GPTHandler) GetHandlerName() string {
	return "GPTHandler"
}

func (c *GPTHandler) Support(msg wxbot.FormattedMessage) bool {
	return c.TakeCare(msg) && (msg.At || selfWxid == msg.ReferSender) && !msg.Command
}

func (c *GPTHandler) Handle(ctx context.Context, msg wxbot.FormattedMessage) error {
	go wxbot.ResponseWithGPT(ctx, msg.Chat, msg.Content)
	return nil
}

var _ Handler = (*GPTHandler)(nil)
