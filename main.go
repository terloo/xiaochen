package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/terloo/xiaochen/notify/period"
	"github.com/terloo/xiaochen/ws"
	"github.com/terloo/xiaochen/wxbot"
)

func init() {
	initLogger()
}

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// 校验登录状态
	wxid, err := wxbot.GetWxid(ctx)
	if err != nil || wxid == "" {
		log.Fatal(err)
	}

	go func() {
		wg.Add(1)
		period.StartPeriodNotifier(ctx)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		ws.StartReceiveMessage(ctx)
		wg.Done()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	cancel()
	log.Println("Waiting for task exit...")
	wg.Wait()
}

func initLogger() {
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds)
	log.SetPrefix("[xiaochen] ")
}
