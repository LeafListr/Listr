package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Endpoint string

type HttpClient struct {
	hC *http.Client
	hH http.Header
	e  Endpoint
}

func NewHTTPClient(endpoint Endpoint, headers http.Header) *HttpClient {
	return &HttpClient{
		hC: &http.Client{Timeout: 30 * time.Second},
		hH: headers,
		e:  endpoint,
	}
}

func (c *HttpClient) Query(ctx context.Context, query string, variables map[string]interface{}, method string) ([]byte, error) {
	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return c.do(ctx, method, reqBody)
}

func (c *HttpClient) do(ctx context.Context, method string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, string(c.e), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header = c.hH
	req.Header.Set("Content-Type", "application/json")

	resp, hCErr := c.hC.Do(req)
	if hCErr != nil {
		return nil, hCErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bs, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	return bs, nil
}
