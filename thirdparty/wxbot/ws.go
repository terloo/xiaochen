package wxbot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"github.com/terloo/xiaochen/config"
)

func StartReceiveMessage(ctx context.Context) <-chan FormattedMessage {
	wxBotHost := config.NewLoader("main.wxBotHost").Get()
	url := fmt.Sprintf("ws://%s/ws/generalMsg", wxBotHost)
	ws, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resultChan := make(chan FormattedMessage, 10)

	go func() {
		<-ctx.Done()
		log.Println("close websocket connection...")
		_ = ws.Close()
		close(resultChan)
	}()

	go func() {
		for {
			if err := ctx.Err(); err != nil {
				log.Println("stop receive message")
				return
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Println("websocket panic: ", r)
						// panic后重新建立连接
						ws, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
						if err != nil {
							log.Fatalf("reconnect websocket error: %+v", err)
						}
						log.Println("reconnect websocket")
					}
				}()

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
						return
					}
					if formattedMessage.Self && !formattedMessage.At {
						// 暂不处理自己发送的消息
						return
					}
					resultChan <- formattedMessage
				}
				return
			}()
		}
	}()

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
