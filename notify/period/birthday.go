package period

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"github.com/Lofanmi/chinese-calendar-golang/lunar"
	"github.com/terloo/xiaochen/util"

	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/notify"
	"github.com/terloo/xiaochen/thirdparty/apispace"
	"github.com/terloo/xiaochen/wxbot"
)

type BirthdayNotifier struct {
}

var _ notify.Notifier = (*BirthdayNotifier)(nil)

type BirthdayPair struct {
	Key   family.People
	Value int
}

func (b *BirthdayNotifier) Notify(ctx context.Context, notified ...string) {
	msg := ""
	var todayBirth []family.People

	// 排序
	remainingDays := GetRemainingDays()
	pairs := make([]BirthdayPair, len(remainingDays))
	i := 0
	for k, v := range remainingDays {
		pairs[i] = BirthdayPair{k, v}
		i++
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Value < pairs[j].Value })

	// 通知生日
	for _, pair := range pairs {
		p := pair.Key
		remainingDay := pair.Value

		if remainingDay == -1 {
			continue
		}
		if remainingDay == 0 {
			todayBirth = append(todayBirth, p)
			continue
		}
		if remainingDay == 1 {
			msg += fmt.Sprintf("重大提醒！明天就是%s的生日了！\n", p.NickName)
			continue
		}
		if remainingDay == 7 || remainingDay <= 5 {
			msg += fmt.Sprintf("距离%s的生日只剩%d天了！\n", p.NickName, remainingDay)
			continue
		}
		if remainingDay == 20 {
			msg += fmt.Sprintf("还有%d天%s的生日就到了！\n", remainingDay, p.NickName)
			continue
		}
	}
	if msg == "" && len(todayBirth) == 0 {
		return
	}

	_ = wxbot.SendMsg(ctx, msg, notified...)

	// 通知生日花语
	for _, p := range todayBirth {
		flowerMsg := fmt.Sprintf("今天是%s的生日(%d-%d)，生日快乐！\n", p.NickName, p.Birthday.Month, p.Birthday.Day)
		flower, err := apispace.GetBirthdayFlower(ctx, p)
		if err != nil {
			log.Println(err)
			continue
		}
		flowerMsg += fmt.Sprintln(flower.BirthdayFlower)
		flowerMsg += fmt.Sprintln(flower.BirthdayFlowerContent)
		flowerMsg += fmt.Sprintln(flower.FlowerLng)
		flowerMsg += fmt.Sprintln(flower.FlowerLngContent)
		flowerMsg += fmt.Sprintln(flower.Birthstone)
		flowerMsg += fmt.Sprintln(flower.BirthstoneContent)

		_ = wxbot.SendMsg(ctx, flowerMsg, notified...)
	}
}

func GetRemainingDays() map[family.People]int {

	result := make(map[family.People]int)
	for _, p := range family.Families {
		result[p] = getRemainingDay(p)
	}
	return result
}

func getRemainingDay(p family.People) int {
	if p.Birthday.Date == "" {
		return -1
	}
	birthDay, err := time.ParseInLocation(util.DateLayout, p.Birthday.Date, time.Local)
	if err != nil {
		log.Printf("计算生日(%s)剩余时间错误 %s", p.Birthday.Date, err.Error())
		return -1
	}

	now := time.Now()
	nowYear := now.Year()
	nowMonth := now.Month()
	nowDay := now.Day()

	if !p.Birthday.Lunar {
		nextBirthDay := time.Date(nowYear, birthDay.Month(), birthDay.Day(), 0, 0, 0, 0, time.Local)
		if nextBirthDay.Year() == nowYear && nextBirthDay.Month() == nowMonth && nextBirthDay.Day() == nowDay {
			return 0
		}
		if nextBirthDay.Before(now) {
			nextBirthDay = time.Date(nowYear+1, birthDay.Month(), birthDay.Day(), 0, 0, 0, 0, time.Local)
		}
		return int(nextBirthDay.Sub(time.Date(nowYear, nowMonth, nowDay, 0, 0, 0, 0, time.Local)).Hours() / 24)
	} else {
		// 获取国历生日对应的农历月日
		_, month, day, _ := lunar.FromSolarTimestamp(birthDay.Unix())
		ZoneName, _ := birthDay.Zone()
		if ZoneName == "CDT" {
			// 特殊处理一下夏令时导致的天数差1
			day += 1
		}
		// 获取今年农历生日
		currentYearLunarBirthday := calendar.ByLunar(int64(nowYear), month, day, 0, 0, 0, false)
		nextBirthDay := time.Date(
			int(currentYearLunarBirthday.Solar.GetYear()),
			time.Month(currentYearLunarBirthday.Solar.GetMonth()),
			int(currentYearLunarBirthday.Solar.GetDay()),
			0, 0, 0, 0, time.Local,
		)
		if nextBirthDay.Year() == nowYear && nextBirthDay.Month() == nowMonth && nextBirthDay.Day() == nowDay {
			return 0
		}
		if nextBirthDay.Before(now) {
			// 考虑闰月的因素
			currentYearLunarBirthday = calendar.ByLunar(int64(nowYear), month, day, 0, 0, 0, true)
			nextBirthDay = time.Date(
				int(currentYearLunarBirthday.Solar.GetYear()),
				time.Month(currentYearLunarBirthday.Solar.GetMonth()),
				int(currentYearLunarBirthday.Solar.GetDay()),
				0, 0, 0, 0, time.Local,
			)
			if nextBirthDay.Before(now) {
				nextYearLunarBirthday := calendar.ByLunar(int64(nowYear+1), month, day, 0, 0, 0, false)
				nextBirthDay = time.Date(
					int(nextYearLunarBirthday.Solar.GetYear()),
					time.Month(nextYearLunarBirthday.Solar.GetMonth()),
					int(nextYearLunarBirthday.Solar.GetDay()),
					0, 0, 0, 0, time.Local,
				)
			}
		}
		return int(nextBirthDay.Sub(time.Date(nowYear, nowMonth, nowDay, 0, 0, 0, 0, time.Local)).Hours() / 24)
	}
}
