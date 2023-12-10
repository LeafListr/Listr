package translation

import (
	apiModels "github.com/Linkinlog/LeafListr/internal/api/models"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
)

type APITranslator struct{}

func (aT *APITranslator) TranslateAPILocation(l *models.Location) *apiModels.Location {
	return &apiModels.Location{
		Id:      l.Id,
		Name:    l.Name,
		Address: l.Address,
		City:    l.City,
		State:   l.State,
		ZipCode: l.ZipCode,
	}
}

func (aT *APITranslator) TranslateAPILocations(ls []*models.Location) []*apiModels.Location {
	apiLs := make([]*apiModels.Location, 0)
	for _, l := range ls {
		apiLs = append(apiLs, aT.TranslateAPILocation(l))
	}
	return apiLs
}

func (aT *APITranslator) TranslateAPIDispensary(d *models.Dispensary) *apiModels.Dispensary {
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

func (aT *APITranslator) TranslateAPIDispensaries(ds []*models.Dispensary) []*apiModels.Dispensary {
	apiDs := make([]*apiModels.Dispensary, 0)
	for _, d := range ds {
		apiDs = append(apiDs, aT.TranslateAPIDispensary(d))
	}
	return apiDs
}

func (aT *APITranslator) TranslateAPIProduct(p *models.Product) *apiModels.Product {
	apiP := &apiModels.Product{
		Id:      p.Id,
		Brand:   p.Brand,
		Name:    p.Name,
		Images:  p.Images,
		Ctg:     apiModels.Category(p.Ctg),
		SubCtg:  p.SubCtg,
		Variant: p.Variant,
		Price: &apiModels.Price{
			Total:           p.Price.Total,
			DiscountedTotal: p.Price.DiscountedTotal,
			IsDiscounted:    p.Price.IsDiscounted,
		},
	}

	for _, c := range p.C {
		apiP.C = append(apiP.C, aT.TranslateAPICannabinoid(c))
	}

	for _, terp := range p.T {
		apiP.T = append(apiP.T, aT.TranslateAPITerpene(terp))
	}

	return apiP
}

func (aT *APITranslator) TranslateAPIProducts(ps []*models.Product) []*apiModels.Product {
	apiPs := make([]*apiModels.Product, 0)
	for _, p := range ps {
		if p != nil {
			apiPs = append(apiPs, aT.TranslateAPIProduct(p))
		}
	}
	return apiPs
}

func (aT *APITranslator) TranslateAPICategory(c *models.Category) *apiModels.Category {
	return (*apiModels.Category)(c)
}

func (aT *APITranslator) TranslateAPICategories(cs []*models.Category) []*apiModels.Category {
	apiCs := make([]*apiModels.Category, 0)
	for _, c := range cs {
		apiCs = append(apiCs, aT.TranslateAPICategory(c))
	}
	return apiCs
}

func (aT *APITranslator) TranslateAPITerpene(terp *models.Terpene) *apiModels.Terpene {
	return &apiModels.Terpene{
		Name:        terp.Name,
		Description: terp.Description,
		Value:       terp.Value,
	}
}

func (aT *APITranslator) TranslateAPITerpenes(terps []*models.Terpene) []*apiModels.Terpene {
	apiTerps := make([]*apiModels.Terpene, 0)
	for _, terp := range terps {
		apiTerps = append(apiTerps, aT.TranslateAPITerpene(terp))
	}
	return apiTerps
}

func (aT *APITranslator) TranslateAPICannabinoid(c *models.Cannabinoid) *apiModels.Cannabinoid {
	return &apiModels.Cannabinoid{
		Name:        c.Name,
		Description: c.Description,
		Value:       c.Value,
	}
}

func (aT *APITranslator) TranslateAPICannabinoids(cs []*models.Cannabinoid) []*apiModels.Cannabinoid {
	apiCs := make([]*apiModels.Cannabinoid, 0)
	for _, c := range cs {
		apiCs = append(apiCs, aT.TranslateAPICannabinoid(c))
	}
	return apiCs
}

func (aT *APITranslator) TranslateAPIOffer(o *models.Offer) *apiModels.Offer {
	return &apiModels.Offer{
		Id:          o.Id,
		Description: o.Description,
	}
}

func (aT *APITranslator) TranslateAPIOffers(os []*models.Offer) []*apiModels.Offer {
	apiOs := make([]*apiModels.Offer, 0)
	for _, o := range os {
		apiOs = append(apiOs, aT.TranslateAPIOffer(o))
	}
	return apiOs
}

func NewAPITranslator() translation.APITranslatable {
	return &APITranslator{}
}
