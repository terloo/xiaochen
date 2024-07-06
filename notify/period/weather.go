package period

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"golang.org/x/exp/slices"

	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/thirdparty/caiyun"
	"github.com/terloo/xiaochen/wxbot"
)

type WeatherNotifier struct {
}

var _ notify.Notifier = (*WeatherNotifier)(nil)

func (w *WeatherNotifier) Notify(ctx context.Context, notified ...string) {

	weatherMap := make(map[string]*caiyun.DailyResp, len(caiyun.WeatherLocation))
	for city, location := range caiyun.WeatherLocation {
		dailyWeather, err := caiyun.GetDailyWeather(ctx, location)
		if err != nil {
			log.Printf("获取%s天气异常：%+v", city, err)
			dailyWeather = nil
		}
		weatherMap[city] = dailyWeather
	}

	wxMsg := "今日天气预报：\n\n"
	for city, daily := range weatherMap {
		wxMsg += fmt.Sprintf("地区：%s\n", city)
		if daily.Status != "ok" {
			wxMsg += "获取天气信息异常"
			continue
		}
		wxMsg += fmt.Sprintf("\t主要天气：%s\n", caiyun.SkyconMap[daily.Skycon08H20H[0].Value].Sino)
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

type WeatherHourlyNotifier struct {
	LastHourWeather map[string]caiyun.Skycon
}

var _ notify.Notifier = (*WeatherHourlyNotifier)(nil)

func (w *WeatherHourlyNotifier) Notify(ctx context.Context, notified ...string) {

	// 早上7点重置
	if time.Now().Hour() == 7 {
		w.LastHourWeather = nil
	}

	if w.LastHourWeather == nil {
		w.LastHourWeather = make(map[string]caiyun.Skycon)
	}

	AlertMsg := ""
	WeatherConvertMsg := ""
	for city, location := range caiyun.WeatherLocation {
		hourlyResp, alertResp, err := caiyun.GetHourlyWeather(ctx, location)
		if err != nil {
			log.Printf("获取%s天气异常：%+v", city, err)
			hourlyResp = nil
			continue
		}

		lastSkycon, ok := w.LastHourWeather[city]
		currentSkycon := caiyun.SkyconMap[hourlyResp.Skycon[0].Value]
		if !ok {
			w.LastHourWeather[city] = currentSkycon
			continue
		}

		// 天气转变
		log.Printf("ciry: %s, lastSkycon: %s, currentSkycon: %s", city, lastSkycon, currentSkycon)
		currentSkyconPriority := slices.Index(caiyun.SkyconPriority, currentSkycon)
		lastSkyconPriority := slices.Index(caiyun.SkyconPriority, lastSkycon)
		if (currentSkyconPriority < 15 || lastSkyconPriority < 15) && math.Abs(float64(currentSkyconPriority-lastSkyconPriority)) > 1 {
			WeatherConvertMsg += fmt.Sprintf("%s的天气即将由%s转为%s\n", city, lastSkycon.Sino, currentSkycon.Sino)
		}
		w.LastHourWeather[city] = currentSkycon

		// 预警信息
		if alertResp != nil && len(alertResp.Description) != 0 {
			log.Printf("%s产生预警信息：%s", city, alertResp.Description)
			AlertMsg += fmt.Sprintf("%s\n", alertResp.Description)
		}

	}

	if len(AlertMsg) != 0 {
		_ = wxbot.SendMsg(ctx, AlertMsg, notified...)
	}
	if len(WeatherConvertMsg) != 0 {
		_ = wxbot.SendMsg(ctx, WeatherConvertMsg, notified...)
	}

}
