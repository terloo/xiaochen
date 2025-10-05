package period

import (
	"context"
	"fmt"
	"log"

	"github.com/terloo/xiaochen/service/notify"
	"github.com/terloo/xiaochen/thirdparty/tsanghi"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

type TickerNotifier struct {
	Tickers []string
}

var _ notify.Notifier = (*TickerNotifier)(nil)

func (t *TickerNotifier) Notify(ctx context.Context, notified ...string) {
	wxMsg := "股票行情：\n\n"

	if len(t.Tickers) == 0 {
		return
	}
	for _, ticker := range t.Tickers {
		todayTickerData, err := tsanghi.GetTodayTicker(ctx, ticker)
		if err != nil {
			log.Printf("get today ticket error: %+v\n", err)
			return
		}
		wxMsg += fmt.Sprintf("\t代码：%s\n", ticker)
		wxMsg += fmt.Sprintf("\t开：%.2f\n", todayTickerData.Open)
		wxMsg += fmt.Sprintf("\t收：%.2f\n", todayTickerData.Close)
		wxMsg += fmt.Sprintf("\t高：%.2f\n", todayTickerData.High)
		wxMsg += fmt.Sprintf("\t低：%.2f\n", todayTickerData.Low)
		wxMsg += fmt.Sprintf("\t成交量：%d\n", int(todayTickerData.Volume))
		wxMsg += fmt.Sprintf("\n")
	}

	_ = wxbot.SendMsg(ctx, wxMsg, notified...)
}
