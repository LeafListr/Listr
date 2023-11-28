package cache

import (
	"log/slog"
	"sync"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"
)

type cacheItem struct {
	value     any
	timeStamp int64
}

type inMemCache struct {
	mu sync.RWMutex
	m  map[string]cacheItem
	ttl time.Duration
}

func NewCache() cache.Cacher {
	return &inMemCache{
		m: make(map[string]cacheItem),
		ttl: time.Minute * 5,
	}
}

func (c *inMemCache) GetOrRetrieve(key string, retriever func() (any, error)) (any, error) {
	value, err := c.Get(key)
	if err == cache.ErrKeyNotFound || err == cache.ErrKeyExpired {
		value, err = retriever()
		if err != nil {
			return nil, err
		}
		if err = c.Set(key, value); err != nil {
			return nil, err
		}
	}
	return value, err
}

func (c *inMemCache) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	if !ok {
		return nil, cache.ErrKeyNotFound
	} else if time.Now().UnixNano() - value.timeStamp > int64(c.ttl) {
		return nil, cache.ErrKeyExpired
	}
	slog.Debug("cache hit", slog.String("key", key))
	return value.value, nil
}

func (c *inMemCache) Set(key string, value any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	slog.Debug("cache set", slog.String("key", key))
	c.m[key] = cacheItem{value: value, timeStamp: time.Now().UnixNano()}
	return nil
}
