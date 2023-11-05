package graphql

import (
	"context"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Client
type Client interface {
	Query(ctx context.Context, query string, variables map[string]interface{}, method string) ([]byte, error)
}
