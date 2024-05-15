package main_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/terloo/xiaochen/almanac"
	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/wxbot"
)

func TestReport(t *testing.T) {
	ctx := context.Background()
	wxbot.KeepAlive(ctx)
	wxbot.ReportTest(ctx, family.TestChatroomWxid)
	wxbot.ReportWeather(ctx, family.TestChatroomWxid)
	for _, notifier := range notify.Notifiers {
		notifier.Notify(ctx, family.TestChatroomWxid)
	}
}

func TestContact(t *testing.T) {
	ctx := context.Background()
	contacts, _ := wxbot.GetContacts(ctx)
	log.Println(contacts)
}

func TestFoo(t *testing.T) {
	layout := "2006-01-02"
	parse, _ := time.ParseInLocation(layout, "1989-05-11", time.Local)
	log.Println(parse)
}

func TestCal(t *testing.T) {
	month := almanac.NewMonth(2024, 5)
	wxbot.SendMsg(context.Background(), family.TestChatroomWxid, month.FormatCal())
	log.Println(month.FormatCal())
}
