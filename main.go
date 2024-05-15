package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/ws"
	"github.com/terloo/xiaochen/wxbot"
)

func init() {
	initLogger()
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 校验登录状态
	wxid, err := wxbot.GetWxid(ctx)
	if err != nil || wxid == "" {
		log.Fatal(err)
	}

	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(cron.VerbosePrintfLogger(log.Default())),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)
	defer func() {
		stop := c.Stop()
		<-stop.Done()
	}()

	c.AddFunc("0 0 */2 * * *", func() {
		wxbot.KeepAlive(ctx)
	})
	c.AddFunc("0 0 7 * * *", func() {
		wxbot.ReportWeather(ctx, family.MomWxid)
		wxbot.ReportWeather(ctx, family.FamilyChatroomWxid)
	})
	c.AddFunc("0 0 16 * * *", func() {
		wxbot.ReportTicker(ctx, family.MomWxid, []string{"600959"})
	})
	c.AddFunc("0 0 11 * * *", func() {
		for _, notifier := range notify.Notifiers {
			notifier.Notify(ctx, family.FamilyChatroomWxid)
		}
	})

	go c.Start()
	go ws.StartReceiveMessage(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

func initLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[xiaochen]")
}
