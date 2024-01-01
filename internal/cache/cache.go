package cache

import (
	"errors"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExpired  = errors.New("key expired")
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Cacher
type Cacher interface {
	Set(key string, ttl time.Duration, value any) error
	Get(key string) (any, error)
	GetOrRetrieve(key string, ttl time.Duration, retriever func() (any, error)) (any, error)
}
