package repository

import "github.com/Linkinlog/LeafList/internal/repository/models"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Repository
type Repository interface {
	GetProducts(menuId string) ([]*models.Product, error)
	GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error)
	GetCategories(menuId string) ([]models.Category, error)
	GetTerpenes(menuId string) ([]*models.Terpene, error)
	GetOffers(menuId string) ([]*models.Offer, error)
}
