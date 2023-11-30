package workflow

import (
	"context"

	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Manager
type Manager interface {
	LocationManager
	ProductManager
	CategoryManager
	TerpeneManager
	CannabinoidManager
	OfferManager
	ErrorManager
}

type LocationManager interface {
	Location(dispensary, menuId string) (*models.Location, error)
	Locations(dispensary string) ([]*models.Location, error)
}

type ProductManager interface {
	Product(dispensary, menuId, productId string) (*models.Product, error)
	Products(dispensary, menuId string) ([]*models.Product, error)
	ProductsForCategory(dispensary, menuId string, category models.Category) ([]*models.Product, error)
	ProductsForSubCategory(dispensary, menuId string, products []*models.Product, subCategory string) ([]*models.Product, error)
	ProductsForBrands(dispensary, menuId string, products []*models.Product, brands []string) ([]*models.Product, error)
	ProductsExcludingBrands(dispensary, menuId string, products []*models.Product, brands []string) ([]*models.Product, error)
	ProductsForPriceRange(dispensary, menuId string, products []*models.Product, min, max float64) ([]*models.Product, error)
}

type CategoryManager interface {
	Categories(dispensary, menuId string) ([]*models.Category, error)
}

type TerpeneManager interface {
	Terpenes(dispensary, menuId string) ([]*models.Terpene, error)
}

type CannabinoidManager interface {
	Cannabinoids(dispensary, menuId string) ([]*models.Cannabinoid, error)
}

type OfferManager interface {
	Offers(dispensary, menuId string) ([]*models.Offer, error)
}

type ErrorManager interface {
	LogError(err error, context context.Context)
}
