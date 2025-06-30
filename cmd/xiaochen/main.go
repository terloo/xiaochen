package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/terloo/xiaochen/config"
	"github.com/terloo/xiaochen/service/message"
	"github.com/terloo/xiaochen/service/music"
	"github.com/terloo/xiaochen/service/notify/period"
)

func init() {
	initLogger()
}

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		wg.Add(1)
		period.StartPeriodNotifier(ctx)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		message.StartConsumer(ctx)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		music.StartPeriodNetease(ctx)
		wg.Done()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	cancel()
	log.Println("Waiting for task exit...")
	wg.Wait()
	log.Println("Bye!")
}

func initLogger() {
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds)
	log.SetPrefix("[xiaochen] ")
}
