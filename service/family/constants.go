package family

import (
	"log"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/lunar"
	"github.com/terloo/xiaochen/storage"

	"github.com/terloo/xiaochen/util"
)

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

	families, err := storage.WxFamilyRepo.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	familyChatroom, err := storage.WxChatRepo.FindByName("Family")
	if err != nil {
		log.Fatal(err)
	}
	FamilyChatroomWxid = familyChatroom.Wxid
	testChatroom, err := storage.WxChatRepo.FindByName("Test")
	if err != nil {
		log.Fatal(err)
	}
	TestChatroomWxid = testChatroom.Wxid
	momChatroom, err := storage.WxChatRepo.FindByName("Mom")
	if err != nil {
		log.Fatal(err)
	}
	MomWxid = momChatroom.Wxid

	for _, f := range families {
		NameToWxid[f.NickName] = f.Wxid
		WxidToName[f.Wxid] = f.NickName
	}

	// 计算生日月和天
	for _, p := range families {
		if p.Birthday == "" {
			continue
		}
		birthDay, err := time.Parse(util.DateLayout, p.Birthday)
		if err != nil {
			log.Printf("计算生日(%s)剩余时间错误 %s", p.Birthday, err.Error())
			continue
		}
		people := People{
			NickName: p.NickName,
			Wxid:     p.Wxid,
		}
		people.Birthday.Date = p.Birthday
		people.Birthday.Lunar = p.Lunar
		if !p.Lunar {
			people.Birthday.Month = int(birthDay.Month())
			people.Birthday.Day = birthDay.Day()
		} else {
			_, month, day, _ := lunar.FromSolarTimestamp(birthDay.Unix())
			ZoneName, _ := birthDay.Zone()
			if ZoneName == "CDT" {
				// 特殊处理一下夏令时导致的天数差1
				day += 1
			}
			people.Birthday.Month = int(month)
			people.Birthday.Day = int(day)
		}
		Families = append(Families, people)
	}

}
