package curaleaf_test

import (
	"errors"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf"

	"github.com/Linkinlog/LeafListr/internal/factory/factoryfakes"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/repository/repositoryfakes"
	"github.com/stretchr/testify/assert"
)

func TestManager_Location(t *testing.T) {
	tests := map[string]struct {
		dispensary string
		menuId     string
		wantErr    bool
	}{
		"Valid Location": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			wantErr:    false,
		},
		"Invalid Location": {
			dispensary: "snoraleaf",
			menuId:     "zzz",
			wantErr:    true,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			m := curaleaf.NewWorkflow()
			location, err := m.Location(tt.dispensary, tt.menuId, "MEDICAL")
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, location)
			}
		})
	}
}

func TestManager_Locations(t *testing.T) {
	tests := map[string]struct {
		dispensary string
		wantErr    bool
	}{
		"Valid Location": {
			dispensary: "curaleaf",
			wantErr:    false,
		},
		"Invalid Location": {
			dispensary: "snoraleaf",
			wantErr:    true,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			m := curaleaf.NewWorkflow()
			location, err := m.Locations(tt.dispensary, "MEDICAL")
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, location)
			}
		})
	}
}

func TestManager_Product(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		productId     string
		expectedError error
		wantErr       bool
	}{
		"Valid Product": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			productId:  "654e3feeba407e0001ab1ab0",
			wantErr:    false,
		},
		"Invalid Dispensary": {
			dispensary:    "invalid",
			menuId:        "LMR070",
			menuType:      "MEDICAL",
			productId:     "Rizz-Machine",
			wantErr:       true,
			expectedError: errors.New("invalid dispensary"),
		},
		"Invalid MenuId": {
			dispensary:    "curaleaf",
			menuId:        "Invalid",
			menuType:      "MEDICAL",
			productId:     "Rizz-Machine",
			wantErr:       true,
			expectedError: errors.New("invalid menu"),
		},
		"Invalid ProductId": {
			dispensary:    "curaleaf",
			menuId:        "LMR070",
			menuType:      "MEDICAL",
			productId:     "Invalid",
			wantErr:       true,
			expectedError: errors.New("invalid product"),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p := &models.Product{
				Id:   tt.productId,
				Name: "Rizzer",
				Ctg:  "",
				C:    nil,
				T:    nil,
			}

			r := &repositoryfakes.FakeRepository{}
			r.GetProductReturns(p, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			product, err := m.Product(tt.dispensary, tt.menuId, tt.menuType, tt.productId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, product.Id, tt.productId)
			}
		})
	}
}

func TestManager_Products(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		expectedError error
		wantErr       bool
	}{
		"Valid Products Request": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			products := []*models.Product{
				{Id: "Product1", Name: "Product One"},
				{Id: "Product2", Name: "Product Two"},
			}

			r := &repositoryfakes.FakeRepository{}
			r.GetProductsReturns(products, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			result, err := m.Products(tt.dispensary, tt.menuId, tt.menuType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, products, result)
			}
		})
	}
}

func TestManager_ProductsForCategory(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		category      string
		expectedError error
		wantErr       bool
	}{
		"Valid Category Request": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			category:   "Edibles",
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			products := []*models.Product{
				{Id: "Product1", Name: "Product One", Ctg: "Edibles"},
				{Id: "Product2", Name: "Product Two", Ctg: "Edibles"},
			}

			r := &repositoryfakes.FakeRepository{}
			r.GetProductsForCategoryReturns(products, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			result, err := m.ProductsForCategory(tt.dispensary, tt.menuId, tt.menuType, models.Category(tt.category))
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, products, result)
			}
		})
	}
}

func TestManager_Categories(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		expectedError error
		wantErr       bool
	}{
		"Valid Categories Request": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			edible := models.Category("Edibles")
			flower := models.Category("Flower")

			categories := []*models.Category{&edible, &flower}

			r := &repositoryfakes.FakeRepository{}
			r.GetCategoriesReturns(categories, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			result, err := m.Categories(tt.dispensary, tt.menuId, tt.menuType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, categories, result)
			}
		})
	}
}

func TestManager_Terpenes(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		expectedError error
		wantErr       bool
	}{
		"Valid Terpenes Request": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			terpenes := []*models.Terpene{
				{Name: "Myrcene"},
				{Name: "Limonene"},
			}

			r := &repositoryfakes.FakeRepository{}
			r.GetTerpenesReturns(terpenes, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			result, err := m.Terpenes(tt.dispensary, tt.menuId, tt.menuType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, terpenes, result)
			}
		})
	}
}

func TestManager_Cannabinoids(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		expectedError error
		wantErr       bool
	}{
		"Valid Cannabinoids Request": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cannabinoids := []*models.Cannabinoid{
				{Name: "CBD"},
				{Name: "THC"},
			}

			r := &repositoryfakes.FakeRepository{}
			r.GetCannabinoidsReturns(cannabinoids, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			result, err := m.Cannabinoids(tt.dispensary, tt.menuId, tt.menuType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, cannabinoids, result)
			}
		})
	}
}

func TestManager_Offers(t *testing.T) {
	tests := map[string]struct {
		dispensary    string
		menuId        string
		menuType      string
		expectedError error
		wantErr       bool
	}{
		"Valid Offers Request": {
			dispensary: "curaleaf",
			menuId:     "LMR070",
			menuType:   "MEDICAL",
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			offers := []*models.Offer{
				{Id: "Offer1", Description: "Discount 1"},
				{Id: "Offer2", Description: "Discount 2"},
			}

			r := &repositoryfakes.FakeRepository{}
			r.GetOffersReturns(offers, tt.expectedError)

			f := &factoryfakes.FakeRepositoryFactory{}
			f.FindByDispensaryMenuReturns(r, nil)

			m := curaleaf.Workflow{F: f}

			result, err := m.Offers(tt.dispensary, tt.menuId, tt.menuType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, offers, result)
			}
		})
	}
}
