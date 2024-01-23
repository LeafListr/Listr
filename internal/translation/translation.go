package translation

import (
	apiModels "github.com/Linkinlog/LeafListr/internal/api/models"
	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . APITranslatable
type APITranslatable interface {
	TranslateLocation(*models.Location) *apiModels.Location
	TranslateLocations([]*models.Location) []*apiModels.Location
	TranslateDispensary(*models.Dispensary) *apiModels.Dispensary
	TranslateDispensaries([]*models.Dispensary) []*apiModels.Dispensary
	TranslateProduct(*models.Product) *apiModels.Product
	TranslateProducts([]*models.Product) []*apiModels.Product
	TranslateCategory(string) string
	TranslateCategories([]string) []string
	TranslateTerpene(*models.Terpene) *apiModels.Terpene
	TranslateTerpenes([]*models.Terpene) []*apiModels.Terpene
	TranslateCannabinoid(*models.Cannabinoid) *apiModels.Cannabinoid
	TranslateCannabinoids([]*models.Cannabinoid) []*apiModels.Cannabinoid
	TranslateOffer(*models.Offer) *apiModels.Offer
	TranslateOffers([]*models.Offer) []*apiModels.Offer
}
