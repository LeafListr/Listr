package translation

import (
	"github.com/Linkinlog/LeafListr/internal/client/curaleaf"
	"github.com/Linkinlog/LeafListr/internal/models"
)

type Translatable interface {
	TranslateLocations(ls []*curaleaf.Location) []*models.Dispensary
	TranslateProduct(*curaleaf.Product) *models.Product
	TranslateProducts([]*curaleaf.Product) []*models.Product
	TranslateCategories([]*curaleaf.Category) []*models.Category
	TranslateTerpene(*curaleaf.TerpeneObj) *models.Terpene
	TranslateCannabinoid(cannabinoid *curaleaf.CannabinoidObj) *models.Cannabinoid
	TranslateOffers([]*curaleaf.Offer) []*models.Offer
}

type Translator struct{}

func (t *Translator) TranslateLocations(ls []*curaleaf.Location) []*models.Dispensary {
	var dispensaries []*models.Dispensary
	for _, l := range ls {
		if l == nil {
			return dispensaries
		}
		dispensaries = append(dispensaries, &models.Dispensary{
			UniqueId:   l.UniqueId,
			Name:       l.Name,
			OrderTypes: l.OrderTypes,
			MenuTypes:  l.MenuTypes,
			IsOpened:   l.IsOpened,
			Location: models.Location{
				Coordinates: models.Coordinates{
					Latitude:  l.Location.Coordinates.Latitude,
					Longitude: l.Location.Coordinates.Longitude,
				},
				Address: l.Location.Address,
				City:    l.Location.City,
				State:   l.Location.City,
				ZipCode: l.Location.ZipCode,
			},
		})
	}
	return dispensaries
}

func (t *Translator) TranslateProducts(ps []*curaleaf.Product) []*models.Product {
	var products []*models.Product
	for _, p := range ps {
		if p == nil {
			return products
		}
		products = append(products, t.TranslateProduct(p))
	}
	return products
}

func (t *Translator) TranslateProduct(p *curaleaf.Product) *models.Product {
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
			Variant:         v.Option,
			Total:           v.Price,
			DiscountedTotal: v.SpecialPrice,
		}
		product.P = append(product.P, tempPrice)
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

func (t *Translator) TranslateCategories(cs []*curaleaf.Category) []*models.Category {
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

func (t *Translator) TranslateTerpene(terp *curaleaf.TerpeneObj) *models.Terpene {
	if terp == nil {
		return &models.Terpene{}
	}

	return &models.Terpene{
		Name:        terp.Terpene.Name,
		Description: terp.Terpene.Description,
		Value:       terp.Value,
	}
}

func (t *Translator) TranslateCannabinoid(c *curaleaf.CannabinoidObj) *models.Cannabinoid {
	if c == nil {
		return &models.Cannabinoid{}
	}
	return &models.Cannabinoid{
		Name:        c.Cannabinoid.Name,
		Description: c.Cannabinoid.Description,
		Value:       c.Value,
	}
}

func (t *Translator) TranslateOffers(os []*curaleaf.Offer) []*models.Offer {
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

func NewTranslator() Translatable {
	return &Translator{}
}
