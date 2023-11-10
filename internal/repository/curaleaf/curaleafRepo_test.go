package curaleaf_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Linkinlog/LeafList/internal/client/clientfakes"
	cm "github.com/Linkinlog/LeafList/internal/client/curaleaf"
	"github.com/Linkinlog/LeafList/internal/models"
	"github.com/Linkinlog/LeafList/internal/repository/curaleaf"
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
	cr := curaleaf.NewRepository(&c)
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
		assert.Equal(t, models.Category(expectedProd.Category.Key), prod.Ctg)

		// Assertions for Prices (Variants)
		assert.Equal(t, len(expectedProd.Variants), len(prod.P))
		for j, price := range prod.P {
			expectedVariant := expectedProd.Variants[j]
			assert.Equal(t, expectedVariant.Option, price.Variant)
			assert.Equal(t, expectedVariant.Price, price.Total)
			assert.Equal(t, expectedVariant.SpecialPrice, price.DiscountedTotal)
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
	cr := curaleaf.NewRepository(&c)
	ps, getErr := cr.GetProductsForCategory("foo", "category")
	if getErr != nil {
		t.Fatal(getErr)
	}
	if ps[0].Id != allProductsResp.Data.DispensaryMenu.Products[0].ID {
		t.Fatal("Wrong id Champ")
	}
}

func TestRepository_GetCategories(t *testing.T) {
	vapes := models.Category("VAPORIZERS")
	flower := models.Category("FLOWER")
	type args struct {
		menuId string
	}
	tests := map[string]struct {
		args    args
		want    []*models.Category
		wantErr assert.ErrorAssertionFunc
	}{
		"all categories": {
			args: args{menuId: "abc123"},
			want: []*models.Category{&vapes, &flower},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			categoriesResponse := allCategoriesResponseSample()
			bs, err := json.Marshal(categoriesResponse)
			if err != nil {
				t.Fatal(err)
			}
			c := clientfakes.FakeClient{}
			c.QueryReturns(bs, nil)
			r := curaleaf.NewRepository(&c)
			got, getErr := r.GetCategories(tt.args.menuId)
			if !tt.wantErr(t, getErr, fmt.Sprintf("GetCategories(%v)", tt.args.menuId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetCategories(%v)", tt.args.menuId)
		})
	}
}

func TestRepository_GetTerpenes(t *testing.T) {
	type args struct {
		menuId string
	}
	tests := map[string]struct {
		args    args
		want    []*models.Terpene
		wantErr assert.ErrorAssertionFunc
	}{
		"all terpenes": {
			args: args{menuId: "foo"},
			want: []*models.Terpene{
				{
					Name:        "B-Myrcene",
					Description: "Big strong and here to sing along",
					Value:       4.20,
				},
				{
					Name:        "B-Pinene",
					Description: "A great way to get around",
					Value:       6.9,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			allProductsResp := allProductResponseSample()
			bs, err := json.Marshal(allProductsResp)
			if err != nil {
				t.Fatal(err)
			}
			c := clientfakes.FakeClient{}
			c.QueryReturns(bs, nil)
			r := curaleaf.NewRepository(&c)
			got, getErr := r.GetTerpenes(tt.args.menuId)
			if !tt.wantErr(t, getErr, fmt.Sprintf("GetTerpenes(%v)", tt.args.menuId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetTerpenes(%v)", tt.args.menuId)
		})
	}
}

func TestRepository_GetOffers(t *testing.T) {
	type args struct {
		menuId string
	}
	tests := map[string]struct {
		args    args
		want    []*models.Offer
		wantErr assert.ErrorAssertionFunc
	}{
		"all offers": {
			args: args{menuId: "foo"},
			want: []*models.Offer{
				{
					Id:          "cool-offer-5",
					Description: "100% off",
				},
				{
					Id:          "cool-bad-5",
					Description: "1000% extra",
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			allOffersResp := allOffersResponseSample()
			bs, err := json.Marshal(allOffersResp)
			if err != nil {
				t.Fatal(err)
			}
			c := clientfakes.FakeClient{}
			c.QueryReturns(bs, nil)
			r := curaleaf.NewRepository(&c)
			got, getErr := r.GetOffers(tt.args.menuId)
			if !tt.wantErr(t, getErr, fmt.Sprintf("GetOffers(%v)", tt.args.menuId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetOffers(%v)", tt.args.menuId)
		})
	}
}

func TestRepository_GetProduct(t *testing.T) {
	type args struct {
		menuId    string
		productId string
	}
	tests := map[string]struct {
		args    args
		want    *models.Product
		wantErr assert.ErrorAssertionFunc
	}{
		"product": {
			args: args{
				menuId:    "foo",
				productId: "abc121",
			},
			want: &models.Product{
				Id:   "abc121",
				Name: "Cheesey Wheezey",
				P: []*models.Price{
					{
						Variant:         "Size 1",
						Total:           7.99,
						DiscountedTotal: 3.99,
					},
					{
						Variant:         "Size 2",
						Total:           8.99,
						DiscountedTotal: 0,
					},
				},
				C: []*models.Cannabinoid{
					{
						Name:        "THC",
						Description: "THC is a cannabinoid",
						Value:       4.20,
					},
					{
						Name:        "CBD",
						Description: "CBD is a cannabinoid",
						Value:       6.90,
					},
				},
				T: []*models.Terpene{
					{
						Name:        "B-Myrcene",
						Description: "Big strong and here to sing along",
						Value:       4.20,
					},
					{
						Name:        "B-Pinene",
						Description: "A great way to get around",
						Value:       6.90,
					},
				},
				Ctg: "OUTDOORS",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			prodResp := productResponseSample()
			bs, err := json.Marshal(prodResp)
			if err != nil {
				t.Fatal(err)
			}
			c := clientfakes.FakeClient{}
			c.QueryReturns(bs, nil)
			r := curaleaf.NewRepository(&c)
			got, getErr := r.GetProduct(tt.args.menuId, tt.args.productId)
			if !tt.wantErr(t, getErr, fmt.Sprintf("GetProduct(%v, %v)", tt.args.menuId, tt.args.productId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetProduct(%v, %v)", tt.args.menuId, tt.args.productId)
		})
	}
}

func allOffersResponseSample() *cm.Response {
	offers := []*cm.Offer{
		{
			Description: "Super cool offer",
			Id:          "cool-offer-5",
			Title:       "100% off",
		},
		{
			Description: "Super bad offer",
			Id:          "cool-bad-5",
			Title:       "1000% extra",
		},
	}
	return cm.NewResponse(nil, offers, nil)
}

func allCategoriesResponseSample() *cm.Response {
	cats := []*cm.Category{
		{
			DisplayName: "Vaporizers",
			Key:         "VAPORIZERS",
		},
		{
			DisplayName: "Flower",
			Key:         "FLOWER",
		},
	}

	return cm.NewResponse(nil, nil, cats)
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

func productResponseSample() *cm.Response {
	product := productSample()

	return cm.NewResponse([]*cm.Product{product}, nil, nil)
}

func allProductResponseSample() *cm.Response {
	product := productSample()

	return cm.NewResponse([]*cm.Product{product}, nil, nil)
}
