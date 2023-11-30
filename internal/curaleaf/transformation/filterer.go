package transformation

import (
	"strings"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
)

type filterer struct{}

func NewFilterer() transformation.Filterer {
	return &filterer{}
}

func (f *filterer) SubCategory(subCategoryName string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		if strings.EqualFold(product.SubCtg, subCategoryName) {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func (f *filterer) Price(min, max float64, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		for _, variant := range product.V {
			if variant.Price.DiscountedTotal > 0 && variant.Price.DiscountedTotal >= min && variant.Price.DiscountedTotal <= max {
				filteredProducts = append(filteredProducts, product)
				break
			} else if variant.Price.Total > 0 && variant.Price.Total >= min && variant.Price.Total <= max {
				filteredProducts = append(filteredProducts, product)
				break
			}
		}
	}
	return filteredProducts
}

func (f *filterer) Brands(brandNames []string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		for _, brandName := range brandNames {
			if brandName != "" && strings.EqualFold(product.Brand, brandName) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	}
	return filteredProducts
}

func (f *filterer) NotBrands(brandNames []string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		for _, brandName := range brandNames {
			if brandName != "" && !strings.EqualFold(product.Brand, brandName) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	}
	return filteredProducts
}
