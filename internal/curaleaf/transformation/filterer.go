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
		price := product.Price.Total
		if product.Price.IsDiscounted {
			price = product.Price.DiscountedTotal
		}

		if price >= min && price <= max {
			filteredProducts = append(filteredProducts, product)
		} else if max == 0 && min > 0 && price >= min {
			filteredProducts = append(filteredProducts, product)
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

func (f *filterer) Variants(variantNames []string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		for _, variantName := range variantNames {
			if variantName != "" && strings.EqualFold(product.Variant, variantName) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	}
	return filteredProducts
}
