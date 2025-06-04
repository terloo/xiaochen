package period

import (
	"context"
	"fmt"
	"strings"

	"github.com/terloo/almanac"
	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/util"
	"github.com/terloo/xiaochen/wxbot"
)

type FestivalNotifier struct {
	util.Clock
}

var _ notify.Notifier = (*FestivalNotifier)(nil)

func (c *FestivalNotifier) Notify(ctx context.Context, notified ...string) {
	now := c.Clock.Now()
	almDay := almanac.NewDay(almanac.NewTime(now.Year(), int(now.Month()), now.Day(), 0, 0, 0))
	msg := "@xiaochen "
	for _, f := range almDay.Events.Festival {
		msg += f + ","
	}
	for _, f := range almDay.Lunar.Events.Festival {
		msg += f + ","
	}
	if msg != "@xiaochen " {
		msg = strings.TrimSuffix(msg, ",")
		msg += fmt.Sprintf("到了，能讲一下它的来历和习俗吗？\n")
		_ = wxbot.SendMsg(ctx, msg, notified...)
	}
}
