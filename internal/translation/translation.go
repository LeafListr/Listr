package translation

import (
	apiModels "github.com/Linkinlog/LeafListr/internal/api/models"
	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

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
