package repository

import (
	"github.com/Linkinlog/LeafList/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Repository
type Repository interface {
	GetLocations(longitude, latitude float64) ([]*models.Brand, error)
	GetProduct(menuId, productId string) (*models.Product, error)
	GetProducts(menuId string) ([]*models.Product, error)
	GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error)
	GetCategories(menuId string) ([]models.Category, error)
	GetTerpenes(menuId string) ([]*models.Terpene, error)
	GetCannabinoids(menuId string) ([]*models.Cannabinoid, error)
	GetOffers(menuId string) ([]*models.Offer, error)
}
