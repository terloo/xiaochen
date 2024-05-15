package family

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/lunar"
)

type Constant struct {
	FamilyChatroomWxid string   `json:"family_chatroom_wxid"`
	TestChatroomWxid   string   `json:"test_chatroom_wxid"`
	MomWxid            string   `json:"mom_wxid"`
	Families           []People `json:"families"`
}

type People struct {
	NickName string   `json:"nick_name"`
	Wxid     string   `json:"wxid"`
	Birthday Birthday `json:"birthday"`
}

type Birthday struct {
	Date  string `json:"date"`
	Lunar bool   `json:"lunar"`
	Month int    `json:"month,omitempty"`
	Day   int    `json:"day,omitempty"`
}

var Families []People

var NameToWxid = map[string]string{}

var WxidToName = map[string]string{}

var FamilyChatroomWxid string

var TestChatroomWxid string

var MomWxid string

func init() {

	file, err := os.ReadFile("families.json")
	if err != nil {
		log.Fatal(err)
	}
	cons := &Constant{}
	err = json.Unmarshal(file, &cons)
	if err != nil {
		log.Fatal(err)
	}

	Families = cons.Families
	FamilyChatroomWxid = cons.FamilyChatroomWxid
	TestChatroomWxid = cons.TestChatroomWxid
	MomWxid = cons.MomWxid

	for _, f := range Families {
		NameToWxid[f.NickName] = f.Wxid
		WxidToName[f.Wxid] = f.NickName
	}

	// 计算生日月和天
	for i, p := range Families {
		if p.Birthday.Date == "" {
			continue
		}
		timeFormat := "2006-01-02"
		birthDay, err := time.Parse(timeFormat, p.Birthday.Date)
		if err != nil {
			log.Printf("计算生日(%s)剩余时间错误 %s", p.Birthday.Date, err.Error())
			continue
		}
		if !p.Birthday.Lunar {
			p.Birthday.Month = int(birthDay.Month())
			p.Birthday.Day = birthDay.Day()
		} else {
			_, month, day, _ := lunar.FromSolarTimestamp(birthDay.Unix())
			ZoneName, _ := birthDay.Zone()
			if ZoneName == "CDT" {
				// 特殊处理一下夏令时导致的天数差1
				day += 1
			}
			p.Birthday.Month = int(month)
			p.Birthday.Day = int(day)
		}
		Families[i] = p
	}

}
