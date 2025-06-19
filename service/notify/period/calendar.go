package period

import (
	"context"
	"fmt"
	"strings"

	"github.com/terloo/almanac"
	"github.com/terloo/xiaochen/service/notify"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
	"github.com/terloo/xiaochen/util"
)

type CalendarNotifier struct {
	util.Clock
}

var _ notify.Notifier = (*CalendarNotifier)(nil)

func (c *CalendarNotifier) Notify(ctx context.Context, notified ...string) {
	now := c.Clock.Now()
	almDay := almanac.NewDay(almanac.NewTime(now.Year(), int(now.Month()), now.Day(), 0, 0, 0))
	lunar := almDay.Lunar

	solarMsg := fmt.Sprintf("公历%s星期%s", now.Format(util.DateLayout), util.IntToWeekday[almDay.Week])
	solarSpecialDay := ""
	for _, f := range almDay.Events.Festival {
		if f == "国庆节假日" {
			// 特殊处理一下国庆节假日
			continue
		}
		solarSpecialDay += f + "，"
	}
	for _, f := range almDay.Events.Important {
		solarSpecialDay += f + "，"
	}
	solarSpecialDay = strings.TrimSuffix(solarSpecialDay, "，")
	if len(solarSpecialDay) != 0 {
		solarMsg += fmt.Sprintf("（%s）", solarSpecialDay)
	}

	lunarMsg := fmt.Sprintf("农历%s%s年%s%s月%s",
		lunar.Year2,
		almDay.GetChineseZodiacName(),
		lunar.LeapStr,
		lunar.MonthName,
		lunar.DayName,
	)
	lunarSpecialDay := ""
	for _, f := range lunar.Events.Important {
		lunarSpecialDay += f + "，"
	}
	lunarSpecialDay = strings.TrimSuffix(lunarSpecialDay, "，")
	if len(lunarSpecialDay) != 0 {
		lunarMsg += fmt.Sprintf("（%s）", lunarSpecialDay)
	}

	msg := fmt.Sprintf("今天是%s，%s", solarMsg, lunarMsg)

	_ = wxbot.SendMsg(ctx, msg, notified...)
}
