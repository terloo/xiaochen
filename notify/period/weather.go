package period

import (
	"context"
	"fmt"
	"log"

	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/thirdparty/gaode"
	"github.com/terloo/xiaochen/wxbot"
)

type WeatherNotifier struct {
}

var _ notify.Notifier = (*WeatherNotifier)(nil)

func (w *WeatherNotifier) Notify(ctx context.Context, notified ...string) {
	citys := []string{"江油市", "成都市"}
	weathers := make([]*gaode.Weather, len(citys))
	for i, city := range citys {
		weather, err := gaode.GetWeathre(ctx, city)
		if err != nil {
			log.Println(err)
		}
		weathers[i] = weather
	}

	wxMsg := "今日天气预报：\n\n"
	for i, weather := range weathers {
		wxMsg += fmt.Sprintf("城市：%s\n", citys[i])
		cast := weather.Forecasts[0].Casts[0]
		wxMsg += fmt.Sprintf("\t日期：%s\n", cast.Date)
		wxMsg += fmt.Sprintf("\t白昼天气：%s\n", cast.Dayweather)
		wxMsg += fmt.Sprintf("\t白昼温度：%s\n", cast.Daytemp)
		wxMsg += fmt.Sprintf("\t夜晚天气：%s\n", cast.Nightweather)
		wxMsg += fmt.Sprintf("\t夜晚温度：%s\n", cast.Nighttemp)
		wxMsg += fmt.Sprintf("\t风力：%s级\n", cast.Daypower)
		wxMsg += fmt.Sprintf("\n")
	}

	// 发送消息
	_ = wxbot.SendMsg(ctx, wxMsg, notified...)
}
