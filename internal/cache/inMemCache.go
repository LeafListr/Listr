package cache

import (
	"errors"
	"log/slog"
	"sync"
	"time"
)

type cacheItem struct {
	value     any
	ttl       time.Duration
	timeStamp int64
}

type inMemCache struct {
	mu sync.RWMutex
	m  map[string]cacheItem
}

func NewCache() Cacher {
	return &inMemCache{
		m: make(map[string]cacheItem),
	}
}

func (c *inMemCache) GetOrRetrieve(key string, ttl time.Duration, retriever func() (any, error)) (any, error) {
	value, err := c.Get(key)
	if err == nil {
		return value, nil
	}
	if !errors.Is(err, ErrKeyNotFound) && !errors.Is(err, ErrKeyExpired) {
		return nil, err
	}

	retrievedValue, retrieveErr := retriever()
	if retrieveErr != nil {
		return nil, retrieveErr
	}

	cacheSetErr := c.Set(key, ttl, retrievedValue)
	if cacheSetErr != nil {
		return nil, cacheSetErr
	}

	return retrievedValue, nil
}

func (c *inMemCache) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	if time.Now().UnixNano()-value.timeStamp > int64(value.ttl) {
		return nil, ErrKeyExpired
	}
	slog.Debug("cache hit", slog.String("key", key))
	return value.value, nil
}

func (c *inMemCache) Set(key string, ttl time.Duration, value any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	slog.Debug("cache set", slog.String("key", key))
	c.m[key] = cacheItem{value: value, timeStamp: time.Now().UnixNano(), ttl: ttl}
	return nil
}
