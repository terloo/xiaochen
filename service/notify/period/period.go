package period

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/thirdparty/wxbot"

	"github.com/terloo/xiaochen/util"
)

func StartPeriodNotifier(ctx context.Context) {
	printfLogger := cron.VerbosePrintfLogger(log.New(log.Writer(), "[period_notifier]  ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(printfLogger),
		cron.WithChain(cron.Recover(printfLogger)),
	)

	c.AddFunc("0 0 */2 * * *", func() {
		wxbot.KeepAlive(ctx)
	})

	weatherNotifier := &WeatherNotifier{}
	c.AddFunc("0 0 7 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		weatherNotifier.Notify(_ctx, family.FamilyChatroomWxid, family.MomWxid)
	})

	weatherHourlyNotifier := &WeatherHourlyNotifier{}
	c.AddFunc("0 0 7-23 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		weatherHourlyNotifier.Notify(_ctx, family.FamilyChatroomWxid, family.MomWxid)
	})

	birthdayNotifier := &BirthdayNotifier{
		&util.RealClock{},
		family.Families,
	}
	c.AddFunc("0 0 11 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		birthdayNotifier.Notify(_ctx, family.FamilyChatroomWxid)
	})

	tickerNotifier := &TickerNotifier{
		Tickers: []string{"600959"},
	}
	c.AddFunc("0 0 16 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		tickerNotifier.Notify(_ctx, family.MomWxid)
	})

	calendarNotifier := &CalendarNotifier{&util.RealClock{}}
	c.AddFunc("0 0 8 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		calendarNotifier.Notify(_ctx, family.FamilyChatroomWxid)
	})

	festivalNotifier := &FestivalNotifier{&util.RealClock{}}
	c.AddFunc("10 0 8 * * *", func() {
		_ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		defer cancelFunc()
		festivalNotifier.Notify(_ctx, family.FamilyChatroomWxid)
	})

	c.Start()
	<-ctx.Done()
	stop := c.Stop()
	<-stop.Done()
}
