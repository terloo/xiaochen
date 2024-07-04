package caiyun

import (
	"github.com/terloo/xiaochen/config"
)

var openKey = config.NewLoader("thirdparty.caiyun.openKey")

var openHost = "http://api.caiyunapp.com/"

var WeatherLocation = map[string]string{
	"成都市华府大道": "104.06,30.53",
	"江油市":     "104.73,31.78",
	"武都镇":     "104.78,31.88",
	"北京市西城区":  "116.35,39.86",
}

var SkyconMap = map[string]Skycon{}

var SkyconPriority []Skycon

func init() {
	for _, s := range skycons {
		SkyconMap[s.Enum] = s
		SkyconPriority = append(SkyconPriority, s)
	}
}

type Skycon struct {
	Enum string
	Sino string
}

var skycons = []Skycon{
	{Enum: "STORM_SNOW", Sino: "暴雪"},
	{Enum: "HEAVY_SNOW", Sino: "大雪"},
	{Enum: "MODERATE_SNOW", Sino: "中雪"},
	{Enum: "LIGHT_SNOW", Sino: "小雪"},
	{Enum: "STORM_RAIN", Sino: "暴雨"},
	{Enum: "HEAVY_RAIN", Sino: "大雨"},
	{Enum: "MODERATE_RAIN", Sino: "中雨"},
	{Enum: "LIGHT_RAIN", Sino: "小雨"},
	{Enum: "FOG", Sino: "雾"},
	{Enum: "SAND", Sino: "沙尘"},
	{Enum: "DUST", Sino: "浮尘"},
	{Enum: "HEAVY_HAZE", Sino: "重度雾霾"},
	{Enum: "MODERATE_HAZE", Sino: "中度雾霾"},
	{Enum: "LIGHT_HAZE", Sino: "轻度雾霾"},
	{Enum: "WIND", Sino: "大风"},
	{Enum: "CLOUDY", Sino: "阴"},
	{Enum: "PARTLY_CLOUDY_DAY", Sino: "多云"},
	{Enum: "PARTLY_CLOUDY_NIGHT", Sino: "多云"},
	{Enum: "CLEAR_DAY", Sino: "晴"},
	{Enum: "CLEAR_NIGHT", Sino: "晴"},
}
