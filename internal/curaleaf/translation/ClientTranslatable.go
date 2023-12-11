package translation

import (
	"log/slog"
	"strings"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
)

type ClientTranslator struct{}

func NewClientTranslator() translation.ClientTranslatable {
	return &ClientTranslator{}
}

func (cT *ClientTranslator) TranslateClientLocation(l client.Location) *models.Location {
	return &models.Location{
		Id:      l.UniqueId,
		Name:    l.Name,
		Address: l.Location.Address,
		City:    l.Location.City,
		State:   l.Location.State,
		ZipCode: l.Location.ZipCode,
	}
}

func (cT *ClientTranslator) TranslateClientLocations(ls []client.Location) []*models.Location {
	locations := make([]*models.Location, 0)
	for _, l := range ls {
		locations = append(locations, cT.TranslateClientLocation(l))
	}
	return locations
}

func (cT *ClientTranslator) TranslateClientProducts(ps []client.Product) []*models.Product {
	products := make([]*models.Product, 0)

	for _, p := range ps {
		if len(p.Variants) == 0 {
			baseProduct := cT.TranslateClientProduct(p)
			products = append(products, baseProduct)
			slog.Debug("no variants for product", slog.String("product name", p.Name))
			continue
		}

		for _, v := range p.Variants {
			variantProduct := cT.TranslateClientProduct(p)
			variantProduct.Id = v.Id
			variantProduct.Price = &models.Price{
				Total:           v.Price,
				DiscountedTotal: v.SpecialPrice,
				IsDiscounted:    v.IsSpecial,
			}
			variantProduct.Variant = v.Option

			products = append(products, variantProduct)
		}
	}

	return products
}

func (cT *ClientTranslator) TranslateClientProduct(p client.Product) *models.Product {
	product := &models.Product{
		Id:     p.ID,
		Brand:  strings.TrimSpace(p.Brand.Name),
		Name:   p.Name,
		Ctg:    models.Category(p.Category.Key),
		SubCtg: strings.ToLower(p.Subcategory.Key),
	}
	for _, t := range p.LabResults.Terpenes {
		tempTerp := &models.Terpene{
			Name:        t.Terpene.Name,
			Description: t.Terpene.Description,
			Value:       t.Value,
		}
		product.T = append(product.T, tempTerp)
	}
	for _, c := range p.LabResults.Cannabinoids {
		tempCanna := &models.Cannabinoid{
			Name:        c.Cannabinoid.Name,
			Description: c.Cannabinoid.Description,
			Value:       c.Value,
		}
		product.C = append(product.C, tempCanna)
	}
	for _, i := range p.Images {
		product.Images = append(product.Images, i.URL)
	}
	return product
}

func (cT *ClientTranslator) TranslateClientCategory(category client.Category) *models.Category {
	c := models.Category(category.Key)
	return &c
}

func (cT *ClientTranslator) TranslateClientCategories(cs []client.Category) []*models.Category {
	categories := make([]*models.Category, 0)
	for _, c := range cs {
		categories = append(categories, cT.TranslateClientCategory(c))
	}
	return categories
}

func (cT *ClientTranslator) TranslateClientTerpene(terp client.TerpeneObj) *models.Terpene {
	return &models.Terpene{
		Name:        terp.Terpene.Name,
		Description: terp.Terpene.Description,
		Value:       terp.Value,
	}
}

func (cT *ClientTranslator) TranslateClientTerpenes(i []client.TerpeneObj) []*models.Terpene {
	ts := make([]*models.Terpene, 0)
	for _, t := range i {
		ts = append(ts, cT.TranslateClientTerpene(t))
	}
	return ts
}

func (cT *ClientTranslator) TranslateClientCannabinoid(c client.CannabinoidObj) *models.Cannabinoid {
	return &models.Cannabinoid{
		Name:        c.Cannabinoid.Name,
		Description: c.Cannabinoid.Description,
		Value:       c.Value,
	}
}

func (cT *ClientTranslator) TranslateClientCannabinoids(i []client.CannabinoidObj) []*models.Cannabinoid {
	cs := make([]*models.Cannabinoid, 0)
	for _, c := range i {
		cs = append(cs, cT.TranslateClientCannabinoid(c))
	}
	return cs
}

func (cT *ClientTranslator) TranslateClientOffers(os []client.Offer) []*models.Offer {
	offers := make([]*models.Offer, 0)
	for _, o := range os {
		offers = append(offers, cT.TranslateClientOffer(o))
	}
	return offers
}

func (cT *ClientTranslator) TranslateClientOffer(offer client.Offer) *models.Offer {
	return &models.Offer{
		Id:          offer.Id,
		Description: offer.Title,
	}
}
