package workflow

import (
	"errors"
	"github.com/Linkinlog/LeafList/internal/client"
	"github.com/Linkinlog/LeafList/internal/factory"
	"github.com/Linkinlog/LeafList/internal/models"
	"github.com/Linkinlog/LeafList/internal/repository/curaleaf"
	"log"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Manager
type Manager interface {
	Locations(dispensary string, longitude, latitude float64) ([]*models.Brand, error)
	Product(dispensary, menuId, productId string) (*models.Product, error)
	AllProducts(dispensary, menuId string) ([]*models.Product, error)
	AllProductsForCategory(dispensary, menuId string, category models.Category) ([]*models.Product, error)
	Categories(dispensary, menuId string) ([]models.Category, error)
	Terpenes(dispensary, menuId string) ([]*models.Terpene, error)
	Cannabinoids(dispensary, menuId string) ([]*models.Cannabinoid, error)
	Offers(dispensary, menuId string) ([]*models.Offer, error)
	LogError(err error)
}

type Workflow struct {
	f factory.RepositoryFactory
}

func NewWorkflowManager() Manager {
	return &Workflow{
		f: factory.NewRepoFactory(),
	}
}

func (w *Workflow) Locations(dispensary string, longitude float64, latitude float64) ([]*models.Brand, error) {
	if dispensary != "curaleaf" {
		return nil, errors.New("TODO")
	}
	repo := curaleaf.NewRepository(client.NewHTTPClient(curaleaf.GqlEndpoint, curaleaf.Headers))
	return repo.GetLocations(longitude, latitude)
}

func (w *Workflow) Product(dispensary, menuId, productId string) (*models.Product, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetProduct(menuId, productId)
}

func (w *Workflow) AllProducts(dispensary, menuId string) ([]*models.Product, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetProducts(menuId)
}

func (w *Workflow) AllProductsForCategory(dispensary, menuId string, category models.Category) ([]*models.Product, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetProductsForCategory(menuId, category)
}

func (w *Workflow) Categories(dispensary string, menuId string) ([]models.Category, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetCategories(menuId)
}

func (w *Workflow) Terpenes(dispensary, menuId string) ([]*models.Terpene, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetTerpenes(menuId)
}

func (w *Workflow) Cannabinoids(dispensary, menuId string) ([]*models.Cannabinoid, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetCannabinoids(menuId)
}

func (w *Workflow) Offers(dispensary, menuId string) ([]*models.Offer, error) {
	repo, err := w.f.Find(dispensary, menuId)
	if err != nil {
		return nil, err
	}
	return repo.GetOffers(menuId)
}

func (w *Workflow) LogError(err error) {
	log.Printf("workflow error: %v", err)
}
