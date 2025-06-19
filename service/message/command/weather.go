package command

import (
	"context"

	"github.com/terloo/xiaochen/service/notify/period"
)

type Weather struct {
}

var _ Handler = (*Weather)(nil)

func (w *Weather) CommandName() string {
	return "天气"
}

func (w *Weather) Exec(ctx context.Context, caller string, args []string) error {
	weatherNotifier := period.WeatherNotifier{}
	weatherNotifier.Notify(ctx, caller)
	return nil
}
