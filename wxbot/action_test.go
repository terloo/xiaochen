package wxbot_test

import (
	"context"
	"testing"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/wxbot"
)

func TestTricker(t *testing.T) {
	wxbot.ReportTicker(context.Background(), family.TestChatroomWxid, []string{"600959"})
}

func TestGPT(t *testing.T) {
	wxbot.ResponseWithGPT(context.Background(), family.TestChatroomWxid, "你好啊!")
}

func TestReportHoliday(t *testing.T) {
	wxbot.ReportHoliday(context.Background(), family.TestChatroomWxid)
}
