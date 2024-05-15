package wxbot

import (
	"context"
	"fmt"
	"log"

	"github.com/terloo/xiaochen/family"
	gaode2 "github.com/terloo/xiaochen/thirdparty/gaode"
	"github.com/terloo/xiaochen/thirdparty/gpt"
	"github.com/terloo/xiaochen/thirdparty/tsanghi"
)

func ReportTest(ctx context.Context, wxid string) {
	_ = SendMsg(ctx, wxid, "test test")
}

func KeepAlive(ctx context.Context) {
	_ = SendMsg(ctx, family.TestChatroomWxid, "1")
}

func ReportWeather(ctx context.Context, wxid string) {
	citys := []string{"江油市", "成都市"}
	weathers := make([]*gaode2.Weather, len(citys))
	for i, city := range citys {
		weather, err := gaode2.GetWeathre(ctx, city)
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
	_ = SendMsg(ctx, wxid, wxMsg)
}

func ReportHoliday(ctx context.Context, wxid string) {
}

func ReportTicker(ctx context.Context, wxid string, tickers []string) {
	wxMsg := "股票行情：\n\n"

	if len(tickers) == 0 {
		return
	}
	for _, ticker := range tickers {
		todayTickerData, err := tsanghi.GetTodayTicker(ctx, ticker)
		if err != nil {
			log.Println(err)
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

	_ = SendMsg(ctx, wxid, wxMsg)
}

func ResponseWithGPT(ctx context.Context, wxid string, message string) {
	s, err := gpt.Completion(ctx, message)
	respMessage := s
	if err != nil {
		respMessage = err.Error()
	}
	_ = SendMsg(ctx, wxid, respMessage)
}
