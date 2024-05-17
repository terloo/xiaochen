package period

import (
	"context"
	"testing"

	"github.com/terloo/xiaochen/family"
)

func TestCalendarNotifier(t *testing.T) {
	(&CalendarNotifier{}).Notify(context.Background(), family.TestChatroomWxid)
}
