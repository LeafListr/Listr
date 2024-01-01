package beyondhello

import (
	"strconv"

	"github.com/Linkinlog/LeafListr/internal/models"
)

func (r *Response) translateProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	for _, p := range r.Hits {
		for _, w := range p.InventoryPotencies {
			product := models.Product{
				Id:     strconv.Itoa(p.ProductId),
				Brand:  p.Brand,
				Name:   p.Name,
				SubCtg: p.KindSubtype,
				Ctg:    models.Category(p.Kind),
				Weight: w.PriceId,
				Price:  p.translatePrice(w.PriceId),
				C:      p.translateCannabinoids(w.PriceId),
				T:      p.translateTerpenes(w.PriceId),
			}
			for _, i := range p.Photos {
				product.Images = append(product.Images, i.Urls.Medium)
			}
			products = append(products, &product)
		}
	}
	return products, nil
}

func (h *hit) translateCannabinoids(weight string) []*models.Cannabinoid {
	cannabinoids := make([]*models.Cannabinoid, 0)
	for _, l := range h.LabResults {
		if l.PriceId == weight {
			for _, lr := range l.LabResults {
				switch lr.CompoundName {
				case "THC", "THCA", "CBD", "CBDA", "CBG", "CBGA", "CBC", "CBCA", "THCV", "CBDV", "CBN", "CBNA", "CBT", "CBTA", "THCA-A", "THCV-A", "CBDV-A", "CBGA-A", "CBG-A", "CBC-A", "CBN-A", "CBDA-A", "CBD-A", "CBT-A", "CBTA-A":
					cannabinoids = append(cannabinoids, &models.Cannabinoid{
						Name:  lr.CompoundName,
						Value: lr.Value,
					})
				}
			}
		}
	}
	return cannabinoids
}

func (h *hit) translateTerpenes(weight string) []*models.Terpene {
	terpenes := make([]*models.Terpene, 0)
	for _, l := range h.LabResults {
		if l.PriceId == weight {
			for _, lr := range l.LabResults {
				switch lr.CompoundName {
				case "THC", "THCA", "CBD", "CBDA", "CBG", "CBGA", "CBC", "CBCA", "THCV", "CBDV", "CBN", "CBNA", "CBT", "CBTA", "THCA-A", "THCV-A", "CBDV-A", "CBGA-A", "CBG-A", "CBC-A", "CBN-A", "CBDA-A", "CBD-A", "CBT-A", "CBTA-A":
					continue
				default:
					terpenes = append(terpenes, &models.Terpene{
						Name:  lr.CompoundName,
						Value: lr.Value,
					})
				}
			}
		}
	}
	return terpenes
}

func (h *hit) translatePrice(weight string) *models.Price {
	var priceField string
	var specialPriceField string

	switch weight {
	case "half_gram":
		priceField = h.PriceHalfGram
		specialPriceField = h.DiscountedPriceHalfGram
	case "gram":
		priceField = h.PriceGram
		specialPriceField = h.DiscountedPriceGram
	case "eighth_ounce":
		priceField = h.PriceEighthOunce
		specialPriceField = h.DiscountedPriceEighthOunce
	case "quarter_ounce":
		priceField = h.PriceQuarterOunce
		specialPriceField = h.DiscountedPriceQuarterOunce
	case "half_ounce":
		priceField = h.PriceHalfOunce
		specialPriceField = h.DiscountedPriceHalfOunce
	case "ounce":
		priceField = h.PriceOunce
		specialPriceField = h.DiscountedPriceOunce
	case "two_gram":
		priceField = h.PriceTwoGram
		specialPriceField = h.DiscountedPriceTwoGram
	}

	price := &models.Price{}

	if priceField != "" {
		price.Total, _ = strconv.ParseFloat(priceField, 64)
	}

	if specialPriceField != "" {
		price.DiscountedTotal, _ = strconv.ParseFloat(specialPriceField, 64)
		price.IsDiscounted = true
	}

	return price
}
