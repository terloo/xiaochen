package gaode

import "github.com/terloo/xiaochen/config"

var openKey = config.NewLoader("thirdparty.gaode.openKey")

var openHost = "https://restapi.amap.com/"

var weatherMap = map[string]string{
	"成都市": "510100",
	"江油市": "510781",
}
