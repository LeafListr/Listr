package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Client
type Client interface {
	Query(ctx context.Context, body string, method string) ([]byte, error)
	SetEndpoint(e Endpoint)
}

type Endpoint string

type HttpClient struct {
	hC      *http.Client
	headers http.Header
	e       Endpoint
	c       cache.Cacher
}

func NewHTTPClient(endpoint Endpoint, headers http.Header, c cache.Cacher) *HttpClient {
	return &HttpClient{
		hC:      &http.Client{Timeout: 30 * time.Second},
		headers: headers,
		e:       endpoint,
		c:       c,
	}
}

func (hc *HttpClient) SetEndpoint(e Endpoint) {
	hc.e = e
}

func (hc *HttpClient) Query(ctx context.Context, body string, method string) ([]byte, error) {
	cacheKey := fmt.Sprintf("%s-%s", method, body)
	cachedData, err := hc.c.Get(cacheKey)
	if err == nil {
		return cachedData.([]byte), nil
	}

	responseData, err := hc.do(ctx, method, []byte(body))
	if err != nil {
		return nil, err
	}

	go func() {
		err = hc.c.Set(cacheKey, 5*time.Minute, responseData)
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	return responseData, nil
}

func (hc *HttpClient) do(ctx context.Context, method string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, string(hc.e), bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}

	req.Header = hc.headers

	resp, hCErr := hc.hC.Do(req)
	if hCErr != nil {
		return []byte{}, hCErr
	}
	defer resp.Body.Close()

	respBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return []byte{}, readErr
	}

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("unexpected status code: %d\n response: %s", resp.StatusCode, string(respBytes))
	}

	return respBytes, nil
}
