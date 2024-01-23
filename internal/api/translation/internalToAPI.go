package translation

import (
	apiModels "github.com/Linkinlog/LeafListr/internal/api/models"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
)

type internalToAPI struct{}

func (i *internalToAPI) TranslateLocation(l *models.Location) *apiModels.Location {
	return &apiModels.Location{
		Id:      l.Id,
		Name:    l.Name,
		Address: l.Address,
		City:    l.City,
		State:   l.State,
		ZipCode: l.ZipCode,
	}
}

func (i *internalToAPI) TranslateLocations(ls []*models.Location) []*apiModels.Location {
	apiLs := make([]*apiModels.Location, 0)
	for _, l := range ls {
		apiLs = append(apiLs, i.TranslateLocation(l))
	}
	return apiLs
}

func (i *internalToAPI) TranslateDispensary(d *models.Dispensary) *apiModels.Dispensary {
	apiD := &apiModels.Dispensary{Name: d.Name}
	for _, loc := range d.Locations {
		temp := apiModels.Location{
			Name:    loc.Name,
			Address: loc.Address,
			City:    loc.City,
			State:   loc.State,
			ZipCode: loc.ZipCode,
		}
		apiD.Locations = append(apiD.Locations, &temp)
	}
	return apiD
}

func (i *internalToAPI) TranslateDispensaries(ds []*models.Dispensary) []*apiModels.Dispensary {
	apiDs := make([]*apiModels.Dispensary, 0)
	for _, d := range ds {
		apiDs = append(apiDs, i.TranslateDispensary(d))
	}
	return apiDs
}

func (i *internalToAPI) TranslateProduct(p *models.Product) *apiModels.Product {
	apiP := &apiModels.Product{
		Id:      p.Id,
		Brand:   p.Brand,
		Name:    p.Name,
		Images:  p.Images,
		Ctg:     string(p.Ctg),
		SubCtg:  p.SubCtg,
		Variant: p.Weight,
		Price: &apiModels.Price{
			Total:           p.Price.Total,
			DiscountedTotal: p.Price.DiscountedTotal,
			IsDiscounted:    p.Price.IsDiscounted,
		},
	}

	for _, c := range p.C {
		apiP.C = append(apiP.C, i.TranslateCannabinoid(c))
	}

	for _, terp := range p.T {
		apiP.T = append(apiP.T, i.TranslateTerpene(terp))
	}

	return apiP
}

func (i *internalToAPI) TranslateProducts(ps []*models.Product) []*apiModels.Product {
	apiPs := make([]*apiModels.Product, 0)
	for _, p := range ps {
		if p != nil {
			apiPs = append(apiPs, i.TranslateProduct(p))
		}
	}
	return apiPs
}

func (i *internalToAPI) TranslateCategory(c string) string {
	return string(c)
}

func (i *internalToAPI) TranslateCategories(cs []string) []string {
	apiCs := make([]string, 0)
	for _, c := range cs {
		apiCs = append(apiCs, i.TranslateCategory(c))
	}
	return apiCs
}

func (i *internalToAPI) TranslateTerpene(terp *models.Terpene) *apiModels.Terpene {
	return &apiModels.Terpene{
		Name:        terp.Name,
		Description: terp.Description,
		Value:       terp.Value,
	}
}

func (i *internalToAPI) TranslateTerpenes(terps []*models.Terpene) []*apiModels.Terpene {
	apiTerps := make([]*apiModels.Terpene, 0)
	for _, terp := range terps {
		apiTerps = append(apiTerps, i.TranslateTerpene(terp))
	}
	return apiTerps
}

func (i *internalToAPI) TranslateCannabinoid(c *models.Cannabinoid) *apiModels.Cannabinoid {
	return &apiModels.Cannabinoid{
		Name:        c.Name,
		Description: c.Description,
		Value:       c.Value,
	}
}

func (i *internalToAPI) TranslateCannabinoids(cs []*models.Cannabinoid) []*apiModels.Cannabinoid {
	apiCs := make([]*apiModels.Cannabinoid, 0)
	for _, c := range cs {
		apiCs = append(apiCs, i.TranslateCannabinoid(c))
	}
	return apiCs
}

func (i *internalToAPI) TranslateOffer(o *models.Offer) *apiModels.Offer {
	return &apiModels.Offer{
		Id:          o.Id,
		Description: o.Description,
	}
}

func (i *internalToAPI) TranslateOffers(os []*models.Offer) []*apiModels.Offer {
	apiOs := make([]*apiModels.Offer, 0)
	for _, o := range os {
		apiOs = append(apiOs, i.TranslateOffer(o))
	}
	return apiOs
}

func NewAPITranslator() translation.APITranslatable {
	return &internalToAPI{}
}
