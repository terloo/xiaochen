package period

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/util"
)

func TestRemainDay(t *testing.T) {
	b := &BirthdayNotifier{
		&util.SpyClock{
			SpyTimeStr: "2024-05-15",
		},
		family.Families,
	}
	m := b.GetRemainingDays()
	msg := ""
	for p, day := range m {
		if day == -1 {
			continue
		}
		typeStr := "国历"
		if p.Birthday.Lunar {
			typeStr = "农历"
		}
		nickName := StrPad(p.NickName, 4, " ", "RIGHT")
		msg += fmt.Sprintf("%-10s %s\t %s\t 还有%d天过生\n", nickName, p.Birthday.Date, typeStr, day)
	}
	log.Println(msg)
	// wxbot.SendMsg(context.Background(), family.TestChatroomWxid, msg)
}

// StrPad
// input string 原字符串
// padLength int 规定补齐后的字符串位数
// padString string 自定义填充字符串
// padType string 填充类型:LEFT(向左填充,自动补齐位数), 默认右侧
func StrPad(input string, padLength int, padString string, padType string) string {

	output := ""
	inputLen := len(input)

	if inputLen >= padLength {
		return input
	}

	padStringLen := len(padString)
	needFillLen := padLength - inputLen

	if diffLen := padStringLen - needFillLen; diffLen > 0 {
		padString = padString[diffLen:]
	}

	for i := 1; i <= needFillLen; i += padStringLen {
		output += padString
	}
	switch padType {
	case "LEFT":
		return output + input
	default:
		return input + output
	}
}

func TestBirthdayNotifier(t *testing.T) {
	(&BirthdayNotifier{
		&util.SpyClock{
			SpyTimeStr: "2024-05-15",
		},
		family.Families,
	}).Notify(context.Background(), family.TestChatroomWxid)
}
