package caches

import (
	"context"
	"fmt"
	"sync"

	"code.gopub.tech/errors"
)

var ErrNoSuchValue = fmt.Errorf("no such value")

func GetDefaultCache() ICache  { return defaultCache }
func SetDefaultCache(c ICache) { defaultCache = c }

var defaultCache = NewSimpleCache()

func NewSimpleCache() ICache { return new(cache) }

type Fetcher func(ctx context.Context, key string) (any, error)

type ICache interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any) error
	GetOrFetch(ctx context.Context, key string, fetcher Fetcher) (any, error)
}

type cache struct {
	m sync.Map
}

func (c *cache) Get(ctx context.Context, key string) (any, error) {
	v, ok := c.m.Load(key)
	if ok {
		return v, nil
	}
	return nil, errors.Wrapf(ErrNoSuchValue, "key=%v", key)
}

func (c *cache) GetOrFetch(ctx context.Context, key string, fetcher Fetcher) (any, error) {
	v, ok := c.m.Load(key)
	if ok {
		return v, nil
	}
	v, err := fetcher(ctx, key)
	if err != nil {
		return v, err
	}
	c.m.LoadOrStore(key, v)
	return v, nil
}

func (c *cache) Set(ctx context.Context, key string, value any) error {
	c.m.Store(key, value)
	return nil
}
