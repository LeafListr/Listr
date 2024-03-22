package workflow

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/repository"

	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
)

const (
	ShortTTL   = 5 * time.Minute
	DayTTL     = 24 * time.Hour
	HalfDayTTL = 12 * time.Hour
)

func NewWorkflowParams(dispensary, menuId string, recreational bool) WorkflowParams {
	return WorkflowParams{
		Dispensary:   dispensary,
		MenuId:       menuId,
		Recreational: recreational,
	}
}

type WorkflowParams struct {
	Dispensary   string
	MenuId       string
	Recreational bool
}

func NewWorkflow(c cache.Cacher) *Workflow {
	return &Workflow{
		C: c,
	}
}

type Workflow struct {
	C cache.Cacher
}

func (w *Workflow) Locations(wp WorkflowParams, repo repository.LocationRepository) ([]*models.Location, error) {
	queryKey := fmt.Sprintf("locations-%d-%d-%s-%t", 0, 0, wp.Dispensary, wp.Recreational)
	locations, cacheErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.Locations(0, 0)
	})
	if cacheErr != nil {
		return []*models.Location{}, cacheErr
	}
	locs := locations.([]*models.Location)
	if len(locs) != 0 {
		sort.SliceStable(locs, func(i, j int) bool {
			return locs[i].Name < locs[j].Name
		})
	}
	return locs, nil
}

func (w *Workflow) Product(wp WorkflowParams, productId string, repo repository.ProductRepository) (*models.Product, error) {
	queryKey := fmt.Sprintf("product-%s-%s-%s-%t", wp.Dispensary, wp.MenuId, productId, wp.Recreational)
	product, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProduct(productId)
	})
	if retErr != nil {
		return &models.Product{}, retErr
	}

	return product.(*models.Product), nil
}

func (w *Workflow) Products(wp WorkflowParams, repo repository.ProductRepository) ([]*models.Product, error) {
	queryKey := fmt.Sprintf("products-%s-%s-%t", wp.Dispensary, wp.MenuId, wp.Recreational)
	products, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProducts()
	})
	if retErr != nil {
		return []*models.Product{}, retErr
	}

	return products.([]*models.Product), nil
}

func (w *Workflow) ProductsInCategory(wp WorkflowParams, category string, repo repository.ProductRepository) ([]*models.Product, error) {
	queryKey := fmt.Sprintf("products-%s-%s-%s-%t", wp.Dispensary, wp.MenuId, string(category), wp.Recreational)
	products, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProductsForCategory(category)
	})
	if retErr != nil {
		return []*models.Product{}, retErr
	}

	return products.([]*models.Product), nil
}

func (w *Workflow) Categories(wp WorkflowParams, repo repository.CategoryRepository) ([]string, error) {
	queryKey := fmt.Sprintf("categories-%s-%s-%t", wp.Dispensary, wp.MenuId, wp.Recreational)
	categories, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetCategories()
	})
	if retErr != nil {
		return []string{}, retErr
	}
	cats := categories.([]string)
	if len(cats) != 0 {
		sort.SliceStable(cats, func(i, j int) bool {
			return cats[i] < cats[j]
		})
	}

	return cats, nil
}

func (w *Workflow) Subcategories(wp WorkflowParams, category string, repo repository.CategoryRepository) ([]string, error) {
	queryKey := fmt.Sprintf("subcategories-%s-%s-%t", wp.Dispensary, wp.MenuId, wp.Recreational)
	categories, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetSubcategories(category)
	})
	if retErr != nil {
		return []string{}, retErr
	}
	cats := categories.([]string)
	if len(cats) != 0 {
		sort.SliceStable(cats, func(i, j int) bool {
			return cats[i] < cats[j]
		})
	}

	return cats, nil
}

func (w *Workflow) Terpenes(wp WorkflowParams, repo repository.TerpeneRepository) ([]*models.Terpene, error) {
	queryKey := fmt.Sprintf("terpenes-%s-%s-%t", wp.Dispensary, wp.MenuId, wp.Recreational)
	terpenes, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetTerpenes()
	})
	if retErr != nil {
		return []*models.Terpene{}, retErr
	}

	terps := terpenes.([]*models.Terpene)
	if len(terps) != 0 {
		sort.SliceStable(terps, func(i, j int) bool {
			return terps[i].Name < terps[j].Name
		})
	}

	return terps, nil
}

func (w *Workflow) Cannabinoids(wp WorkflowParams, repo repository.CannabinoidRepository) ([]*models.Cannabinoid, error) {
	f := factory.NewRepoFactory(wp.Dispensary, wp.MenuId, wp.Recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Cannabinoid{}, fmt.Errorf("couldn't find dispensary by menu for cannabinoids. Dispensary=%s, MenuId=%s. Err: %v", wp.Dispensary, wp.MenuId, err)
	}
	queryKey := fmt.Sprintf("cannabinoids-%s-%s-%t", wp.Dispensary, wp.MenuId, wp.Recreational)
	cannabinoids, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetCannabinoids()
	})
	if retErr != nil {
		return []*models.Cannabinoid{}, retErr
	}

	return cannabinoids.([]*models.Cannabinoid), nil
}

func (w *Workflow) Offers(wp WorkflowParams, repo repository.OfferRepository) ([]*models.Offer, error) {
	f := factory.NewRepoFactory(wp.Dispensary, wp.MenuId, wp.Recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Offer{}, fmt.Errorf("couldn't find dispensary by menu for offers. Dispensary=%s, MenuId=%s. Err: %v", wp.Dispensary, wp.MenuId, err)
	}
	queryKey := fmt.Sprintf("offers-%s-%s-%t", wp.Dispensary, wp.MenuId, wp.Recreational)
	offers, retErr := w.C.GetOrRetrieve(queryKey, HalfDayTTL, func() (any, error) {
		return repo.GetOffers()
	})
	if retErr != nil {
		return []*models.Offer{}, retErr
	}

	return offers.([]*models.Offer), nil
}

func (w *Workflow) Filter(fp *transformation.FilterParams, products []*models.Product) ([]*models.Product, error) {
	f := transformation.NewFilterer(fp)

	return f.Filter(products), nil
}

func (w *Workflow) Sort(sp *transformation.SortParams, products []*models.Product) error {
	s := transformation.NewSorterer(sp)
	s.Sort(products)

	return nil
}

func (w *Workflow) LogError(err error, context context.Context) {
	slog.InfoContext(context, err.Error())
}

func (w *Workflow) RepoFromFactory(params WorkflowParams) (repository.Repository, error) {
	f := factory.NewRepoFactory(params.Dispensary, params.MenuId, params.Recreational)
	if params.MenuId == "" {
		repo, err := f.FindByDispensary()
		if err != nil {
			return nil, err
		}
		return repo, nil
	}
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return nil, err
	}
	return repo, nil
}
