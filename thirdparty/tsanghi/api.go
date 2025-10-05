package tsanghi

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/client"
	"github.com/terloo/xiaochen/util"
)

func GetTodayTicker(ctx context.Context, code string) (*TickerData, error) {
	param := url.Values{
		"token":      []string{openKey.Get()},
		"ticker":     []string{code},
		"start_date": []string{time.Now().Format(util.DateLayout)},
	}
	b, err := client.HttpGet(ctx, openHost+"daily", nil, param)
	if err != nil {
		return nil, err
	}

	ticker := &Ticker{}
	err = json.Unmarshal(b, ticker)
	if err != nil {
		return nil, err
	}

	if len(ticker.Data) < 1 {
		return nil, errors.New("no ticket data")
	}
	data := ticker.Data[0]
	return &data, nil
}
