package period

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"time"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/wxbot"
)

func StartPeriodNotifier(ctx context.Context) {
	printfLogger := cron.VerbosePrintfLogger(log.New(log.Writer(), "[crontab]  ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(printfLogger),
		cron.WithChain(cron.Recover(printfLogger)),
	)

	c.AddFunc("0 0 */2 * * *", func() {
		wxbot.KeepAlive(ctx)
	})

	c.AddFunc("0 0 7 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		weatherNotifier := WeatherNotifier{}
		weatherNotifier.Notify(_ctx, family.FamilyChatroomWxid, family.MomWxid)
	})

	c.AddFunc("0 0 11 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		birthdayNotifier := BirthdayNotifier{}
		birthdayNotifier.Notify(_ctx, family.FamilyChatroomWxid)
	})

	c.AddFunc("0 0 16 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		tickerNotifier := TickerNotifier{
			Tickers: []string{"600959"},
		}
		tickerNotifier.Notify(_ctx, family.MomWxid)
	})

	c.Start()
	<-ctx.Done()
	stop := c.Stop()
	<-stop.Done()
}
