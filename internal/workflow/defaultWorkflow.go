package workflow

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"

	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
)

const (
	ShortTTL   = 5 * time.Minute
	DayTTL     = 24 * time.Hour
	HalfDayTTL = 12 * time.Hour
)

type DefaultWorkflow struct {
	C cache.Cacher
}

func NewWorkflow(c cache.Cacher) Workflow {
	return &DefaultWorkflow{
		C: c,
	}
}

func (w *DefaultWorkflow) Location(wp WorkflowParams) (*models.Location, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensary()
	if err != nil {
		return &models.Location{}, fmt.Errorf("couldn't find dispensary by menu for location. Dispensary=%s, MenuId=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("location-%s-%s-%t", wp.dispensary, wp.menuId, wp.recreational)
	location, cacheErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.Location()
	})
	if cacheErr != nil {
		return &models.Location{}, cacheErr
	}
	return location.(*models.Location), nil
}

func (w *DefaultWorkflow) Locations(wp WorkflowParams) ([]*models.Location, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensary()
	if err != nil {
		return []*models.Location{}, fmt.Errorf("couldn't find dispensary by menu for locations. Dispensary=%s. Err: %v", wp.dispensary, err)
	}
	queryKey := fmt.Sprintf("locations-%d-%d-%s-%t", 0, 0, wp.dispensary, wp.recreational)
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

func (w *DefaultWorkflow) Product(wp WorkflowParams, productId string) (*models.Product, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return &models.Product{}, fmt.Errorf("couldn't find dispensary by menu for product. Dispensary=%s, Location=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("product-%s-%s-%s-%t", wp.dispensary, wp.menuId, productId, wp.recreational)
	product, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProduct(productId)
	})
	if retErr != nil {
		return &models.Product{}, retErr
	}

	return product.(*models.Product), nil
}

func (w *DefaultWorkflow) Products(wp WorkflowParams) ([]*models.Product, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Product{}, fmt.Errorf("couldn't find dispensary by menu for products. Dispensary=%s, Location=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("products-%s-%s-%t", wp.dispensary, wp.menuId, wp.recreational)
	products, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProducts()
	})
	if retErr != nil {
		return []*models.Product{}, retErr
	}

	return products.([]*models.Product), nil
}

func (w *DefaultWorkflow) ProductsInCategory(wp WorkflowParams, category string) ([]*models.Product, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Product{}, fmt.Errorf("couldn't find dispensary by menu for products for category. Dispensary=%s, Location=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("products-%s-%s-%s-%t", wp.dispensary, wp.menuId, string(category), wp.recreational)
	products, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProductsForCategory(category)
	})
	if retErr != nil {
		return []*models.Product{}, retErr
	}

	return products.([]*models.Product), nil
}

func (w *DefaultWorkflow) Categories(wp WorkflowParams) ([]string, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []string{}, fmt.Errorf("couldn't find dispensary by menu for categories. Dispensary=%s, MenuId=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("categories-%s-%s-%t", wp.dispensary, wp.menuId, wp.recreational)
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

func (w *DefaultWorkflow) Terpenes(wp WorkflowParams) ([]*models.Terpene, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Terpene{}, fmt.Errorf("couldn't find dispensary by menu for terpenes. Dispensary=%s, MenuId=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("terpenes-%s-%s-%t", wp.dispensary, wp.menuId, wp.recreational)
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

func (w *DefaultWorkflow) Cannabinoids(wp WorkflowParams) ([]*models.Cannabinoid, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Cannabinoid{}, fmt.Errorf("couldn't find dispensary by menu for cannabinoids. Dispensary=%s, MenuId=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("cannabinoids-%s-%s-%t", wp.dispensary, wp.menuId, wp.recreational)
	cannabinoids, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetCannabinoids()
	})
	if retErr != nil {
		return []*models.Cannabinoid{}, retErr
	}

	return cannabinoids.([]*models.Cannabinoid), nil
}

func (w *DefaultWorkflow) Offers(wp WorkflowParams) ([]*models.Offer, error) {
	f := factory.NewRepoFactory(wp.dispensary, wp.menuId, wp.recreational)
	repo, err := f.FindByDispensaryMenu()
	if err != nil {
		return []*models.Offer{}, fmt.Errorf("couldn't find dispensary by menu for offers. Dispensary=%s, MenuId=%s. Err: %v", wp.dispensary, wp.menuId, err)
	}
	queryKey := fmt.Sprintf("offers-%s-%s-%t", wp.dispensary, wp.menuId, wp.recreational)
	offers, retErr := w.C.GetOrRetrieve(queryKey, HalfDayTTL, func() (any, error) {
		return repo.GetOffers()
	})
	if retErr != nil {
		return []*models.Offer{}, retErr
	}

	return offers.([]*models.Offer), nil
}

func (w *DefaultWorkflow) Filter(fp *transformation.FilterParams, products []*models.Product) ([]*models.Product, error) {
	f := transformation.NewFilterer(fp)

	return f.Filter(products), nil
}

func (w *DefaultWorkflow) Sort(sp *transformation.SortParams, products []*models.Product) error {
	s := transformation.NewSorterer(sp)
	s.Sort(products)

	return nil
}

func (w *DefaultWorkflow) LogError(err error, context context.Context) {
	slog.InfoContext(context, err.Error())
}
