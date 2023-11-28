package repository

import (
	"errors"

	"github.com/Linkinlog/LeafListr/internal/models"
)

var (
	InvalidJSONError = errors.New("invalid json")
	ResourceNotFound = errors.New("resource not found")
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
	Location(menuId string) (*models.Location, error)
	Locations(longitude, latitude float64) ([]*models.Location, error)
}

type ProductRepository interface {
	GetProduct(menuId, productId string) (*models.Product, error)
	GetProducts(menuId string) ([]*models.Product, error)
	GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error)
}

type CategoryRepository interface {
	GetCategories(menuId string) ([]*models.Category, error)
}

type TerpeneRepository interface {
	GetTerpenes(menuId string) ([]*models.Terpene, error)
}

type CannabinoidRepository interface {
	GetCannabinoids(menuId string) ([]*models.Cannabinoid, error)
}

type OfferRepository interface {
	GetOffers(menuId string) ([]*models.Offer, error)
}
