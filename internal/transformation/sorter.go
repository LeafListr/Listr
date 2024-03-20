package transformation

import (
	"slices"
	"strings"

	"github.com/Linkinlog/LeafListr/internal/models"
)

type SortParams struct {
	Top3Terps     [3]string
	Top3Asc       bool
	PriceAsc      bool
	PriceDesc     bool
	THCAsc        bool
	THCDesc       bool
	TerpAsc       bool
	TerpDesc      bool
	GramPriceAsc  bool
	GramPriceDesc bool
}

type sorter struct {
	Sp *SortParams
}

func NewSorterer(sp *SortParams) Sorter {
	return &sorter{
		Sp: sp,
	}
}

func (s *sorter) Sort(products []*models.Product) {
	if s.Sp.Top3Terps != [3]string{} {
		s.Top3Terps(products, s.Sp.Top3Terps, s.Sp.Top3Asc)
	}
	if s.Sp.PriceAsc {
		s.PriceAsc(products)
	}
	if s.Sp.PriceDesc {
		s.PriceDesc(products)
	}
	if s.Sp.THCAsc {
		s.THCAsc(products)
	}
	if s.Sp.THCDesc {
		s.THCDesc(products)
	}
	if s.Sp.TerpAsc {
		s.TerpAsc(products)
	}
	if s.Sp.TerpDesc {
		s.TerpDesc(products)
	}
	if s.Sp.GramPriceAsc {
		s.GramPriceAsc(products)
	}
	if s.Sp.GramPriceDesc {
		s.GramPriceDesc(products)
	}
}

func (s *sorter) PriceAsc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productAPrice := productA.P.DiscountedTotal
		productBPrice := productB.P.DiscountedTotal

		if productAPrice == 0 {
			productAPrice = productA.P.Total
		}

		if productBPrice == 0 {
			productBPrice = productB.P.Total
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
		productAPrice := productA.P.DiscountedTotal
		productBPrice := productB.P.DiscountedTotal

		if productAPrice == 0 {
			productAPrice = productA.P.Total
		}

		if productBPrice == 0 {
			productBPrice = productB.P.Total
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

func (s *sorter) THCAsc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productATHC := productA.THC()
		productBTHC := productB.THC()

		if productATHC < productBTHC {
			return -1
		}

		if productATHC > productBTHC {
			return 1
		}

		return 0
	})
}

func (s *sorter) THCDesc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productATHC := productA.THC()
		productBTHC := productB.THC()

		if productATHC > productBTHC {
			return -1
		}

		if productATHC < productBTHC {
			return 1
		}

		return 0
	})
}

func (s *sorter) TerpAsc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productATerp := productA.TotalTerps()
		productBTerp := productB.TotalTerps()

		if productATerp < productBTerp {
			return -1
		}

		if productATerp > productBTerp {
			return 1
		}

		return 0
	})
}

func (s *sorter) TerpDesc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productATerp := productA.TotalTerps()
		productBTerp := productB.TotalTerps()

		if productATerp > productBTerp {
			return -1
		}

		if productATerp < productBTerp {
			return 1
		}

		return 0
	})
}

func (s *sorter) GramPriceAsc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productAGPrice := productA.PricePerGram()
		productBGPrice := productB.PricePerGram()

		if productAGPrice < productBGPrice {
			return -1
		}

		if productAGPrice > productBGPrice {
			return 1
		}

		return 0
	})
}

func (s *sorter) GramPriceDesc(products []*models.Product) {
	slices.SortStableFunc(products, func(productA, productB *models.Product) int {
		productAGPrice := productA.PricePerGram()
		productBGPrice := productB.PricePerGram()

		if productAGPrice > productBGPrice {
			return -1
		}

		if productAGPrice < productBGPrice {
			return 1
		}

		return 0
	})
}

func (s *sorter) Top3Terps(products []*models.Product, terpenes [3]string, asc bool) {
	// TODO simplify this
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
			if asc {
				return 1
			}
			return -1
		}

		if productAScore < productBScore {
			if asc {
				return -1
			}
			return 1
		}

		return 0
	})
}
