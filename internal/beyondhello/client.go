package beyondhello

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Endpoint string

type HttpClient struct {
	hC      *http.Client
	headers http.Header
	e       Endpoint
}

func NewHTTPClient(endpoint Endpoint, headers http.Header) *HttpClient {
	return &HttpClient{
		hC:      &http.Client{Timeout: 30 * time.Second},
		headers: headers,
		e:       endpoint,
	}
}

func (c *HttpClient) Query(ctx context.Context, body string, method string) ([]byte, error) {
	return c.do(ctx, method, []byte(body))
}

func (c *HttpClient) do(ctx context.Context, method string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, string(c.e), bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}

	req.Header = c.headers
	req.Header.Set("Content-Type", "application/json")

	resp, hCErr := c.hC.Do(req)
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
