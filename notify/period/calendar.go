package period

import (
	"context"
	"fmt"

	"github.com/terloo/xiaochen/almanac"
	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/util"
	"github.com/terloo/xiaochen/wxbot"
)

type CalendarNotifier struct {
	util.Clock
}

var _ notify.Notifier = (*CalendarNotifier)(nil)

func (c *CalendarNotifier) Notify(ctx context.Context, notified ...string) {
	now := c.Clock.Now()
	almDay := almanac.NewDay(almanac.NewTime(now.Year(), int(now.Month()), now.Day(), 0, 0, 0))
	lunar := almDay.Lunar
	msg := fmt.Sprintf("今天是公历%s星期%s，农历%s%s年%s月%s", now.Format(util.DateLayout), util.IntToWeekday[almDay.Week], lunar.Year2, almDay.GetChineseZodiacName(), lunar.MonthName, lunar.DayName)
	_ = wxbot.SendMsg(ctx, msg, notified...)
}
