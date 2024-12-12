package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/lunar"
	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/notify/period"
	"github.com/terloo/xiaochen/util"
	"github.com/terloo/xiaochen/wxbot"
)

type Birthday struct {
}

var _ Handler = (*Birthday)(nil)

func (z *Birthday) CommandName() string {
	return "生日"
}

func (z *Birthday) Exec(ctx context.Context, caller string, args []string) error {
	b := &period.BirthdayNotifier{
		&util.RealClock{},
		family.Families,
	}
	m := b.GetRemainingDays()
	msg := ""
	for p, remainDay := range m {
		if remainDay == -1 {
			continue
		}

		if p.Birthday.Date == "" {
			continue
		}
		birthDay, err := time.ParseInLocation(util.DateLayout, p.Birthday.Date, time.Local)
		if err != nil {
			log.Printf("计算生日(%s)剩余时间错误 %s", p.Birthday.Date, err.Error())
			continue
		}

		nickName := StrPad(p.NickName, 4, " ", "RIGHT")
		if p.Birthday.Lunar {
			typeStr := "农历"
			_, month, day, _ := lunar.FromSolarTimestamp(birthDay.Unix())
			msg += fmt.Sprintf("%-10s %d-%d\t %s\t 还有%d天过生\n", nickName, month, day, typeStr, remainDay)
		} else {
			typeStr := "国历"
			month, day := birthDay.Month(), birthDay.Day()
			msg += fmt.Sprintf("%-10s %d-%d\t %s\t 还有%d天过生\n", nickName, month, day, typeStr, remainDay)
		}
	}
	wxbot.SendMsg(ctx, msg, caller)
	return nil
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
