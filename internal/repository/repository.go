package repository

import (
	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Repository
type Repository interface {
	MenuRepository
	ProductRepository
	CategoryRepository
	CannabinoidRepository
	TerpeneRepository
	OfferRepository
}

type MenuRepository interface {
	GetMenu(menuId string) (*models.Dispensary, error)
	GetMenus(longitude, latitude float64) ([]*models.Dispensary, error)
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
