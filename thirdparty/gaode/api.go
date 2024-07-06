package gaode

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/terloo/xiaochen/client"
)

func GetWeathre(ctx context.Context, city string) (*Weather, error) {
	city_nu, ok := weatherMap[city]
	if !ok {
		return nil, errors.New("unknown city: " + city)
	}

	param := url.Values{
		"city":       []string{city_nu},
		"key":        []string{openKey.Get()},
		"extensions": []string{"all"},
	}
	b, err := client.HttpGet(ctx, openHost+"v3/weather/weatherInfo", nil, param)
	if err != nil {
		return nil, err
	}

	gaodeWather := &Weather{}
	err = json.Unmarshal(b, gaodeWather)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gaodeWather, nil
}
