package graphql_test

import (
	"context"
	"encoding/json"
	"github.com/Linkinlog/LeafList/internal/graphql/curaleaf"
	"net/http"
	"testing"

	"github.com/Linkinlog/LeafList/internal/graphql"
)

func TestSend(t *testing.T) {
	tests := map[string]struct {
		ctx           context.Context
		query         string
		endpoint      graphql.Endpoint
		variables     map[string]interface{}
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
			ctx:      context.Background(),
			endpoint: "https://example.com",
			query:    `<>/?`,
			variables: map[string]interface{}{
				"invalid": make(chan int),
			},
			expectedError: true,
		},
		"curaleaf - all products": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			query:    curaleaf.AllProductQuery(),
			variables: map[string]interface{}{
				"dispensaryUniqueId": curaleaf.GbgId,
				"menuType":           curaleaf.MenuType,
			},
			response: curaleaf.AllProductResponse{},
		},
		"curaleaf - all oral products": {
			ctx:      context.Background(),
			endpoint: curaleaf.Endpoint,
			headers:  map[string][]string{"authority": {curaleaf.Authority}},
			query:    curaleaf.AllProductForCategoryQuery(),
			variables: map[string]interface{}{
				"dispensaryUniqueId": curaleaf.GbgId,
				"menuType":           curaleaf.MenuType,
				"category":           "oral",
			},
			response: curaleaf.AllProductForCategoryResponse{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.headers == nil {
				tt.headers = make(map[string][]string)
			}
			c := graphql.NewHTTPClient(tt.endpoint, tt.headers)
			bs, err := c.Query(tt.ctx, tt.query, tt.variables, "POST")
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
