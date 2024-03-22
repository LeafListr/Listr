package curaleaf

import (
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/Linkinlog/LeafListr/internal/models"
)

type ClientTranslator struct {
	numberRegex *regexp.Regexp
}

func NewClientTranslator() *ClientTranslator {
	return &ClientTranslator{
		numberRegex: regexp.MustCompile(`^-?(?:\d+|\d+\.\d+)`),
	}
}

func (cT *ClientTranslator) TranslateClientLocation(l Location) *models.Location {
	slug := ""
	if l.Location.StateSlug != "" && l.Location.State != "" {
		slug = "/" + l.Location.StateSlug + "/" + l.Slug
	}
	return &models.Location{
		Id:            l.UniqueId,
		Name:          l.Name,
		Address:       l.Location.Address,
		City:          l.Location.City,
		State:         l.Location.State,
		ZipCode:       l.Location.ZipCode,
		LocationTypes: l.MenuTypes,
		Slug:          slug,
	}
}

func (cT *ClientTranslator) TranslateClientLocations(ls []Location) []*models.Location {
	locations := make([]*models.Location, 0)
	for _, l := range ls {
		locations = append(locations, cT.TranslateClientLocation(l))
	}
	return locations
}

func (cT *ClientTranslator) TranslateClientProducts(ps []Product, menuUrl string) []*models.Product {
	products := make([]*models.Product, 0)

	for _, p := range ps {
		if len(p.Variants) == 0 {
			baseProduct := cT.TranslateClientProduct(p, menuUrl)
			products = append(products, baseProduct)
			slog.Debug("no variants for product", slog.String("product name", p.Name))
			continue
		}

		for _, v := range p.Variants {
			variantProduct := cT.TranslateClientProduct(p, menuUrl)
			variantProduct.Id = v.Id
			variantProduct.P = &models.Price{
				Total:           v.Price,
				DiscountedTotal: v.SpecialPrice,
				IsDiscounted:    v.IsSpecial,
			}
			variantProduct.Weight = v.Option

			if variantProduct.Ctg == ProductCategoryVape {
				split := strings.Split(p.CardDescription, "•")

				for j := range split {
					if strings.Contains(split[j], "THC") {
						thcSplit := strings.Split(split[j], "-")
						for k := range thcSplit {
							thc := cT.numberRegex.FindString(thcSplit[k])
							thcVal, err := strconv.ParseFloat(thc, 64)
							if err == nil {
								variantProduct.C = append(variantProduct.C, &models.Cannabinoid{
									Name:        "THC (Tetrahydrocannabinol)",
									Description: "Tetrahydrocannabinol",
									Value:       thcVal,
								})
							}
						}
					}
				}
			}
			variantProduct.P.PerGram = variantProduct.PricePerGram()
			products = append(products, variantProduct)
		}
	}

	return products
}

func (cT *ClientTranslator) TranslateClientProduct(p Product, menuUrl string) *models.Product {
	slug := "/products/" + p.ID
	permalink := ""
	if menuUrl != "" {
		permalink = menuUrl + slug
	}
	product := &models.Product{
		Id:        p.ID,
		Brand:     strings.TrimSpace(p.Brand.Name),
		Name:      p.Name,
		Ctg:       p.Category.Key,
		SubCtg:    strings.ToLower(p.Subcategory.Key),
		Slug:      slug,
		Permalink: permalink,
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

func (cT *ClientTranslator) TranslateClientCategory(category Category) string {
	category.Key = strings.ToLower(category.Key)
	return category.Key
}

func (cT *ClientTranslator) TranslateClientCategories(cs []Category) []string {
	categories := make([]string, 0)
	for _, c := range cs {
		categories = append(categories, cT.TranslateClientCategory(c))
	}
	return categories
}

func (cT *ClientTranslator) TranslateClientSubCategories(ps []Product) []string {
	subCatMap := make(map[string]struct{})
	for _, p := range ps {
		subCatMap[p.Subcategory.DisplayName] = struct{}{}
	}
	subCategories := make([]string, 0)
	for k := range subCatMap {
		subCategories = append(subCategories, k)
	}
	return subCategories
}

func (cT *ClientTranslator) TranslateClientTerpene(terp TerpeneObj) *models.Terpene {
	return &models.Terpene{
		Name:        terp.Terpene.Name,
		Description: terp.Terpene.Description,
		Value:       terp.Value,
	}
}

func (cT *ClientTranslator) TranslateClientTerpenes(i []TerpeneObj) []*models.Terpene {
	ts := make([]*models.Terpene, 0)
	for _, t := range i {
		ts = append(ts, cT.TranslateClientTerpene(t))
	}
	return ts
}

func (cT *ClientTranslator) TranslateClientCannabinoid(c CannabinoidObj) *models.Cannabinoid {
	return &models.Cannabinoid{
		Name:        c.Cannabinoid.Name,
		Description: c.Cannabinoid.Description,
		Value:       c.Value,
	}
}

func (cT *ClientTranslator) TranslateClientCannabinoids(i []CannabinoidObj) []*models.Cannabinoid {
	cs := make([]*models.Cannabinoid, 0)
	for _, c := range i {
		cs = append(cs, cT.TranslateClientCannabinoid(c))
	}
	return cs
}

func (cT *ClientTranslator) TranslateClientOffers(os []Offer) []*models.Offer {
	offers := make([]*models.Offer, 0)
	for _, o := range os {
		offers = append(offers, cT.TranslateClientOffer(o))
	}
	return offers
}

func (cT *ClientTranslator) TranslateClientOffer(offer Offer) *models.Offer {
	return &models.Offer{
		Id:          offer.Id,
		Description: offer.Title,
	}
}
