package transformation

import (
	"slices"
	"strings"

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

func (s *sorter) Top3Terps(products []*models.Product, terpenes [3]string) {
	findTerpeneValue := func(p *models.Product, terpeneName string) float64 {
		for _, t := range p.T {
			if strings.EqualFold(t.Name, terpeneName) {
				return t.Value
			}
		}
		return 0
	}

	scoreProduct := func(p *models.Product) float64 {
		score := 0.0
		score += findTerpeneValue(p, terpenes[0]) * 5
		score += findTerpeneValue(p, terpenes[1]) * 2
		score += findTerpeneValue(p, terpenes[2])
		return score
	}

	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productAScore := scoreProduct(productA)
		productBScore := scoreProduct(productB)

		if productAScore > productBScore {
			return -1
		}

		if productAScore < productBScore {
			return 1
		}

		return 0
	})
}
