package translation

import (
	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
)

type ClientTranslator struct{}

func (cT *ClientTranslator) TranslateClientLocation(l *client.Location) *models.Location {
	return &models.Location{
		Name:    l.UniqueId,
		Address: l.Location.Address,
		City:    l.Location.City,
		State:   l.Location.City,
		ZipCode: l.Location.ZipCode,
	}
}

func (cT *ClientTranslator) TranslateClientLocations(ls []*client.Location) []*models.Location {
	var locations []*models.Location
	for _, l := range ls {
		if l == nil {
			return locations
		}
		locations = append(locations, cT.TranslateClientLocation(l))
	}
	return locations
}

func (cT *ClientTranslator) TranslateClientProducts(ps []*client.Product) []*models.Product {
	var products []*models.Product
	for _, p := range ps {
		products = append(products, cT.TranslateClientProduct(p))
	}
	return products
}

func (cT *ClientTranslator) TranslateClientProduct(p *client.Product) *models.Product {
	if p == nil {
		return &models.Product{}
	}
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

func (cT *ClientTranslator) TranslateClientCategory(category *client.Category) []*models.Category {
	// TODO implement me
	panic("implement me")
}

func (cT *ClientTranslator) TranslateClientCategories(cs []*client.Category) []*models.Category {
	var categories []*models.Category
	for _, c := range cs {
		if c == nil {
			return categories
		}
		tempC := models.Category(c.Key)
		categories = append(categories, &tempC)
	}
	return categories
}

func (cT *ClientTranslator) TranslateClientTerpene(terp *client.TerpeneObj) *models.Terpene {
	if terp == nil {
		return &models.Terpene{}
	}

	return &models.Terpene{
		Name:        terp.Terpene.Name,
		Description: terp.Terpene.Description,
		Value:       terp.Value,
	}
}

func (cT *ClientTranslator) TranslateClientTerpenes(i *[]client.TerpeneObj) *models.Terpene {
	// TODO implement me
	panic("implement me")
}

func (cT *ClientTranslator) TranslateClientCannabinoid(c *client.CannabinoidObj) *models.Cannabinoid {
	if c == nil {
		return &models.Cannabinoid{}
	}
	return &models.Cannabinoid{
		Name:        c.Cannabinoid.Name,
		Description: c.Cannabinoid.Description,
		Value:       c.Value,
	}
}

func (cT *ClientTranslator) TranslateClientCannabinoids(i *[]client.CannabinoidObj) *models.Cannabinoid {
	// TODO implement me
	panic("implement me")
}

func (cT *ClientTranslator) TranslateClientOffers(os []*client.Offer) []*models.Offer {
	var offers []*models.Offer
	for _, o := range os {
		if o == nil {
			return offers
		}
		tempOffer := &models.Offer{
			Id:          o.Id,
			Description: o.Title,
		}
		offers = append(offers, tempOffer)
	}
	return offers
}

func (cT *ClientTranslator) TranslateClientOffer(offer *client.Offer) []*models.Offer {
	// TODO implement me
	panic("implement me")
}

func NewClientTranslator() translation.ClientTranslatable {
	return &ClientTranslator{}
}
