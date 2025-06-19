package apispace

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/lunar"
	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/service/family"

	"github.com/terloo/xiaochen/client"
	"github.com/terloo/xiaochen/util"
)

func GetBirthdayFlower(ctx context.Context, p family.People) (*BirthdayFlowerData, error) {
	birthDay, err := time.Parse(util.DateLayout, p.Birthday.Date)
	if err != nil {
		return nil, errors.Errorf("解析日期(%s)错误 %s", p.Birthday.Date, err.Error())
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
		return nil, errors.WithMessagef(err, "花语接口HTTP错误：%s", string(b))
	}

	if len(birthdayFlower.Data) == 0 {
		return nil, errors.Errorf("花语接口返回值错误：%s", string(b))
	}
	result := birthdayFlower.Data[0]
	return &result, nil
}
