package repository_test

import (
	"encoding/json"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/api/models"

	cm "github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/repository"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/translation"

	"github.com/Linkinlog/LeafListr/internal/client/clientfakes"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	allProductsResp := allProductResponseSample()
	bs, err := json.Marshal(allProductsResp)
	if err != nil {
		t.Fatal(err)
	}
	c := clientfakes.FakeClient{}
	c.QueryReturns(bs, nil)
	cr := repository.NewRepository(&c, translation.NewClientTranslator())
	ps, getErr := cr.GetProducts("foo")
	if getErr != nil {
		t.Fatal(getErr)
	}
	if ps[0].Id != allProductsResp.Data.DispensaryMenu.Products[0].ID {
		t.Fatal("Wrong id Champ")
	}

	// Assertions
	assert.NotNil(t, ps)
	expectedProducts := allProductsResp.Data.DispensaryMenu.Products
	assert.Equal(t, len(expectedProducts), len(ps))

	for i, prod := range ps {
		expectedProd := expectedProducts[i]

		assert.Equal(t, expectedProd.ID, prod.Id)
		assert.Equal(t, models.Category(expectedProd.Category.Key), models.Category(prod.Ctg))

		// Assertions for Prices (Variants)
		assert.Equal(t, len(expectedProd.Variants), len(prod.V))
		for j, v := range prod.V {
			expectedVariant := expectedProd.Variants[j]
			assert.Equal(t, expectedVariant.Option, v.Name)
			assert.Equal(t, expectedVariant.Price, v.Price.Total)
			assert.Equal(t, expectedVariant.SpecialPrice, v.Price.DiscountedTotal)
		}

		// Assertions for Cannabinoids
		assert.Equal(t, len(expectedProd.LabResults.Cannabinoids), len(prod.C))
		for j, canna := range prod.C {
			expectedCanna := expectedProd.LabResults.Cannabinoids[j]
			assert.Equal(t, expectedCanna.Cannabinoid.Name, canna.Name)
			assert.Equal(t, expectedCanna.Cannabinoid.Description, canna.Description)
			assert.Equal(t, expectedCanna.Value, canna.Value)
		}

		// Assertions for Terpenes
		assert.Equal(t, len(expectedProd.LabResults.Terpenes), len(prod.T))
		for j, terp := range prod.T {
			expectedTerp := expectedProd.LabResults.Terpenes[j]
			assert.Equal(t, expectedTerp.Terpene.Name, terp.Name)
			assert.Equal(t, expectedTerp.Terpene.Description, terp.Description)
			assert.Equal(t, expectedTerp.Value, terp.Value)
		}
	}
}

func TestGetProductsForCategory(t *testing.T) {
	allProductsResp := allProductResponseSample()
	bs, err := json.Marshal(allProductsResp)
	if err != nil {
		t.Fatal(err)
	}
	c := clientfakes.FakeClient{}
	c.QueryReturns(bs, nil)
	cr := repository.NewRepository(&c, translation.NewClientTranslator())
	ps, getErr := cr.GetProductsForCategory("foo", "category")
	if getErr != nil {
		t.Fatal(getErr)
	}
	if ps[0].Id != allProductsResp.Data.DispensaryMenu.Products[0].ID {
		t.Fatal("Wrong id Champ")
	}
}

func productSample() *cm.Product {
	product := &cm.Product{
		Brand: cm.Brand{
			Description: "Big Dispensary",
			Id:          "brand-1",
			Image: cm.Image{
				URL: "https://example.com/image.png",
			},
			Name: "Super cool product",
			Slug: "brand-slug",
		},
		Category: cm.Category{
			DisplayName: "Outdoors",
			Key:         "OUTDOORS",
		},
		DescriptionHtml: "",
		Effects:         nil,
		ID:              "abc121",
		Images: []cm.Image{
			{
				URL: "https://example.com/image",
			},
		},
		LabResults: cm.LabResult{
			Cannabinoids: []cm.CannabinoidObj{
				{
					Cannabinoid: cm.Cannabinoid{
						Description: "THC is a cannabinoid",
						Name:        "THC",
					},
					Unit:  "PERCENTAGE",
					Value: 4.20,
				},
				{
					Cannabinoid: cm.Cannabinoid{
						Description: "CBD is a cannabinoid",
						Name:        "CBD",
					},
					Unit:  "PERCENTAGE",
					Value: 6.90,
				},
			},
			Terpenes: []cm.TerpeneObj{
				{
					Terpene: cm.Terpene{
						Description: "Big strong and here to sing along",
						Name:        "B-Myrcene",
					},
					UnitSymbol: "%",
					Value:      4.20,
				},
				{
					Terpene: cm.Terpene{
						Description: "A great way to get around",
						Name:        "B-Pinene",
					},
					UnitSymbol: "%",
					Value:      6.90,
				},
			},
			THC: cm.THC{
				Formatted: "4.20%",
				Range:     []float64{4.20},
			},
		},
		Name: "Cheesey Wheezey",
		Offers: []cm.Offer{
			{
				Description: "A great deal",
				Id:          "internal-offer-1",
				Title:       "This is on the product",
			},
		},
		Strain: cm.Strain{
			Key:         "wisdom",
			DisplayName: "hybrid",
		},
		Subcategory: cm.Subcategory{
			Key:         "DISPOSABLES",
			DisplayName: "Disposables",
		},
		Variants: []cm.Variant{
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

func allProductResponseSample() *cm.Response {
	product := productSample()

	return cm.NewResponse([]*cm.Product{product}, nil, nil)
}
