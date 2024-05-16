package period_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/notify/period"
)

func TestRemainDay(t *testing.T) {
	m := period.GetRemainingDays()
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
	(&period.BirthdayNotifier{}).Notify(context.Background(), family.TestChatroomWxid)
}
