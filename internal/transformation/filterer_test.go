package transformation_test

import (
	"testing"

	"github.com/Linkinlog/LeafListr/internal/transformation"

	"github.com/Linkinlog/LeafListr/internal/models"
)

func TestSubCategory(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		subCategoryName  string
		products         []*models.Product
		expectedProducts []*models.Product
	}{
		"empty products": {
			subCategoryName:  "test",
			products:         []*models.Product{},
			expectedProducts: []*models.Product{},
		},
		"empty subCategoryName": {
			subCategoryName: "",
			products: []*models.Product{
				{
					Id:     "test",
					Name:   "test name",
					Ctg:    "test category",
					SubCtg: "test sub category",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
				{
					Id:     "test 2",
					Name:   "test name 2",
					Ctg:    "test category 2",
					SubCtg: "test sub category 2",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
			},
			expectedProducts: []*models.Product{},
		},
		"valid subCategoryName": {
			subCategoryName: "test sub category",
			products: []*models.Product{
				{
					Id:     "test",
					Name:   "test name",
					Ctg:    "test category",
					SubCtg: "test sub category",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
				{
					Id:     "test 2",
					Name:   "test name 2",
					Ctg:    "test category 2",
					SubCtg: "test sub category 2",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
			},
			expectedProducts: []*models.Product{
				{
					Id:     "test",
					Name:   "test name",
					Ctg:    "test category",
					SubCtg: "test sub category",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filterer := transformation.NewFilterer(nil)
			filteredProducts := filterer.SubCategory(test.subCategoryName, test.products)
			if len(filteredProducts) != len(test.expectedProducts) {
				t.Errorf("expected %d products, got %d", len(test.expectedProducts), len(filteredProducts))
			}
			for i := range filteredProducts {
				if filteredProducts[i].Id != test.expectedProducts[i].Id {
					t.Errorf("expected product id %s, got %s", test.expectedProducts[i].Id, filteredProducts[i].Id)
				}
			}
		})
	}
}

func TestPrice(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		min              float64
		max              float64
		products         []*models.Product
		expectedProducts []*models.Product
	}{
		"empty products": {
			min:              0,
			max:              0,
			products:         []*models.Product{},
			expectedProducts: []*models.Product{},
		},
		"valid price range": {
			min: 10,
			max: 20,
			products: []*models.Product{
				{
					Id: "test",
					P: &models.Price{
						Total:        15,
						IsDiscounted: false,
					},
				},
				{
					Id: "test 2",
					P: &models.Price{
						Total:        25,
						IsDiscounted: false,
					},
				},
			},
			expectedProducts: []*models.Product{
				{
					Id: "test",
					P: &models.Price{
						Total:        15,
						IsDiscounted: false,
					},
				},
			},
		},
		"empty price range": {
			min: 0,
			max: 0,
			products: []*models.Product{
				{
					Id: "test",
					P: &models.Price{
						Total:        15,
						IsDiscounted: false,
					},
				},
			},
			expectedProducts: []*models.Product{},
		},
		"valid max price": {
			min: 0,
			max: 20,
			products: []*models.Product{
				{
					Id: "test",
					P: &models.Price{
						Total:           15,
						DiscountedTotal: 5,
						IsDiscounted:    true,
					},
				},
				{
					Id: "test 2",
					P: &models.Price{
						Total:           25,
						DiscountedTotal: 21,
						IsDiscounted:    true,
					},
				},
				{
					Id: "test 3",
					P: &models.Price{
						Total:           35,
						DiscountedTotal: 15,
						IsDiscounted:    true,
					},
				},
			},
			expectedProducts: []*models.Product{
				{
					Id: "test",
					P: &models.Price{
						Total:           15,
						DiscountedTotal: 5,
						IsDiscounted:    true,
					},
				},
				{
					Id: "test 3",
					P: &models.Price{
						Total:           35,
						DiscountedTotal: 15,
						IsDiscounted:    true,
					},
				},
			},
		},
		"valid min price": {
			min: 10,
			max: 0,
			products: []*models.Product{
				{
					Id: "test",
					P: &models.Price{
						Total:           15,
						DiscountedTotal: 5,
						IsDiscounted:    true,
					},
				},
				{
					Id: "test 2",
					P: &models.Price{
						Total:           25,
						DiscountedTotal: 21,
						IsDiscounted:    true,
					},
				},
				{
					Id: "test 3",
					P: &models.Price{
						Total:           35,
						DiscountedTotal: 15,
						IsDiscounted:    true,
					},
				},
			},
			expectedProducts: []*models.Product{
				{
					Id: "test 2",
					P: &models.Price{
						Total:           25,
						DiscountedTotal: 21,
						IsDiscounted:    true,
					},
				},
				{
					Id: "test 3",
					P: &models.Price{
						Total:           35,
						DiscountedTotal: 15,
						IsDiscounted:    true,
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filterer := transformation.NewFilterer(nil)
			filteredProducts := filterer.Price(test.min, test.max, test.products)
			if len(filteredProducts) != len(test.expectedProducts) {
				t.Fatalf("expected %d products, got %d", len(test.expectedProducts), len(filteredProducts))
			}
			for i := range filteredProducts {
				if filteredProducts[i].Id != test.expectedProducts[i].Id {
					t.Errorf("expected product id %s, got %s", test.expectedProducts[i].Id, filteredProducts[i].Id)
				}
			}
		})
	}
}

func TestBrand(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		brandName        string
		products         []*models.Product
		expectedProducts []*models.Product
	}{
		"empty products": {
			brandName:        "test",
			products:         []*models.Product{},
			expectedProducts: []*models.Product{},
		},
		"empty brandName": {
			brandName: "",
			products: []*models.Product{
				{
					Id:     "test",
					Name:   "test name",
					Ctg:    "test category",
					SubCtg: "test sub category",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
				{
					Id:     "test 2",
					Name:   "test name 2",
					Ctg:    "test category 2",
					SubCtg: "test sub category 2",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
			},
			expectedProducts: []*models.Product{},
		},
		"valid brandName": {
			brandName: "test brand",
			products: []*models.Product{
				{
					Id:     "test",
					Brand:  "test brand",
					Name:   "test name",
					Ctg:    "test category",
					SubCtg: "test sub category",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
				{
					Id:     "test 2",
					Brand:  "test brand 2",
					Name:   "test name 2",
					Ctg:    "test category 2",
					SubCtg: "test sub category 2",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
			},
			expectedProducts: []*models.Product{
				{
					Id:     "test",
					Brand:  "test brand",
					Name:   "test name",
					Ctg:    "test category",
					SubCtg: "test sub category",
					Images: []string{},
					C:      []*models.Cannabinoid{},
					T:      []*models.Terpene{},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filterer := transformation.NewFilterer(nil)
			filteredProducts := filterer.Brands([]string{test.brandName}, test.products)
			if len(filteredProducts) != len(test.expectedProducts) {
				t.Errorf("expected %d products, got %d", len(test.expectedProducts), len(filteredProducts))
			}
			for i := range filteredProducts {
				if filteredProducts[i].Id != test.expectedProducts[i].Id {
					t.Errorf("expected product id %s, got %s", test.expectedProducts[i].Id, filteredProducts[i].Id)
				}
			}
		})
	}
}

func TestNotBrand(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		brandName        []string
		products         []*models.Product
		expectedProducts []*models.Product
	}{
		"empty products": {
			brandName:        []string{"test"},
			products:         []*models.Product{},
			expectedProducts: []*models.Product{},
		},
		"empty brandName": {
			brandName: []string{},
			products: []*models.Product{
				{Brand: "test brand"},
				{Brand: "test brand 2"},
			},
			expectedProducts: []*models.Product{
				{Brand: "test brand"},
				{Brand: "test brand 2"},
			},
		},
		"valid brandName": {
			brandName: []string{"test brand"},
			products: []*models.Product{
				{Brand: "test brand"},
				{Brand: "test brand 2"},
			},
			expectedProducts: []*models.Product{
				{Brand: "test brand 2"},
			},
		},
		"valid brandNames": {
			brandName: []string{"test brand", "test brand 2"},
			products: []*models.Product{
				{Brand: "test brand"},
				{Brand: "test brand 2"},
			},
			expectedProducts: []*models.Product{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filterer := transformation.NewFilterer(nil)
			filteredProducts := filterer.NotBrands(test.brandName, test.products)
			if len(filteredProducts) != len(test.expectedProducts) {
				t.Errorf("expected %d products, got %d", len(test.expectedProducts), len(filteredProducts))
			}
			for i := range filteredProducts {
				if filteredProducts[i].Brand != test.expectedProducts[i].Brand {
					t.Errorf("expected product brand %s, got %s", test.expectedProducts[i].Brand, filteredProducts[i].Brand)
				}
			}
		})
	}
}

func TestExcludingTerms(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		excludedTerms    []string
		products         []*models.Product
		expectedProducts []*models.Product
	}{
		"empty products": {
			excludedTerms:    []string{"test"},
			products:         []*models.Product{},
			expectedProducts: []*models.Product{},
		},
		"valid excludedTerms": {
			excludedTerms: []string{"test", "apple"},
			products: []*models.Product{
				{
					Name: "tEst name",
				},
				{
					Name: "APPLE name",
				},
			},
			expectedProducts: []*models.Product{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filterer := transformation.NewFilterer(nil)
			filteredProducts := filterer.ExcludingTerms(test.excludedTerms, test.products)
			if len(filteredProducts) != len(test.expectedProducts) {
				t.Errorf("expected %d products, got %d", len(test.expectedProducts), len(filteredProducts))
			}
			for i := range test.expectedProducts {
				if filteredProducts[i].Name != test.expectedProducts[i].Name {
					t.Errorf("expected product name %s, got %s", test.expectedProducts[i].Name, filteredProducts[i].Name)
				}
			}
		})
	}
}
