package graphql

import (
	"context"
)

type Client interface {
	Query(ctx context.Context, query string, variables map[string]interface{}, method string) ([]byte, error)
}
