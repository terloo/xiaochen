package message

import (
	"context"
	"log"

	"github.com/terloo/xiaochen/wxbot"
)

func StartConsumer(ctx context.Context) {
	for message := range wxbot.StartReceiveMessage(ctx) {
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
