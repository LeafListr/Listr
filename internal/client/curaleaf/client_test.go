package curaleaf_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Linkinlog/LeafList/internal/client/curaleaf"

	curaRepo "github.com/Linkinlog/LeafList/internal/repository/curaleaf"

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
			endpoint: curaRepo.GqlEndpoint,
			query:    curaleaf.AllProductQuery(curaRepo.GbgId, curaRepo.MenuType),
			response: &curaleaf.Response{},
		},
		"curaleaf - all vaporizer products": {
			ctx:      context.Background(),
			endpoint: curaRepo.GqlEndpoint,
			headers:  map[string][]string{"authority": {curaRepo.Authority}},
			query:    curaleaf.AllProductForCategoryQuery(curaRepo.GbgId, curaRepo.MenuType, "VAPORIZERS"),
			response: &curaleaf.Response{},
		},
		"curaleaf - all offers": {
			ctx:      context.Background(),
			endpoint: curaRepo.GqlEndpoint,
			headers:  map[string][]string{"authority": {curaRepo.Authority}},
			query:    curaleaf.AllOffersQuery(curaRepo.GbgId, curaRepo.MenuType),
			response: &curaleaf.Response{},
		},
		"curaleaf - all categories": {
			ctx:      context.Background(),
			endpoint: curaRepo.GqlEndpoint,
			headers:  map[string][]string{"authority": {curaRepo.Authority}},
			query:    curaleaf.AllCategoriesQuery(curaRepo.GbgId, curaRepo.MenuType),
			response: &curaleaf.Response{},
		},
		"curaleaf - all locations": {
			ctx:      context.Background(),
			endpoint: curaRepo.GqlEndpoint,
			headers:  map[string][]string{"authority": {curaRepo.Authority}},
			query:    curaleaf.AllLocationsQuery(-79.5389, 40.3015),
			response: &curaleaf.Response{},
		},
		"curaleaf - invalid menu": {
			ctx:           context.Background(),
			endpoint:      curaRepo.GqlEndpoint,
			query:         curaleaf.AllProductQuery("foo", curaRepo.MenuType),
			response:      &curaleaf.Response{},
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
				if responseObj, ok := tt.response.(*curaleaf.Response); ok {
					if responseObj.ErrorObj != nil && !tt.expectedError {
						t.Fatal(responseObj.Errors[0].Message)
					}
				}
			}
		})
	}
}
