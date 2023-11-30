package manager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/cache"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"

	curaFactory "github.com/Linkinlog/LeafListr/internal/curaleaf/factory"
	curaTransformer "github.com/Linkinlog/LeafListr/internal/curaleaf/transformation"
	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/workflow"
)

type Manager struct {
	F  factory.RepositoryFactory
	TF transformation.Filterer
}

func NewWorkflowManager() workflow.Manager {
	return &Manager{
		F:  curaFactory.NewRepoFactory(cache.NewCache()),
		TF: curaTransformer.NewFilterer(),
	}
}

func (w *Manager) Location(dispensary, menuId string) (*models.Location, error) {
	repo, err := w.F.FindByDispensary(dispensary)
	if err != nil {
		return &models.Location{}, fmt.Errorf("couldn't find dispensary by menu for location. Dispensary=%s, MenuId=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.Location(menuId)
}

func (w *Manager) Locations(dispensary string) ([]*models.Location, error) {
	repo, err := w.F.FindByDispensary(dispensary)
	if err != nil {
		return []*models.Location{}, fmt.Errorf("couldn't find dispensary by menu for locations. Dispensary=%s. Err: %v", dispensary, err)
	}
	return repo.Locations(0, 0)
}

func (w *Manager) Product(dispensary, menuId, productId string) (*models.Product, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return &models.Product{}, fmt.Errorf("couldn't find dispensary by menu for product. Dispensary=%s, Location=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetProduct(menuId, productId)
}

func (w *Manager) Products(dispensary, menuId string) ([]*models.Product, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return []*models.Product{}, fmt.Errorf("couldn't find dispensary by menu for products. Dispensary=%s, Location=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetProducts(menuId)
}

func (w *Manager) ProductsForCategory(dispensary, menuId string, category models.Category) ([]*models.Product, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return []*models.Product{}, fmt.Errorf("couldn't find dispensary by menu for products for category. Dispensary=%s, Location=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetProductsForCategory(menuId, category)
}

func (w *Manager) ProductsForSubCategory(dispensary, menuId string, products []*models.Product, subCategory string) ([]*models.Product, error) {
	return w.TF.SubCategory(subCategory, products), nil
}

func (w *Manager) ProductsExcludingBrands(dispensary, menuId string, products []*models.Product, brands []string) ([]*models.Product, error) {
	return w.TF.NotBrands(brands, products), nil
}

func (w *Manager) ProductsForBrands(dispensary, menuId string, products []*models.Product, brands []string) ([]*models.Product, error) {
	return w.TF.Brands(brands, products), nil
}

func (w *Manager) ProductsForPriceRange(dispensary, menuId string, products []*models.Product, min, max float64) ([]*models.Product, error) {
	return w.TF.Price(min, max, products), nil
}

func (w *Manager) Categories(dispensary, menuId string) ([]*models.Category, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return []*models.Category{}, fmt.Errorf("couldn't find dispensary by menu for categories. Dispensary=%s, MenuId=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetCategories(menuId)
}

func (w *Manager) Terpenes(dispensary, menuId string) ([]*models.Terpene, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return []*models.Terpene{}, fmt.Errorf("couldn't find dispensary by menu for terpenes. Dispensary=%s, MenuId=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetTerpenes(menuId)
}

func (w *Manager) Cannabinoids(dispensary, menuId string) ([]*models.Cannabinoid, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return []*models.Cannabinoid{}, fmt.Errorf("couldn't find dispensary by menu for cannabinoids. Dispensary=%s, MenuId=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetCannabinoids(menuId)
}

func (w *Manager) Offers(dispensary, menuId string) ([]*models.Offer, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return []*models.Offer{}, fmt.Errorf("couldn't find dispensary by menu for offers. Dispensary=%s, MenuId=%s. Err: %v", dispensary, menuId, err)
	}
	return repo.GetOffers(menuId)
}

func (w *Manager) LogError(err error, context context.Context) {
	slog.InfoContext(context, err.Error())
}
