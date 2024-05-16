package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"log"

	"github.com/terloo/xiaochen/handler"
	"github.com/terloo/xiaochen/wxbot"
)

func StartReceiveMessage(ctx context.Context) {
	url := "ws://tx:8080/ws/generalMsg"
	ws, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	isClosed := false
	go func() {
		<-ctx.Done()
		log.Println("close websocket connection...")
		_ = ws.Close()
		isClosed = true
	}()

	defer func() {
		if r := recover(); r != nil {
			log.Println("接收消息panic: ", r)
		}
	}()

	for {
		if isClosed {
			break
		}
		message, err := ReadMessage(ws)
		if err != nil {
			// TODO 判断错误类型，如果连接已关闭则需要重新连接
			continue
		}
		handler.HandleMessage(ctx, *message)
	}

}

func ReadMessage(ws *websocket.Conn) (*wxbot.WxGeneralMsg, error) {
	receiveMsg := &wxbot.WxGeneralMsg{}
	err := ws.ReadJSON(receiveMsg)
	if err != nil {
		return nil, err
	}
	log.Println("receive: ", receiveMsg)
	return receiveMsg, nil
}
