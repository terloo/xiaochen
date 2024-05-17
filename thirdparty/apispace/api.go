package apispace

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/lunar"
	"github.com/terloo/xiaochen/client"
	"github.com/terloo/xiaochen/family"
	"github.com/terloo/xiaochen/util"
)

func GetBirthdayFlower(ctx context.Context, p family.People) (*BirthdayFlowerData, error) {
	birthDay, err := time.Parse(util.DateLayout, p.Birthday.Date)
	if err != nil {
		errorStr := fmt.Sprintf("计算生日(%s)剩余时间错误 %s", p.Birthday.Date, err.Error())
		log.Printf(errorStr)
		return nil, errors.New(errorStr)
	}

	month := int64(birthDay.Month())
	day := int64(birthDay.Day())
	if p.Birthday.Lunar {
		_, month, day, _ = lunar.FromSolarTimestamp(birthDay.Unix())
	}

	param := url.Values{}
	param.Add("moon", strconv.Itoa(int(month)))
	param.Add("day", strconv.Itoa(int(day)))
	header := http.Header{}
	header.Add("X-APISpace-Token", openKey.Get())

	b, err := client.HttpGet(ctx, openHost+"birthday-flowers/api/v1/xzw/birthday_flower/", header, param)
	if err != nil {
		return nil, err
	}
	birthdayFlower := &BirthdayFlower{}
	err = json.Unmarshal(b, birthdayFlower)
	if err != nil {
		return nil, err
	}

	if len(birthdayFlower.Data) == 0 {
		return nil, errors.New("花语接口返回有误：" + string(b))
	}
	result := birthdayFlower.Data[0]
	return &result, nil
}
