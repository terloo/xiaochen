package message

import (
	"context"
	"testing"

	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
)

func TestGPTHandler(t *testing.T) {
	ctx := context.Background()
	handler := NewGPTHandler(CommonHandler{[]string{family.TestChatroomWxid}})
	err := handler.Handle(ctx, wxbot.FormattedMessage{
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[0].Wxid,
		Content: "谁来讲个笑话",
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	err = handler.Handle(ctx, wxbot.FormattedMessage{
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[2].Wxid,
		Content: "让xiaochen讲吧",
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	err = handler.Handle(ctx, wxbot.FormattedMessage{
		Self:    false,
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[2].Wxid,
		Content: "@xiaochen 来一个",
		At:      true,
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	err = handler.Handle(ctx, wxbot.FormattedMessage{
		Self:    false,
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[3].Wxid,
		Content: "@xiaochen 把上面那个笑话重复一遍呢，两次的回复要一模一样",
		At:      true,
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestMCPHandler1(t *testing.T) {
	ctx := context.Background()
	handler := NewGPTHandler(CommonHandler{[]string{family.TestChatroomWxid}})
	err := handler.Handle(ctx, wxbot.FormattedMessage{
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[0].Wxid,
		Content: "@xiaochen 你说今天我该搞点什么菜来吃",
		At:      true,
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	err = handler.Handle(ctx, wxbot.FormattedMessage{
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[0].Wxid,
		Content: "@xiaochen 说下第一道菜的详细做法，约详细越好",
		At:      true,
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestMCPHandler2(t *testing.T) {
	ctx := context.Background()
	handler := NewGPTHandler(CommonHandler{[]string{family.TestChatroomWxid}})
	err := handler.Handle(ctx, wxbot.FormattedMessage{
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[0].Wxid,
		Content: "@xiaochen 帮我查一下跳楼机是谁唱的",
		At:      true,
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	err = handler.Handle(ctx, wxbot.FormattedMessage{
		Chat:    family.TestChatroomWxid,
		Sender:  family.Families[0].Wxid,
		Content: "@xiaochen 把第一个下下来",
		At:      true,
	})
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
}
