package translation

import (
	"github.com/Linkinlog/LeafList/internal/client/curaleaf"
	"github.com/Linkinlog/LeafList/internal/models"
)

type Translatable interface {
	TranslateLocations(ls []curaleaf.Location) []*models.Brand
	TranslateProduct(curaleaf.Product) *models.Product
	TranslateProducts([]curaleaf.Product) []*models.Product
	TranslateCategories([]curaleaf.Category) []models.Category
	TranslateTerpenes([]curaleaf.Product) []*models.Terpene
	TranslateCannabinoids([]curaleaf.Product) []*models.Cannabinoid
	TranslateOffers([]curaleaf.Offer) []*models.Offer
}

type Translator struct{}

func (t *Translator) TranslateLocations(ls []curaleaf.Location) []*models.Brand {
	var mLs []*models.Brand
	for _, l := range ls {
		mLs = append(mLs, &models.Brand{
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
	return mLs
}

func (t *Translator) TranslateProducts(ps []curaleaf.Product) []*models.Product {
	var mPs []*models.Product
	for _, p := range ps {
		mPs = append(mPs, t.TranslateProduct(p))
	}
	return mPs
}

func (t *Translator) TranslateProduct(p curaleaf.Product) *models.Product {
	mp := &models.Product{
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
		mp.P = append(mp.P, tempPrice)
	}
	for _, t := range p.LabResults.Terpenes {
		tempTerp := &models.Terpene{
			Name:        t.Terpene.Name,
			Description: t.Terpene.Description,
			Value:       t.Value,
		}
		mp.T = append(mp.T, tempTerp)
	}
	for _, c := range p.LabResults.Cannabinoids {
		tempCanna := &models.Cannabinoid{
			Name:        c.Cannabinoid.Name,
			Description: c.Cannabinoid.Description,
			Value:       c.Value,
		}
		mp.C = append(mp.C, tempCanna)
	}
	return mp
}

func (t *Translator) TranslateCategories(cs []curaleaf.Category) []models.Category {
	var mCs []models.Category
	for _, c := range cs {
		tempC := models.Category(c.Key)
		mCs = append(mCs, tempC)
	}
	return mCs
}

func (t *Translator) TranslateTerpenes(products []curaleaf.Product) []*models.Terpene {
	var mTerpenes []*models.Terpene
	for _, product := range products {
		for _, terp := range product.LabResults.Terpenes {
			tempTerp := &models.Terpene{
				Name:        terp.Terpene.Name,
				Description: terp.Terpene.Description,
				Value:       terp.Value,
			}
			mTerpenes = append(mTerpenes, tempTerp)
		}
	}
	return mTerpenes
}

func (t *Translator) TranslateCannabinoids(products []curaleaf.Product) []*models.Cannabinoid {
	var mCannabinoids []*models.Cannabinoid
	for _, product := range products {
		for _, cannabinoid := range product.LabResults.Cannabinoids {
			tempCannabinoid := &models.Cannabinoid{
				Name:        cannabinoid.Cannabinoid.Name,
				Description: cannabinoid.Cannabinoid.Description,
				Value:       cannabinoid.Value,
			}
			mCannabinoids = append(mCannabinoids, tempCannabinoid)
		}
	}
	return mCannabinoids
}

func (t *Translator) TranslateOffers(offers []curaleaf.Offer) []*models.Offer {
	var mOffers []*models.Offer
	for _, o := range offers {
		tempOffer := &models.Offer{
			Id:          o.Id,
			Description: o.Title,
		}
		mOffers = append(mOffers, tempOffer)
	}
	return mOffers
}

func NewTranslator() Translatable {
	return &Translator{}
}
