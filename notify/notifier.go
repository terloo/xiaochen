package notify

import "context"

type Notifier interface {
	Notify(ctx context.Context, wxid string)
}

var Notifiers = []Notifier{
	&BirthdayNotifier{},
}
