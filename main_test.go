package main_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/terloo/almanac"
	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/notify/period"
	"github.com/terloo/xiaochen/util"
	"github.com/terloo/xiaochen/wxbot"
)

func TestReport(t *testing.T) {
	ctx := context.Background()
	wxbot.KeepAlive(ctx)
}

func TestContact(t *testing.T) {
	ctx := context.Background()
	contacts, _ := wxbot.GetContacts(ctx)
	log.Println(contacts)
}

func TestFoo(t *testing.T) {
	parse, _ := time.ParseInLocation(util.DateLayout, "1989-05-11", time.Local)
	log.Println(parse)
}

func TestCal(t *testing.T) {
	month := almanac.NewMonth(2024, 5)
	wxbot.SendMsg(context.Background(), month.FormatCal(), family.TestChatroomWxid)
	log.Println(month.FormatCal())
}

func TestWeather(t *testing.T) {
	weatherNotifier := period.WeatherNotifier{}
	weatherNotifier.Notify(context.Background(), family.TestChatroomWxid)
}