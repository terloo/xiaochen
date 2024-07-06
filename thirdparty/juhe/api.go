package juhe

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/terloo/xiaochen/client"
)

func GetZhouGong(ctx context.Context, keyword string) (*ZhouGongResult, error) {
	param := url.Values{
		"q":    []string{keyword},
		"full": []string{"1"},
		"key":  []string{openKey.Get()},
	}
	b, err := client.HttpGet(ctx, openHost+"dream/query", nil, param)
	if err != nil {
		return nil, err
	}

	zhougong := &ZhouGong{}
	err = json.Unmarshal(b, zhougong)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(zhougong.Result) > 0 {
		return &zhougong.Result[0], nil
	}
	return nil, errors.New("未查询到相关关键词")
}
