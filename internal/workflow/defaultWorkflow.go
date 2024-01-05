package workflow

import (
	"context"
	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
	"log/slog"
)

type defaultWorkflow struct {
	//rf factory.RepositoryFactory
	c cache.Cacher
	f transformation.Filterer
}

func NewWorkflow() Workflow {
	return &defaultWorkflow{}
}

func (w *defaultWorkflow) Location(dispensary, menuId, menuType string) (*models.Location, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Locations(dispensary, menuType string) ([]*models.Location, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Product(dispensary, menuId, menuType, productId string) (*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Products(dispensary, menuId, menuType string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Categories(dispensary, menuId, menuType string) ([]*models.Category, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Terpenes(dispensary, menuId, menuType string) ([]*models.Terpene, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Cannabinoids(dispensary, menuId, menuType string) ([]*models.Cannabinoid, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) Offers(dispensary, menuId, menuType string) ([]*models.Offer, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) LogError(err error, context context.Context) {
	slog.InfoContext(context, err.Error())
}

func (w *defaultWorkflow) ProductsForCategory(dispensary, menuId, menuType string, category models.Category) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) ProductsForSubCategory(dispensary, menuId, menuType string, products []*models.Product, subCategory string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) ProductsForBrands(dispensary, menuId, menuType string, products []*models.Product, brands []string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) ProductsExcludingBrands(dispensary, menuId, menuType string, products []*models.Product, brands []string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) ProductsForVariants(dispensary, menuId, menuType string, products []*models.Product, variants []string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) ProductsForPriceRange(dispensary, menuId, menuType string, products []*models.Product, min, max float64) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) SortProductsByPriceAsc(dispensary, menuId, menuType string, products []*models.Product) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) SortProductsByPriceDesc(dispensary, menuId, menuType string, products []*models.Product) {
	// TODO implement me
	panic("implement me")
}

func (w *defaultWorkflow) SortProductsByTop3Terps(dispensary, menuId, menuType string, products []*models.Product, terps [3]string) {
	// TODO implement me
	panic("implement me")
}
