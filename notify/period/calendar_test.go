package period

import (
	"context"
	"testing"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/util"
)

func TestCalendarNotifier(t *testing.T) {
	(&CalendarNotifier{
		&util.SpyClock{
			SpyTimeStr: "2024-06-10",
		},
	}).Notify(context.Background(), family.TestChatroomWxid)
}
