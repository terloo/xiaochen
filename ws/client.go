package ws

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/terloo/xiaochen/handler"
	"github.com/terloo/xiaochen/wxbot"
)

func StartReceiveMessage(ctx context.Context) {
	url := "ws://tx:8080/ws/generalMsg"
	ws, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		defer func() {
			if r := recover(); r != nil {
				log.Println("接收消息panic")
			}
		}()
		message := ReadMessage(ws)
		handler.HandleMessage(ctx, message)
	}
}

func ReadMessage(ws *websocket.Conn) wxbot.WxGeneralMsg {
	_, data, err := ws.ReadMessage()
	if err != nil {
		log.Panicln(err)
	}
	receiveMsg := &wxbot.WxGeneralMsg{}
	err = json.Unmarshal(data, receiveMsg)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("receive: ", string(data))
	return *receiveMsg
}
