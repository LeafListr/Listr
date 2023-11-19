package translation

import (
	apiModels "github.com/Linkinlog/LeafListr/internal/api/models"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . ClientTranslatable
type ClientTranslatable interface {
	TranslateClientLocation(*client.Location) *models.Location
	TranslateClientLocations([]*client.Location) []*models.Location
	TranslateClientProduct(*client.Product) *models.Product
	TranslateClientProducts([]*client.Product) []*models.Product
	TranslateClientCategory(*client.Category) []*models.Category
	TranslateClientCategories([]*client.Category) []*models.Category
	TranslateClientTerpene(*client.TerpeneObj) *models.Terpene
	TranslateClientTerpenes(*[]client.TerpeneObj) *models.Terpene
	TranslateClientCannabinoid(*client.CannabinoidObj) *models.Cannabinoid
	TranslateClientCannabinoids(*[]client.CannabinoidObj) *models.Cannabinoid
	TranslateClientOffer(*client.Offer) []*models.Offer
	TranslateClientOffers([]*client.Offer) []*models.Offer
}

//counterfeiter:generate . APITranslatable
type APITranslatable interface {
	TranslateAPILocation(*models.Location) *apiModels.Location
	TranslateAPILocations([]*models.Location) []*apiModels.Location
	TranslateAPIDispensary(*models.Dispensary) *apiModels.Dispensary
	TranslateAPIDispensaries([]*models.Dispensary) []*apiModels.Dispensary
	TranslateAPIProduct(*models.Product) *apiModels.Product
	TranslateAPIProducts([]*models.Product) []*apiModels.Product
	TranslateAPICategory(*models.Category) *apiModels.Category
	TranslateAPICategories([]*models.Category) []*apiModels.Category
	TranslateAPITerpene(*models.Terpene) *apiModels.Terpene
	TranslateAPITerpenes([]*models.Terpene) []*apiModels.Terpene
	TranslateAPICannabinoid(*models.Cannabinoid) *apiModels.Cannabinoid
	TranslateAPICannabinoids([]*models.Cannabinoid) []*apiModels.Cannabinoid
	TranslateAPIOffer(*models.Offer) *apiModels.Offer
	TranslateAPIOffers([]*models.Offer) []*apiModels.Offer
}
