package notify

import (
	"context"
)

type Notifier interface {
	Notify(ctx context.Context, notified ...string)
}
