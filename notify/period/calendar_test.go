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
			SpyTimeStr: "2023-04-05",
		},
	}).Notify(context.Background(), family.TestChatroomWxid)
}
