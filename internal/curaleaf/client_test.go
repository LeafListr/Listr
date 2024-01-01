package curaleaf_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf"
)

func TestSend(t *testing.T) {
	menuType := "MEDICAL"
	gbgId := "LMR124"
	tests := map[string]struct {
		ctx           context.Context
		query         string
		endpoint      curaleaf.Endpoint
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
			endpoint: curaleaf.GqlEndpoint,
			query:    curaleaf.AllProductQuery(gbgId, menuType),
			response: &curaleaf.Response{},
		},
		"curaleaf - all vaporizer products": {
			ctx:      context.Background(),
			endpoint: curaleaf.GqlEndpoint,
			query:    curaleaf.AllProductForCategoryQuery(gbgId, menuType, "VAPORIZERS"),
			response: &curaleaf.Response{},
		},
		"curaleaf - all offers": {
			ctx:      context.Background(),
			endpoint: curaleaf.GqlEndpoint,
			query:    curaleaf.AllOffersQuery(gbgId, menuType),
			response: &curaleaf.Response{},
		},
		"curaleaf - all categories": {
			ctx:      context.Background(),
			endpoint: curaleaf.GqlEndpoint,
			query:    curaleaf.AllCategoriesQuery(gbgId, menuType),
			response: &curaleaf.Response{},
		},
		"curaleaf - all locations": {
			ctx:      context.Background(),
			endpoint: curaleaf.GqlEndpoint,
			query:    curaleaf.AllLocationsQuery(-79.5389, 40.3015),
			response: &curaleaf.Response{},
		},
		"curaleaf - invalid location": {
			ctx:           context.Background(),
			endpoint:      curaleaf.GqlEndpoint,
			query:         curaleaf.AllProductQuery("foo", menuType),
			response:      &curaleaf.Response{},
			expectedError: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := curaleaf.NewHTTPClient(tt.endpoint)
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
					if responseObj.ErrorObj.Errors != nil && !tt.expectedError {
						t.Fatal(responseObj.Errors[0].Message)
					}
				}
			}
		})
	}
}
