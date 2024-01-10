package workflow

import (
	"context"

	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Workflow
type Workflow interface {
	LocationManager
	ProductManager
	CategoryManager
	TerpeneManager
	CannabinoidManager
	OfferManager
	ErrorManager
	ProductFilter
	ProductSorter
}

type LocationManager interface {
	Location(dispensary, menuId, menuType string) (*models.Location, error)
	Locations(dispensary, menuType string) ([]*models.Location, error)
}

type ProductManager interface {
	Product(dispensary, menuId, menuType, productId string) (*models.Product, error)
	Products(dispensary, menuId, menuType string) ([]*models.Product, error)
}

type ProductFilter interface {
	ProductsForCategory(dispensary, menuId, menuType string, category models.Category) ([]*models.Product, error)
	ProductsForSubCategory(dispensary, menuId, menuType string, products []*models.Product, subCategory string) ([]*models.Product, error)
	ProductsForBrands(dispensary, menuId, menuType string, products []*models.Product, brands []string) ([]*models.Product, error)
	ProductsExcludingBrands(dispensary, menuId, menuType string, products []*models.Product, brands []string) ([]*models.Product, error)
	ProductsForVariants(dispensary, menuId, menuType string, products []*models.Product, variants []string) ([]*models.Product, error)
	ProductsForPriceRange(dispensary, menuId, menuType string, products []*models.Product, min, max float64) ([]*models.Product, error)
}

type ProductSorter interface {
	SortProductsByPriceAsc(dispensary, menuId, menuType string, products []*models.Product)
	SortProductsByPriceDesc(dispensary, menuId, menuType string, products []*models.Product)
	SortProductsByTHCAsc(dispensary, menuId, menuType string, products []*models.Product)
	SortProductsByTHCDesc(dispensary, menuId, menuType string, products []*models.Product)
	SortProductsByTop3Terps(dispensary, menuId, menuType string, products []*models.Product, terps [3]string)
}

type CategoryManager interface {
	Categories(dispensary, menuId, menuType string) ([]models.Category, error)
}

type TerpeneManager interface {
	Terpenes(dispensary, menuId, menuType string) ([]*models.Terpene, error)
}

type CannabinoidManager interface {
	Cannabinoids(dispensary, menuId, menuType string) ([]*models.Cannabinoid, error)
}

type OfferManager interface {
	Offers(dispensary, menuId, menuType string) ([]*models.Offer, error)
}

type ErrorManager interface {
	LogError(err error, context context.Context)
}
