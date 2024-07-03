package period

import (
	"context"
	"fmt"
	"log"

	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/thirdparty/caiyun"
	"github.com/terloo/xiaochen/wxbot"
)

type WeatherNotifier struct {
}

var _ notify.Notifier = (*WeatherNotifier)(nil)

func (w *WeatherNotifier) Notify(ctx context.Context, notified ...string) {

	weatherMap := make(map[string]*caiyun.BaseBody, len(caiyun.WeatherLocation))
	for city, localtion := range caiyun.WeatherLocation {
		weather, err := caiyun.GetDailyWeather(ctx, localtion)
		if err != nil {
			log.Printf("获取天气异常：%+v", err)
			weather = nil
		}
		weatherMap[city] = weather
	}

	wxMsg := "今日天气预报：\n\n"
	for city, weather := range weatherMap {
		wxMsg += fmt.Sprintf("地区：%s\n", city)
		daily := weather.Result.Daily
		if daily.Status != "ok" {
			wxMsg += "获取天气信息异常"
			continue
		}
		wxMsg += fmt.Sprintf("\t主要天气：%s\n", caiyun.SkyconMap[daily.Skycon08H20H[0].Value])
		wxMsg += fmt.Sprintf("\t白天降水概率：%.0f%%\n", daily.Precipitation08H20H[0].Probability)
		wxMsg += fmt.Sprintf("\t白天最高气温：%.0f\n", daily.Temperature08H20H[0].Max)
		wxMsg += fmt.Sprintf("\t白天最低气温：%.0f\n", daily.Temperature08H20H[0].Min)
		wxMsg += fmt.Sprintf("\t白天平均气温：%.0f\n", daily.Temperature08H20H[0].Avg)
		wxMsg += fmt.Sprintf("\t平均相对湿度：%.2f\n", daily.Humidity[0].Avg)
		wxMsg += fmt.Sprintf("\t日出时间：出%s，落%s\n", daily.Astro[0].Sunrise.Time, daily.Astro[0].Sunset.Time)
		wxMsg += fmt.Sprintf("\t舒适指数：%s\n", daily.LifeIndex.Comfort[0].Desc)
		wxMsg += fmt.Sprintf("\t穿衣指数：%s\n", daily.LifeIndex.Dressing[0].Desc)
		wxMsg += fmt.Sprintf("\t紫外线指数：%s\n", daily.LifeIndex.Ultraviolet[0].Desc)
		wxMsg += fmt.Sprintf("\n")
	}

	// 发送消息
	_ = wxbot.SendMsg(ctx, wxMsg, notified...)
}
