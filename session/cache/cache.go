package cache

import (
	"context"
	"sync"

	"github.com/coocood/freecache"
	"github.com/pkg/errors"
)

type LocalCache struct {
	cache         *freecache.Cache
	expireSeconds int
	tmux          sync.Mutex
}

func NewLocalCache(capacity int, expireSeconds int) *LocalCache {
	return &LocalCache{cache: freecache.NewCache(capacity), expireSeconds: expireSeconds}
}

func (l *LocalCache) SetValue(ctx context.Context, k string, v []byte) error {
	err := l.cache.Set([]byte(k), v, l.expireSeconds)
	return errors.Wrapf(err, "set cache [%s]=[%s] error", k, v)
}

func (l *LocalCache) GetValue(ctx context.Context, k string) ([]byte, error) {
	value, err := l.cache.Get([]byte(k))
	if errors.Is(err, freecache.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "get cache key: [%s] error", k)
	}
	return value, err
}

var _ MessageCache = (*LocalCache)(nil)
