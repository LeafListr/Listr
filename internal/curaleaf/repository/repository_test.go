package repository_test

import (
	"encoding/json"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/models"
	repo "github.com/Linkinlog/LeafListr/internal/repository"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/cache"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/repository"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/translation"

	"github.com/Linkinlog/LeafListr/internal/client/clientfakes"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, *models.Location, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, loc *models.Location, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, loc)
				expectedLocs := response.Data.Dispensaries

				expectedLoc := expectedLocs[0]

				assert.Equal(t, expectedLoc.UniqueId, loc.Id)
				assert.Equal(t, expectedLoc.Location.ZipCode, loc.ZipCode)
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, loc *models.Location, err error) {
				assert.Equal(t, &models.Location{}, loc)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty product": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, loc *models.Location, err error) {
				assert.Equal(t, &models.Location{}, loc)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			ls, getErr := cr.Location("ASDF123")
			tt.assertions(t, ls, getErr)
		})
	}
}

func TestGetLocations(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Location, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, ls []*models.Location, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, ls)
				expectedLocations := response.Data.Dispensaries
				assert.Equal(t, len(expectedLocations), len(ls))

				for i, loc := range ls {
					expectedLoc := expectedLocations[i]

					assert.Equal(t, expectedLoc.UniqueId, loc.Id)
					assert.Equal(t, expectedLoc.Location.ZipCode, loc.ZipCode)
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, ls []*models.Location, err error) {
				assert.Equal(t, []*models.Location{}, ls)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, ls []*models.Location, err error) {
				assert.Equal(t, []*models.Location{}, ls)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			ls, getErr := cr.Locations(0, 0)
			tt.assertions(t, ls, getErr)
		})
	}
}

func TestGetProduct(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, *models.Product, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, prod *models.Product, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, prod)
				expectedProducts := response.Data.DispensaryMenu.Products

				expectedProd := expectedProducts[0]

				assert.Equal(t, expectedProd.ID, prod.Id)
				assert.Equal(t, models.Category(expectedProd.Category.Key), prod.Ctg)

				assert.Equal(t, len(expectedProd.LabResults.Cannabinoids), len(prod.C))
				for j, canna := range prod.C {
					expectedCanna := expectedProd.LabResults.Cannabinoids[j]
					assert.Equal(t, expectedCanna.Cannabinoid.Name, canna.Name)
					assert.Equal(t, expectedCanna.Cannabinoid.Description, canna.Description)
					assert.Equal(t, expectedCanna.Value, canna.Value)
				}

				assert.Equal(t, len(expectedProd.LabResults.Terpenes), len(prod.T))
				for j, terp := range prod.T {
					expectedTerp := expectedProd.LabResults.Terpenes[j]
					assert.Equal(t, expectedTerp.Terpene.Name, terp.Name)
					assert.Equal(t, expectedTerp.Terpene.Description, terp.Description)
					assert.Equal(t, expectedTerp.Value, terp.Value)
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, prod *models.Product, err error) {
				assert.Equal(t, &models.Product{}, prod)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty product": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, prod *models.Product, err error) {
				assert.Equal(t, &models.Product{}, prod)
				assert.Error(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			ps, getErr := cr.GetProduct("foo", "bar")
			tt.assertions(t, ps, getErr)
		})
	}
}

func TestGetProducts(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Product, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, ps []*models.Product, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, ps)
				expectedProducts := response.Data.DispensaryMenu.Products
				// assert.Equal(t, len(expectedProducts), len(ps))

				for i, prod := range ps {
					if i >= len(expectedProducts) {
						break
					}
					expectedProd := expectedProducts[i]

					// assert.Equal(t, expectedProd.ID, prod.Id)
					assert.Equal(t, models.Category(expectedProd.Category.Key), prod.Ctg)

					// for _, expectedVariant := range expectedProd.Variants {
					// 	assert.Equal(t, expectedVariant.Price, prod.Price.Total)
					// 	assert.Equal(t, expectedVariant.SpecialPrice, prod.Price.DiscountedTotal)
					// }

					assert.Equal(t, len(expectedProd.LabResults.Cannabinoids), len(prod.C))
					for j, canna := range prod.C {
						expectedCanna := expectedProd.LabResults.Cannabinoids[j]
						assert.Equal(t, expectedCanna.Cannabinoid.Name, canna.Name)
						assert.Equal(t, expectedCanna.Cannabinoid.Description, canna.Description)
						assert.Equal(t, expectedCanna.Value, canna.Value)
					}

					assert.Equal(t, len(expectedProd.LabResults.Terpenes), len(prod.T))
					for j, terp := range prod.T {
						expectedTerp := expectedProd.LabResults.Terpenes[j]
						assert.Equal(t, expectedTerp.Terpene.Name, terp.Name)
						assert.Equal(t, expectedTerp.Terpene.Description, terp.Description)
						assert.Equal(t, expectedTerp.Value, terp.Value)
					}
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, ps []*models.Product, err error) {
				assert.Equal(t, []*models.Product{}, ps)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, ps []*models.Product, err error) {
				assert.Equal(t, []*models.Product{}, ps)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			ps, getErr := cr.GetProducts("foo")
			tt.assertions(t, ps, getErr)
		})
	}
}

func TestGetProductsForCategory(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Product, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, ps []*models.Product, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, ps)
				expectedProducts := response.Data.DispensaryMenu.Products
				// assert.Equal(t, len(expectedProducts), len(ps))

				for i, prod := range ps {
					if i >= len(expectedProducts) {
						break
					}
					expectedProd := expectedProducts[i]

					//	assert.Equal(t, expectedProd.ID, prod.Id)
					assert.Equal(t, models.Category(expectedProd.Category.Key), prod.Ctg)

					//	for _, expectedVariant := range expectedProd.Variants {
					//		assert.Equal(t, expectedVariant.Price, prod.Price.Total)
					//		assert.Equal(t, expectedVariant.SpecialPrice, prod.Price.DiscountedTotal)
					//	}

					assert.Equal(t, len(expectedProd.LabResults.Cannabinoids), len(prod.C))
					for j, canna := range prod.C {
						expectedCanna := expectedProd.LabResults.Cannabinoids[j]
						assert.Equal(t, expectedCanna.Cannabinoid.Name, canna.Name)
						assert.Equal(t, expectedCanna.Cannabinoid.Description, canna.Description)
						assert.Equal(t, expectedCanna.Value, canna.Value)
					}

					assert.Equal(t, len(expectedProd.LabResults.Terpenes), len(prod.T))
					for j, terp := range prod.T {
						expectedTerp := expectedProd.LabResults.Terpenes[j]
						assert.Equal(t, expectedTerp.Terpene.Name, terp.Name)
						assert.Equal(t, expectedTerp.Terpene.Description, terp.Description)
						assert.Equal(t, expectedTerp.Value, terp.Value)
					}
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, ps []*models.Product, err error) {
				assert.Equal(t, []*models.Product{}, ps)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, ps []*models.Product, err error) {
				assert.Equal(t, []*models.Product{}, ps)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			ps, getErr := cr.GetProductsForCategory("foo", "bar")
			tt.assertions(t, ps, getErr)
		})
	}
}

func TestGetCategories(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Category, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, cs []*models.Category, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, cs)
				expectedCategories := response.Data.DispensaryMenu.AllFilters.Categories
				assert.Equal(t, len(expectedCategories), len(cs))

				for i, cat := range cs {
					expectedCategory := expectedCategories[i]

					assert.Equal(t, models.Category(expectedCategory.Key), *cat)
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, cs []*models.Category, err error) {
				assert.Equal(t, []*models.Category{}, cs)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, cs []*models.Category, err error) {
				assert.Equal(t, []*models.Category{}, cs)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			cs, getErr := cr.GetCategories("foo")
			tt.assertions(t, cs, getErr)
		})
	}
}

func TestGetTerpenes(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Terpene, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, ts []*models.Terpene, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, ts)
				expectedTerps := []*client.Terpene{
					{
						Description: "Big strong and here to sing along",
						Name:        "B-Myrcene",
					},
					{
						Description: "A great way to get around",
						Name:        "B-Pinene",
					},
				}
				assert.Equal(t, len(expectedTerps), len(ts))
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, ts []*models.Terpene, err error) {
				assert.Equal(t, []*models.Terpene{}, ts)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, ts []*models.Terpene, err error) {
				assert.Equal(t, []*models.Terpene{}, ts)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			ts, getErr := cr.GetTerpenes("foo")
			tt.assertions(t, ts, getErr)
		})
	}
}

func TestGetCannabinoids(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Cannabinoid, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, cs []*models.Cannabinoid, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, cs)
				expectedCannabinoids := []*client.Cannabinoid{
					{
						Description: "THC is a cannabinoid",
						Name:        "THC",
					},
					{
						Description: "CBD is a cannabinoid",
						Name:        "CBD",
					},
				}
				assert.Equal(t, len(expectedCannabinoids), len(cs))
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, cs []*models.Cannabinoid, err error) {
				assert.Equal(t, []*models.Cannabinoid{}, cs)
				assert.ErrorIs(t, err, repo.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, cs []*models.Cannabinoid, err error) {
				assert.Equal(t, []*models.Cannabinoid{}, cs)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			cs, getErr := cr.GetCannabinoids("foo")
			tt.assertions(t, cs, getErr)
		})
	}
}

func TestGetOffers(t *testing.T) {
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []*models.Offer, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, os []*models.Offer, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, os)
				expectedOffers := response.Data.DispensaryMenu.Offers
				assert.Equal(t, len(expectedOffers), len(os))

				for i, offer := range os {
					expectedOffer := expectedOffers[i]

					assert.Equal(t, expectedOffer.Id, offer.Id)
					assert.Equal(t, expectedOffer.Title, offer.Description)
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, os []*models.Offer, err error) {
				assert.ErrorIs(t, err, repo.InvalidJSONError)

				assert.Equal(t, []*models.Offer{}, os)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, os []*models.Offer, err error) {
				assert.Equal(t, []*models.Offer{}, os)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		tt := tt // shadowed so we can run in parallel
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := repository.NewRepository(c, translation.NewClientTranslator(), cache.NewCache())
			os, getErr := cr.GetOffers("foo")
			tt.assertions(t, os, getErr)
		})
	}
}

func productSample() client.Product {
	product := client.Product{
		Brand: client.Brand{
			Description: "Big Dispensary",
			Id:          "brand-1",
			Image: client.Image{
				URL: "https://example.com/image.png",
			},
			Name: "Super cool product",
			Slug: "brand-slug",
		},
		Category: client.Category{
			DisplayName: "Outdoors",
			Key:         "OUTDOORS",
		},
		DescriptionHtml: "",
		Effects:         nil,
		ID:              "abc121",
		Images: []client.Image{
			{
				URL: "https://example.com/image",
			},
		},
		LabResults: client.LabResult{
			Cannabinoids: []client.CannabinoidObj{
				{
					Cannabinoid: client.Cannabinoid{
						Description: "THC is a cannabinoid",
						Name:        "THC",
					},
					Unit:  "PERCENTAGE",
					Value: 4.20,
				},
				{
					Cannabinoid: client.Cannabinoid{
						Description: "CBD is a cannabinoid",
						Name:        "CBD",
					},
					Unit:  "PERCENTAGE",
					Value: 6.90,
				},
			},
			Terpenes: []client.TerpeneObj{
				{
					Terpene: client.Terpene{
						Description: "Big strong and here to sing along",
						Name:        "B-Myrcene",
					},
					UnitSymbol: "%",
					Value:      4.20,
				},
				{
					Terpene: client.Terpene{
						Description: "A great way to get around",
						Name:        "B-Pinene",
					},
					UnitSymbol: "%",
					Value:      6.90,
				},
			},
			THC: client.THC{
				Formatted: "4.20%",
				Range:     []float64{4.20},
			},
		},
		Name: "Cheesey Wheezey",
		Offers: []client.Offer{
			{
				Description: "A great deal",
				Id:          "internal-offer-1",
				Title:       "This is on the product",
			},
		},
		Strain: client.Strain{
			Key:         "wisdom",
			DisplayName: "hybrid",
		},
		Subcategory: client.Subcategory{
			Key:         "DISPOSABLES",
			DisplayName: "Disposables",
		},
		Variants: []client.Variant{
			{
				Id:           "variant-1",
				IsSpecial:    false,
				Option:       "Size 1",
				Price:        7.99,
				Quantity:     98,
				SpecialPrice: 3.99,
			},
			{
				Id:        "variant-2",
				IsSpecial: false,
				Option:    "Size 2",
				Price:     8.99,
				Quantity:  69,
			},
		},
		CardDescription: "I have a thc percentage as well sometimes 25.5%",
	}
	return product
}

func offerSample() client.Offer {
	return client.Offer{
		Description: "",
		Id:          "6435a162da6e29000119f676",
		Title:       "30% off Last Call",
	}
}

func categorySample() client.Category {
	return client.Category{
		DisplayName: "Flower",
		Key:         "FLOWER",
	}
}

func locationSample() client.Location {
	loc := client.Location{
		UniqueId: "ASDF123",
		Name:     "Cool Location",
		Slug:     "cool-location",
	}
	loc.Location.ZipCode = "12345"
	return loc
}

func responseSample() *client.Response {
	return client.NewResponse([]client.Product{productSample()}, []client.Offer{offerSample()}, []client.Category{categorySample()}, []client.Location{locationSample()})
}
