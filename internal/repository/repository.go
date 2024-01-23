package repository

import (
	"errors"

	"github.com/Linkinlog/LeafListr/internal/models"
)

var (
	InvalidJSONError     = errors.New("invalid json")
	ResourceNotFound     = errors.New("resource not found")
	InvalidCategoryError = errors.New("invalid category")
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Repository
type Repository interface {
	LocationRepository
	ProductRepository
	CategoryRepository
	CannabinoidRepository
	TerpeneRepository
	OfferRepository
}

type LocationRepository interface {
	Location() (*models.Location, error)
	Locations(longitude, latitude float64) ([]*models.Location, error)
}

type ProductRepository interface {
	GetProduct(productId string) (*models.Product, error)
	GetProducts() ([]*models.Product, error)
	GetProductsForCategory(category string) ([]*models.Product, error)
}

type CategoryRepository interface {
	GetCategories() ([]string, error)
}

type TerpeneRepository interface {
	GetTerpenes() ([]*models.Terpene, error)
}

type CannabinoidRepository interface {
	GetCannabinoids() ([]*models.Cannabinoid, error)
}

type OfferRepository interface {
	GetOffers() ([]*models.Offer, error)
}
