package transformation

import (
	"slices"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
)

type sorter struct{}

func NewSorterer() transformation.Sorter {
	return &sorter{}
}

func (s *sorter) PriceAsc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productAPrice := productA.Price.DiscountedTotal
		productBPrice := productB.Price.DiscountedTotal

		if productAPrice == 0 {
			productAPrice = productA.Price.Total
		}

		if productBPrice == 0 {
			productBPrice = productB.Price.Total
		}

		if productAPrice < productBPrice {
			return -1
		}

		if productAPrice > productBPrice {
			return 1
		}

		return 0
	})
}

func (s *sorter) PriceDesc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productAPrice := productA.Price.DiscountedTotal
		productBPrice := productB.Price.DiscountedTotal

		if productAPrice == 0 {
			productAPrice = productA.Price.Total
		}

		if productBPrice == 0 {
			productBPrice = productB.Price.Total
		}

		if productAPrice > productBPrice {
			return -1
		}

		if productAPrice < productBPrice {
			return 1
		}

		return 0
	})
}
