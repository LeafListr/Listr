package workflow

import (
	"context"
	"log/slog"

	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Manager
type Manager interface {
	MenuManager
	ProductManager
	CategoryManager
	TerpeneManager
	CannabinoidManager
	OfferManager
	ErrorManager
}

type MenuManager interface {
	Menu(dispensary, menuId string) (*models.Dispensary, error)
	Menus(dispensary string) ([]*models.Dispensary, error)
}

type ProductManager interface {
	Product(dispensary, menuId, productId string) (*models.Product, error)
	Products(dispensary, menuId string) ([]*models.Product, error)
	ProductsForCategory(dispensary, menuId string, category models.Category) ([]*models.Product, error)
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

type Workflow struct {
	f factory.RepositoryFactory
}

func NewWorkflowManager() Manager {
	return &Workflow{
		f: factory.NewRepoFactory(),
	}
}

func (w *Workflow) Menu(dispensary, menuId string) (*models.Dispensary, error) {
	repo, err := w.f.FindByDispensary(dispensary)
	if err != nil {
		return nil, err
	}
	return repo.GetMenu(menuId)
}

func (w *Workflow) Menus(dispensary string) ([]*models.Dispensary, error) {
	repo, err := w.f.FindByDispensary(dispensary)
	if err != nil {
		return nil, err
	}
	return repo.GetMenus(0, 0)
}

func (w *Workflow) Product(dispensary, menuId, productId string) (*models.Product, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetProduct(menuId, productId)
}

func (w *Workflow) Products(dispensary, menuId string) ([]*models.Product, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetProducts(menuId)
}

func (w *Workflow) ProductsForCategory(dispensary, menuId string, category models.Category) ([]*models.Product, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetProductsForCategory(menuId, category)
}

func (w *Workflow) Categories(dispensary, menuId string) ([]*models.Category, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetCategories(menuId)
}

func (w *Workflow) Terpenes(dispensary, menuId string) ([]*models.Terpene, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetTerpenes(menuId)
}

func (w *Workflow) Cannabinoids(dispensary, menuId string) ([]*models.Cannabinoid, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetCannabinoids(menuId)
}

func (w *Workflow) Offers(dispensary, menuId string) ([]*models.Offer, error) {
	repo, err := w.f.FindByDispensaryMenu(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetOffers(menuId)
}

func (w *Workflow) LogError(err error, context context.Context) {
	slog.ErrorContext(context, err.Error())
}
