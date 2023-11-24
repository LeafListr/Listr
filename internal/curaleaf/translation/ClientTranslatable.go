package translation

import (
	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
)

type ClientTranslator struct{}

func (cT *ClientTranslator) TranslateClientLocation(l client.Location) *models.Location {
	return &models.Location{
		Name:    l.UniqueId,
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
		products = append(products, cT.TranslateClientProduct(p))
	}
	return products
}

func (cT *ClientTranslator) TranslateClientProduct(p client.Product) *models.Product {
	product := &models.Product{
		Id:   p.ID,
		Name: p.Name,
		Ctg:  models.Category(p.Category.Key),
	}
	for _, v := range p.Variants {
		tempPrice := &models.Price{
			Total:           v.Price,
			DiscountedTotal: v.SpecialPrice,
		}
		tempVariant := &models.Variant{
			Name:  v.Option,
			Price: tempPrice,
		}
		product.V = append(product.V, tempVariant)
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

func NewClientTranslator() translation.ClientTranslatable {
	return &ClientTranslator{}
}
