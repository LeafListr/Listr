package transformation

import (
	"github.com/Linkinlog/LeafListr/internal/models"
)

type Filterer interface {
	Filter(products []*models.Product) []*models.Product
	SubCategory(subCategoryName string, products []*models.Product) []*models.Product
	Price(min, max float64, products []*models.Product) []*models.Product
	Brands(brandNames []string, products []*models.Product) []*models.Product
	NotBrands(brandNames []string, products []*models.Product) []*models.Product
	Variants(variantNames []string, products []*models.Product) []*models.Product
	IncludingTerms(terms []string, products []*models.Product) []*models.Product
	ExcludingTerms(terms []string, products []*models.Product) []*models.Product
}

type Sorter interface {
	Sort(products []*models.Product)
	PriceAsc(products []*models.Product)
	PriceDesc(products []*models.Product)
	THCAsc(products []*models.Product)
	THCDesc(products []*models.Product)
	Top3Terps(products []*models.Product, terps [3]string)
}
