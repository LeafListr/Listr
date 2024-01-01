package beyondhello

import (
	"context"
	"log/slog"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/workflow"
)

type Workflow struct{}

func NewWorkflow() workflow.Workflow {
	return &Workflow{}
}

func (w *Workflow) Location(dispensary, menuId, menuType string) (*models.Location, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) Locations(dispensary, menuType string) ([]*models.Location, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) Product(dispensary, menuId, menuType, productId string) (*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) Products(dispensary, menuId, menuType string) ([]*models.Product, error) {
	repo := NewRepository()
	return repo.GetProducts(menuId)
}

func (w *Workflow) Categories(dispensary, menuId, menuType string) ([]*models.Category, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) Terpenes(dispensary, menuId, menuType string) ([]*models.Terpene, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) Cannabinoids(dispensary, menuId, menuType string) ([]*models.Cannabinoid, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) Offers(dispensary, menuId, menuType string) ([]*models.Offer, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) LogError(err error, context context.Context) {
	slog.InfoContext(context, err.Error())
}

func (w *Workflow) ProductsForCategory(dispensary, menuId, menuType string, category models.Category) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) ProductsForSubCategory(dispensary, menuId, menuType string, products []*models.Product, subCategory string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) ProductsForBrands(dispensary, menuId, menuType string, products []*models.Product, brands []string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) ProductsExcludingBrands(dispensary, menuId, menuType string, products []*models.Product, brands []string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) ProductsForVariants(dispensary, menuId, menuType string, products []*models.Product, variants []string) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) ProductsForPriceRange(dispensary, menuId, menuType string, products []*models.Product, min, max float64) ([]*models.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) SortProductsByPriceAsc(dispensary, menuId, menuType string, products []*models.Product) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) SortProductsByPriceDesc(dispensary, menuId, menuType string, products []*models.Product) {
	// TODO implement me
	panic("implement me")
}

func (w *Workflow) SortProductsByTop3Terps(dispensary, menuId, menuType string, products []*models.Product, terps [3]string) {
	// TODO implement me
	panic("implement me")
}
