package message

import (
	"context"
	"fmt"
	"log"

	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

func StartConsumer(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		return
	}
	messageCh := wxbot.StartReceiveMessage(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case message := <-messageCh:
			for _, h := range handlers {
				if h.Support(message) {
					err := h.Handle(ctx, message)
					if err != nil {
						_ = wxbot.SendMsg(ctx, fmt.Sprintf("处理消息失败: %s", err.Error()), message.Chat)
						log.Printf("handle msg err:%v", err)
					}
				}
			}
		}
	}

}
