package client

import (
	"context"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type Endpoint string

//counterfeiter:generate . Client
type Client interface {
	Query(ctx context.Context, body string, method string) ([]byte, error)
}
