package curaleaf_test

import (
	"encoding/json"
	"fmt"
	"github.com/Linkinlog/LeafList/internal/client/clientfakes"
	cm "github.com/Linkinlog/LeafList/internal/client/curaleaf"
	"github.com/Linkinlog/LeafList/internal/repository/curaleaf"
	"github.com/Linkinlog/LeafList/internal/repository/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProducts(t *testing.T) {
	allProductsResp := allProductsResponseSample()
	bs, err := json.Marshal(allProductsResp)
	if err != nil {
		t.Fatal(err)
	}
	c := clientfakes.FakeClient{}
	c.QueryReturns(bs, nil)
	cr := curaleaf.Repository{
		C: &c,
	}
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
		assert.Equal(t, models.Category(expectedProd.Category.DisplayName), prod.Ctg)

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
	allProductsResp := allProductsResponseSample()
	bs, err := json.Marshal(allProductsResp)
	if err != nil {
		t.Fatal(err)
	}
	c := clientfakes.FakeClient{}
	c.QueryReturns(bs, nil)
	cr := curaleaf.Repository{
		C: &c,
	}
	ps, getErr := cr.GetProductsForCategory("foo", "category")
	if getErr != nil {
		t.Fatal(getErr)
	}
	if ps[0].Id != allProductsResp.Data.DispensaryMenu.Products[0].ID {
		t.Fatal("Wrong id Champ")
	}
}

func TestRepository_GetCategories(t *testing.T) {
	type args struct {
		menuId string
	}
	tests := map[string]struct {
		args    args
		want    []models.Category
		wantErr assert.ErrorAssertionFunc
	}{
		"all categories": {
			args: args{menuId: "abc123"},
			want: []models.Category{"VAPORIZERS", "FLOWER"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					return false
				}
				return true
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
			r := &curaleaf.Repository{
				C: &c,
			}
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
			allProductsResp := allProductsResponseSample()
			bs, err := json.Marshal(allProductsResp)
			if err != nil {
				t.Fatal(err)
			}
			c := clientfakes.FakeClient{}
			c.QueryReturns(bs, nil)
			r := &curaleaf.Repository{
				C: &c,
			}
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
			r := &curaleaf.Repository{
				C: &c,
			}
			got, getErr := r.GetOffers(tt.args.menuId)
			if !tt.wantErr(t, getErr, fmt.Sprintf("GetOffers(%v)", tt.args.menuId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetOffers(%v)", tt.args.menuId)
		})
	}
}

func allOffersResponseSample() cm.AllOffersResponse {
	offers := []cm.Offer{
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
	return cm.NewAllOffersResponse(offers)
}

func allCategoriesResponseSample() cm.AllCategoriesResponse {
	cats := []cm.Category{
		{
			DisplayName: "Vaporizers",
			Key:         "VAPORIZERS",
		},
		{
			DisplayName: "Flower",
			Key:         "FLOWER",
		},
	}

	return cm.NewAllCategoriesResponse(cats)
}

func allProductsResponseSample() cm.AllProductsResponse {
	return cm.AllProductsResponse{
		Data: struct {
			DispensaryMenu struct {
				Offers   []cm.Offer   `json:"offers"`
				Products []cm.Product `json:"products"`
			} `json:"dispensaryMenu"`
		}{
			DispensaryMenu: struct {
				Offers   []cm.Offer   `json:"offers"`
				Products []cm.Product `json:"products"`
			}{
				Offers: []cm.Offer{
					{
						Id:    "offer 1",
						Title: "18% off",
					},
					{
						Id:    "offer 2",
						Title: "38% off",
					},
				},
				Products: []cm.Product{
					{
						Brand: cm.Brand{
							Description: "Big Brand",
							Id:          "brand-1",
							Image: cm.Image{
								URL: "https://example.com/image.png",
							},
							Name: "Super cool product",
							Slug: "brand-slug",
						},
						Category: cm.Category{
							DisplayName: "Outdoors",
							Key:         "abc121",
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
								Id:        "variant-1",
								IsSpecial: false,
								Option:    "Size",
								Price:     7.99,
								Quantity:  98,
							},
							{
								Id:        "variant-2",
								IsSpecial: false,
								Option:    "Size",
								Price:     8.99,
								Quantity:  69,
							},
						},
						CardDescription: "I have a thc percentage as well sometimes 25.5%",
					},
				},
			},
		},
	}
}
