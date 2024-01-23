package beyondhello

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/Linkinlog/LeafListr/internal/models"
)

func (r LocationResponse) translateLocations() ([]*models.Location, error) {
	locations := make([]*models.Location, 0)
	for _, l := range r {
		id, err := extractId(l.Content.Rendered)
		if err == nil {
			loc := &models.Location{
				Id:   id,
				Name: strings.ReplaceAll(l.Title.Rendered, "Cannabis Menu", ""),
			}
			if strings.Contains(l.Title.Rendered, "Medical") {
				loc.Name = strings.ReplaceAll(loc.Name, "Medical", "")
				loc.LocationTypes = append(loc.LocationTypes, "Medical")
			} else {
				loc.Name = strings.ReplaceAll(loc.Name, "Adult-Use", "")
				loc.LocationTypes = append(loc.LocationTypes, "Recreational")
			}
			locations = append(locations, loc)
		}
	}
	return locations, nil
}

func extractId(s string) (string, error) {
	re := regexp.MustCompile(`https://api.iheartjane.com/v1/stores/(\d+)/embed.js`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		return "", errors.New("no match")
	}
	return match[1], nil
}

func (r *ProductResponse) translateProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	for _, p := range r.Hits {
		if len(p.AvailableWeights) == 0 {
			// if there are no weights, then the product is a single item
			p.AvailableWeights = append(p.AvailableWeights, "each")
		}
		for _, w := range p.AvailableWeights {
			weightId := strings.Replace(w, " ", "_", -1)
			product := models.Product{
				Id:     strconv.Itoa(p.ProductId),
				Brand:  p.Brand,
				Name:   p.Name,
				SubCtg: p.KindSubtype,
				Ctg:    p.Kind,
				Weight: weightId,
				Price:  p.translatePrice(weightId),
				C:      p.translateCannabinoids(weightId),
				T:      p.translateTerpenes(weightId),
				Images: make([]string, 0),
			}
			for _, i := range p.ProductPhotos {
				product.Images = append(product.Images, i.Id)
			}
			if product.SubCtg == "" {
				product.SubCtg = "default"
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
	case "two_gram":
		priceField = h.PriceTwoGram
		specialPriceField = h.DiscountedPriceTwoGram
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
	case "each":
		priceField = h.PriceEach
		specialPriceField = h.DiscountedPriceEach
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
