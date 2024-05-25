package wxbot

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
)

func StartReceiveMessage(ctx context.Context) <-chan FormattedMessage {
	url := "ws://tx:8080/ws/generalMsg"
	ws, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resultChan := make(chan FormattedMessage, 10)

	isClosed := false
	go func() {
		<-ctx.Done()
		log.Println("close websocket connection...")
		_ = ws.Close()
		isClosed = true
		close(resultChan)
	}()

	go func() {
		for {
			if isClosed {
				break
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Println("websocket panic: ", r)
						// panic后重新建立连接
						ws, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
						if err != nil {
							log.Fatal(err)
						}
						log.Println("reconnect websocket")
					}
				}()

				message, err := ReadMessage(ws)
				if err != nil {
					log.Println(err)
					return
				}
				for _, data := range message.Data {
					formattedMessage, err := FormatMessage(data)
					if err != nil {
						log.Println(err)
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

func ReadMessage(ws *websocket.Conn) (*WxGeneralMsg, error) {
	receiveMsg := &WxGeneralMsg{}
	err := ws.ReadJSON(receiveMsg)
	if err != nil {
		return nil, err
	}
	log.Println("receive message: ", receiveMsg)
	return receiveMsg, nil
}
