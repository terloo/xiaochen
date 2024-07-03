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

func GetDailyWeather(ctx context.Context, location string) (*BaseBody, error) {
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
	return resp, nil
}
