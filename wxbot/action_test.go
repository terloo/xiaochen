package wxbot_test

import (
	"context"
	"testing"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/wxbot"
)

func TestGPT(t *testing.T) {
	wxbot.ResponseWithGPT(context.Background(), family.TestChatroomWxid, "你好啊!")
}
