package transformation

import (
	"strings"

	"github.com/Linkinlog/LeafListr/internal/models"
)

type FilterParams struct {
	SubCategoryName    string
	MinPrice           float64
	MaxPrice           float64
	MinPricePerG       float64
	MaxPricePerG       float64
	IncludedBrandNames []string
	ExcludedBrandNames []string
	Variants           []string
	IncludedTerms      []string
	ExcludedTerms      []string
}

type filterer struct {
	Fp *FilterParams
}

func NewFilterer(fp *FilterParams) Filterer {
	return &filterer{
		Fp: fp,
	}
}

func (f *filterer) Filter(products []*models.Product) []*models.Product {
	filteredProducts := products
	if f.Fp.SubCategoryName != "" {
		filteredProducts = f.SubCategory(f.Fp.SubCategoryName, filteredProducts)
	}
	if f.Fp.MinPricePerG > 0 || f.Fp.MaxPricePerG > 0 {
		filteredProducts = f.PricePerG(f.Fp.MinPricePerG, f.Fp.MaxPricePerG, filteredProducts)
	}
	if f.Fp.MinPrice > 0 || f.Fp.MaxPrice > 0 {
		filteredProducts = f.Price(f.Fp.MinPrice, f.Fp.MaxPrice, filteredProducts)
	}
	if len(f.Fp.IncludedBrandNames) > 0 && f.Fp.IncludedBrandNames[0] != "" {
		filteredProducts = f.Brands(f.Fp.IncludedBrandNames, filteredProducts)
	}
	if len(f.Fp.ExcludedBrandNames) > 0 && f.Fp.ExcludedBrandNames[0] != "" {
		filteredProducts = f.NotBrands(f.Fp.ExcludedBrandNames, filteredProducts)
	}
	if len(f.Fp.Variants) > 0 && f.Fp.Variants[0] != "" {
		filteredProducts = f.Variants(f.Fp.Variants, filteredProducts)
	}
	if len(f.Fp.IncludedTerms) > 0 && f.Fp.IncludedTerms[0] != "" {
		filteredProducts = f.IncludingTerms(f.Fp.IncludedTerms, filteredProducts)
	}
	if len(f.Fp.ExcludedTerms) > 0 && f.Fp.ExcludedTerms[0] != "" {
		filteredProducts = f.ExcludingTerms(f.Fp.ExcludedTerms, filteredProducts)
	}
	return filteredProducts
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

func (f *filterer) PricePerG(min, max float64, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		price := product.P.PerGram

		if price >= min && price <= max {
			filteredProducts = append(filteredProducts, product)
		} else if max == 0 && min > 0 && price >= min {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func (f *filterer) Price(min, max float64, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		price := product.P.Total
		if product.P.IsDiscounted {
			price = product.P.DiscountedTotal
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
		valid := true
		for _, brandName := range brandNames {
			if strings.EqualFold(product.Brand, brandName) {
				valid = false
			}
		}
		if valid {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func (f *filterer) Variants(variantNames []string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		for _, variantName := range variantNames {
			if variantName != "" && strings.EqualFold(product.Weight, variantName) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	}
	return filteredProducts
}

func (f *filterer) IncludingTerms(terms []string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		for _, term := range terms {
			if term != "" && strings.Contains(strings.ToLower(product.Name), strings.ToLower(term)) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	}
	return filteredProducts
}

func (f *filterer) ExcludingTerms(terms []string, products []*models.Product) []*models.Product {
	filteredProducts := make([]*models.Product, 0)
	for _, product := range products {
		valid := true
		for _, term := range terms {
			if term != "" && strings.Contains(strings.ToLower(product.Name), strings.ToLower(term)) {
				valid = false
			}
		}
		if valid {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
