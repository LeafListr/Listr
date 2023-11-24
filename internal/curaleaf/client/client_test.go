package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/repository"
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
			endpoint: repository.GqlEndpoint,
			query:    client.AllProductQuery(repository.GbgId, repository.MenuType),
			response: &client.Response{},
		},
		"curaleaf - all vaporizer products": {
			ctx:      context.Background(),
			endpoint: repository.GqlEndpoint,
			headers:  map[string][]string{"authority": {repository.Authority}},
			query:    client.AllProductForCategoryQuery(repository.GbgId, repository.MenuType, "VAPORIZERS"),
			response: &client.Response{},
		},
		"curaleaf - all offers": {
			ctx:      context.Background(),
			endpoint: repository.GqlEndpoint,
			headers:  map[string][]string{"authority": {repository.Authority}},
			query:    client.AllOffersQuery(repository.GbgId, repository.MenuType),
			response: &client.Response{},
		},
		"curaleaf - all categories": {
			ctx:      context.Background(),
			endpoint: repository.GqlEndpoint,
			headers:  map[string][]string{"authority": {repository.Authority}},
			query:    client.AllCategoriesQuery(repository.GbgId, repository.MenuType),
			response: &client.Response{},
		},
		"curaleaf - all locations": {
			ctx:      context.Background(),
			endpoint: repository.GqlEndpoint,
			headers:  map[string][]string{"authority": {repository.Authority}},
			query:    client.AllLocationsQuery(-79.5389, 40.3015),
			response: &client.Response{},
		},
		"curaleaf - invalid menu": {
			ctx:           context.Background(),
			endpoint:      repository.GqlEndpoint,
			query:         client.AllProductQuery("foo", repository.MenuType),
			response:      &client.Response{},
			expectedError: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.headers == nil {
				tt.headers = make(map[string][]string)
			}
			c := client.NewHTTPClient(tt.endpoint, tt.headers)
			respBytes, err := c.Query(tt.ctx, tt.query, "POST")
			if err != nil && !tt.expectedError {
				t.Fatal(err)
			}

			if json.Valid(respBytes) {
				jsonErr := json.Unmarshal(respBytes, &tt.response)
				if jsonErr != nil {
					t.Fatal(jsonErr)
				}
				if responseObj, ok := tt.response.(*client.Response); ok {
					if responseObj.ErrorObj.Errors != nil && !tt.expectedError {
						t.Fatal(responseObj.Errors[0].Message)
					}
				}
			}
		})
	}
}
