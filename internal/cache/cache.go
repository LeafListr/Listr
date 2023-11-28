package cache

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExpired  = errors.New("key expired")
)

type Cacher interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	GetOrRetrieve(key string, retriever func() (any, error)) (any, error)
}
