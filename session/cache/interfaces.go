package cache

import "context"

type MessageCache interface {
	SetValue(ctx context.Context, k string, v []byte) error
	GetValue(ctx context.Context, k string) ([]byte, error)
}
