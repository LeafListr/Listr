package client_test

import (
	"context"
	"encoding/json"
	"github.com/Linkinlog/LeafList/internal/client/curaleaf"
	"net/http"
	"testing"

	"github.com/Linkinlog/LeafList/internal/client"
)

func TestSend(t *testing.T) {
	tests := map[string]struct {
		ctx           context.Context
		query         string
		endpoint      client.Endpoint
		headers       http.Header
		response      interface{}
		expectedError bool
	}{
		"invalid endpoint / empty response": {
			ctx:      context.Background(),
			endpoint: "https://example.com",
			query:    ``,
		},
		"bad request": {
			ctx:           context.Background(),
			endpoint:      "https://example.com",
			query:         `<>/?`,
			expectedError: true,
		},
		"curaleaf - all products": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			query:    curaleaf.AllProductQuery(curaleaf.GbgId, curaleaf.MenuType),
			response: curaleaf.AllProductsResponse{},
		},
		"curaleaf - all vaporizer products": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			headers:  map[string][]string{"authority": {curaleaf.Authority}},
			query:    curaleaf.AllProductForCategoryQuery(curaleaf.GbgId, curaleaf.MenuType, "VAPORIZERS"),
			response: curaleaf.AllProductsResponse{},
		},
		"curaleaf - all offers": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			headers:  map[string][]string{"authority": {curaleaf.Authority}},
			query:    curaleaf.AllOffersQuery(curaleaf.GbgId, curaleaf.MenuType),
			response: curaleaf.AllProductsResponse{},
		},
		"curaleaf - all categories": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			headers:  map[string][]string{"authority": {curaleaf.Authority}},
			query:    curaleaf.AllCategoriesQuery(curaleaf.GbgId, curaleaf.MenuType),
			response: curaleaf.AllProductsResponse{},
		},
		"curaleaf - all locations": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			headers:  map[string][]string{"authority": {curaleaf.Authority}},
			query:    curaleaf.AllLocationsQuery(-79.5389, 40.3015),
			response: curaleaf.AllProductsResponse{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.headers == nil {
				tt.headers = make(map[string][]string)
			}
			c := client.NewHTTPClient(tt.endpoint, tt.headers)
			bs, err := c.Query(tt.ctx, tt.query, "POST")
			if err != nil && !tt.expectedError {
				t.Fatal(err)
			}

			if json.Valid(bs) {
				jsonErr := json.Unmarshal(bs, &tt.response)
				if jsonErr != nil {
					t.Fatal(jsonErr)
				}
			}
		})
	}
}
