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
}

var SkyconMap = map[string]string{
	"CLEAR_DAY":           "晴",
	"CLEAR_NIGHT":         "晴",
	"PARTLY_CLOUDY_DAY":   "多云",
	"PARTLY_CLOUDY_NIGHT": "多云",
	"CLOUDY":              "阴",
	"LIGHT_HAZE":          "轻度雾霾",
	"MODERATE_HAZE":       "中度雾霾",
	"HEAVY_HAZE":          "重度雾霾",
	"LIGHT_RAIN":          "小雨",
	"MODERATE_RAIN":       "中雨",
	"HEAVY_RAIN":          "大雨",
	"STORM_RAIN":          "暴雨",
	"FOG":                 "雾",
	"LIGHT_SNOW":          "小雪",
	"MODERATE_SNOW":       "中雪",
	"HEAVY_SNOW":          "大雪",
	"STORM_SNOW":          "暴雪",
	"DUST":                "浮尘",
	"SAND":                "沙尘",
	"WIND":                "大风",
}
