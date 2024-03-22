package curaleaf_test

import (
	"encoding/json"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/repository"

	"github.com/Linkinlog/LeafListr/internal/client/clientfakes"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	t.Parallel()
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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "ASDF123", false)
			ls, getErr := cr.Location()
			tt.assertions(t, ls, getErr)
		})
	}
}

func TestGetLocations(t *testing.T) {
	t.Parallel()
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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "ASDF", false)
			ls, getErr := cr.Locations(0, 0)
			tt.assertions(t, ls, getErr)
		})
	}
}

func TestGetProduct(t *testing.T) {
	t.Parallel()
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
				assert.Equal(t, expectedProd.Category.Key, prod.Ctg)

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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "ASDF", false)
			ps, getErr := cr.GetProduct("foo")
			tt.assertions(t, ps, getErr)
		})
	}
}

func TestGetProducts(t *testing.T) {
	t.Parallel()
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
					assert.Equal(t, expectedProd.Category.Key, prod.Ctg)

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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "foo", false)
			ps, getErr := cr.GetProducts()
			tt.assertions(t, ps, getErr)
		})
	}
}

func TestGetProductsForCategory(t *testing.T) {
	t.Parallel()
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
					assert.Equal(t, expectedProd.Category.Key, prod.Ctg)

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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "foo", false)
			ps, getErr := cr.GetProductsForCategory("foo")
			tt.assertions(t, ps, getErr)
		})
	}
}

func TestGetCategories(t *testing.T) {
	t.Parallel()
	response := responseSample()
	bs, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		setup      func(*clientfakes.FakeClient)
		assertions func(*testing.T, []string, error)
	}{
		"valid request": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(bs, nil)
			},
			assertions: func(t *testing.T, cs []string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, cs)
				expectedCategories := response.Data.DispensaryMenu.AllFilters.Categories
				assert.Equal(t, len(expectedCategories), len(cs))

				for i, cat := range cs {
					expectedCategory := expectedCategories[i]

					assert.Equal(t, expectedCategory.Key, cat)
				}
			},
		},
		"invalid request - unexpected end of JSON input": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns(nil, nil)
			},
			assertions: func(t *testing.T, cs []string, err error) {
				assert.Equal(t, []string{}, cs)
				assert.ErrorIs(t, err, repository.InvalidJSONError)
			},
		},
		"invalid request - valid json / invalid expected response, empty products": {
			setup: func(fc *clientfakes.FakeClient) {
				fc.QueryReturns([]byte("{}"), nil)
			},
			assertions: func(t *testing.T, cs []string, err error) {
				assert.Equal(t, []string{}, cs)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "foo", false)
			cs, getErr := cr.GetCategories()
			tt.assertions(t, cs, getErr)
		})
	}
}

func TestGetTerpenes(t *testing.T) {
	t.Parallel()
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
				expectedTerps := []*curaleaf.Terpene{
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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "foo", false)
			ts, getErr := cr.GetTerpenes()
			tt.assertions(t, ts, getErr)
		})
	}
}

func TestGetCannabinoids(t *testing.T) {
	t.Parallel()
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
				expectedCannabinoids := []*curaleaf.Cannabinoid{
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
				assert.ErrorIs(t, err, repository.InvalidJSONError)
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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "foo", false)
			cs, getErr := cr.GetCannabinoids()
			tt.assertions(t, cs, getErr)
		})
	}
}

func TestGetOffers(t *testing.T) {
	t.Parallel()
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
				assert.ErrorIs(t, err, repository.InvalidJSONError)

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
		t.Run(name, func(t *testing.T) {
			c := new(clientfakes.FakeClient)
			tt.setup(c)
			cr := curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), "foo", false)
			os, getErr := cr.GetOffers()
			tt.assertions(t, os, getErr)
		})
	}
}

func productSample() curaleaf.Product {
	product := curaleaf.Product{
		Brand: curaleaf.Brand{
			Description: "Big Dispensary",
			Id:          "brand-1",
			Image: curaleaf.Image{
				URL: "https://example.com/image.png",
			},
			Name: "Super cool product",
			Slug: "brand-slug",
		},
		Category: curaleaf.Category{
			DisplayName: "Outdoors",
			Key:         "OUTDOORS",
		},
		DescriptionHtml: "",
		Effects:         nil,
		ID:              "abc121",
		Images: []curaleaf.Image{
			{
				URL: "https://example.com/image",
			},
		},
		LabResults: curaleaf.LabResult{
			Cannabinoids: []curaleaf.CannabinoidObj{
				{
					Cannabinoid: curaleaf.Cannabinoid{
						Description: "THC is a cannabinoid",
						Name:        "THC",
					},
					Unit:  "PERCENTAGE",
					Value: 4.20,
				},
				{
					Cannabinoid: curaleaf.Cannabinoid{
						Description: "CBD is a cannabinoid",
						Name:        "CBD",
					},
					Unit:  "PERCENTAGE",
					Value: 6.90,
				},
			},
			Terpenes: []curaleaf.TerpeneObj{
				{
					Terpene: curaleaf.Terpene{
						Description: "Big strong and here to sing along",
						Name:        "B-Myrcene",
					},
					UnitSymbol: "%",
					Value:      4.20,
				},
				{
					Terpene: curaleaf.Terpene{
						Description: "A great way to get around",
						Name:        "B-Pinene",
					},
					UnitSymbol: "%",
					Value:      6.90,
				},
			},
			THC: curaleaf.THC{
				Formatted: "4.20%",
				Range:     []float64{4.20},
			},
		},
		Name: "Cheesey Wheezey",
		Offers: []curaleaf.Offer{
			{
				Description: "A great deal",
				Id:          "internal-offer-1",
				Title:       "This is on the product",
			},
		},
		Strain: curaleaf.Strain{
			Key:         "wisdom",
			DisplayName: "hybrid",
		},
		Subcategory: curaleaf.Subcategory{
			Key:         "DISPOSABLES",
			DisplayName: "Disposables",
		},
		Variants: []curaleaf.Variant{
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

func offerSample() curaleaf.Offer {
	return curaleaf.Offer{
		Description: "",
		Id:          "6435a162da6e29000119f676",
		Title:       "30% off Last Call",
	}
}

func categorySample() curaleaf.Category {
	return curaleaf.Category{
		DisplayName: "Flower",
		Key:         "flower",
	}
}

func locationSample() curaleaf.Location {
	loc := curaleaf.Location{
		UniqueId:  "ASDF123",
		Name:      "Cool Location",
		Slug:      "cool-location",
		MenuTypes: []string{"MEDICAL", "RECREATIONAL"},
	}
	loc.Location.ZipCode = "12345"
	return loc
}

func responseSample() *curaleaf.Response {
	return curaleaf.NewResponse([]curaleaf.Product{productSample()}, []curaleaf.Offer{offerSample()}, []curaleaf.Category{categorySample()}, []curaleaf.Location{locationSample()})
}
