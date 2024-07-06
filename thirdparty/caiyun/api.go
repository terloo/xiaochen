package caiyun

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/terloo/xiaochen/client"
)

func GetDailyWeather(ctx context.Context, location string) (*DailyResp, error) {
	param := url.Values{}
	param.Add("dailysteps", "1")
	header := http.Header{}

	_url := fmt.Sprintf("%s/v2.6/%s/%s/daily", openHost, openKey.Get(), location)
	b, err := client.HttpGet(ctx, _url, header, param)
	if err != nil {
		return nil, err
	}
	resp := &BaseBody{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, errors.Wrapf(err, "天气接口返回错误：%s", string(b))
	}
	return &resp.Result.Daily, nil
}

func GetHourlyWeather(ctx context.Context, location string) (*HourlyResp, *AlertResp, error) {
	param := url.Values{}
	param.Add("hourlysteps", "2")
	param.Add("alert", "true")
	header := http.Header{}

	_url := fmt.Sprintf("%s/v2.6/%s/%s/hourly", openHost, openKey.Get(), location)
	b, err := client.HttpGet(ctx, _url, header, param)
	if err != nil {
		return nil, nil, err
	}
	resp := &BaseBody{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "天气接口返回错误：%s", string(b))
	}
	return &resp.Result.Hourly, &resp.Result.Alert, nil
}
