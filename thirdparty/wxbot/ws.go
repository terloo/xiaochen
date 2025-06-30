package wxbot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"

	"github.com/terloo/xiaochen/config"
)

func StartReceiveMessage(ctx context.Context) <-chan FormattedMessage {
	wxBotHost := config.NewLoader("main.wxBotHost").Get()
	url := fmt.Sprintf("ws://%s/ws/generalMsg", wxBotHost)
	retryDuration := 5 * time.Second
	resultChan := make(chan FormattedMessage, 10)

	var ws *websocket.Conn
	go func(ctx context.Context) {
		// 注册退出逻辑
		<-ctx.Done()
		log.Println("close websocket connection...")
		if ws != nil {
			_ = ws.Close()
		}
		close(resultChan)
	}(ctx)

	go func(ctx context.Context) {
		for {
			if err := ctx.Err(); err != nil {
				log.Printf("stop receive message: %+v\n", err)
				return
			}

			var err error
			ws, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
			if err != nil {
				log.Printf("dial to wxbot, try to reconnect after %s, error: %+v\n", retryDuration, err)
				<-time.After(retryDuration)
				continue
			}
			log.Println("dial to wxbot success")

			func(ctx context.Context) {
				defer func() {
					if r := recover(); r != nil {
						log.Println("websocket panic: ", r)
					}
				}()

				log.Println("start receive message")
				for {
					message, err := ReadMessage(ctx, ws)
					if err != nil {
						log.Printf("read message error: %+v\n", err)
						return
					}
					for _, data := range message.Data {
						if err := ctx.Err(); err != nil {
							return
						}
						formattedMessage, err := FormatMessage(data)
						if err != nil {
							log.Printf("format message error: %+v\n", err)
							continue
						}
						if formattedMessage.Self && !formattedMessage.At {
							// 暂不处理自己发送的消息
							continue
						}
						resultChan <- formattedMessage
					}
				}
			}(ctx)
		}
	}(ctx)

	return resultChan
}

func ReadMessage(ctx context.Context, ws *websocket.Conn) (*WxGeneralMsg, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	receiveMsg := &WxGeneralMsg{}
	err := ws.ReadJSON(receiveMsg)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(receiveMsg)
	if err != nil {
		return nil, err
	}
	log.Println("receive message: ", string(marshal))
	return receiveMsg, nil
}
