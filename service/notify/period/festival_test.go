package period

import (
	"context"
	"testing"

	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/util"
)

func TestFestivalNotifier(t *testing.T) {
	(&FestivalNotifier{
		&util.SpyClock{
			SpyTimeStr: "2023-08-18",
		},
	}).Notify(context.Background(), family.TestChatroomWxid)
}
