package workflow

import (
	"context"
	"fmt"
	"log/slog"
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
	F  factory.RepositoryFactory
	TF transformation.Filterer
	S  transformation.Sorter
	C  cache.Cacher
}

func NewWorkflow(f factory.RepositoryFactory, tf transformation.Filterer, s transformation.Sorter, c cache.Cacher) Workflow {
	return &DefaultWorkflow{
		F:  f,
		TF: tf,
		S:  s,
		C:  c,
	}
}

func (w *DefaultWorkflow) Location(dispensary, menuId, menuType string) (*models.Location, error) {
	repo, err := w.F.FindByDispensary(dispensary, menuType)
	if err != nil {
		return &models.Location{}, fmt.Errorf("couldn't find dispensary by menu for location. Dispensary=%s, MenuId=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("location-%s-%s-%s", dispensary, menuId, menuType)
	location, cacheErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.Location(menuId)
	})
	if cacheErr != nil {
		return &models.Location{}, cacheErr
	}
	return location.(*models.Location), nil
}

func (w *DefaultWorkflow) Locations(dispensary, menuType string) ([]*models.Location, error) {
	repo, err := w.F.FindByDispensary(dispensary, menuType)
	if err != nil {
		return []*models.Location{}, fmt.Errorf("couldn't find dispensary by menu for locations. Dispensary=%s. Err: %v", dispensary, err)
	}
	queryKey := fmt.Sprintf("locations-%d-%d-%s-%s", 0, 0, dispensary, menuType)
	locations, cacheErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.Locations(0, 0)
	})
	if cacheErr != nil {
		return []*models.Location{}, cacheErr
	}
	return locations.([]*models.Location), nil
}

func (w *DefaultWorkflow) Product(dispensary, menuId, menuType, productId string) (*models.Product, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return &models.Product{}, fmt.Errorf("couldn't find dispensary by menu for product. Dispensary=%s, Location=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("product-%s-%s-%s-%s", dispensary, menuId, productId, menuType)
	product, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProduct(menuId, productId)
	})
	if retErr != nil {
		return &models.Product{}, retErr
	}

	return product.(*models.Product), nil
}

func (w *DefaultWorkflow) Products(dispensary, menuId, menuType string) ([]*models.Product, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return []*models.Product{}, fmt.Errorf("couldn't find dispensary by menu for products. Dispensary=%s, Location=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("products-%s-%s-%s", dispensary, menuId, menuType)
	products, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProducts(menuId)
	})
	if retErr != nil {
		return []*models.Product{}, retErr
	}

	return products.([]*models.Product), nil
}

func (w *DefaultWorkflow) ProductsInCategory(dispensary, menuId, menuType string, category models.Category) ([]*models.Product, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return []*models.Product{}, fmt.Errorf("couldn't find dispensary by menu for products for category. Dispensary=%s, Location=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("products-%s-%s-%s-%s", dispensary, menuId, string(category), menuType)
	products, retErr := w.C.GetOrRetrieve(queryKey, ShortTTL, func() (any, error) {
		return repo.GetProductsForCategory(menuId, category)
	})
	if retErr != nil {
		return []*models.Product{}, retErr
	}

	return products.([]*models.Product), nil
}

func (w *DefaultWorkflow) Categories(dispensary, menuId, menuType string) ([]models.Category, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return []models.Category{}, fmt.Errorf("couldn't find dispensary by menu for categories. Dispensary=%s, MenuId=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("categories-%s-%s-%s", dispensary, menuId, menuType)
	categories, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetCategories(menuId)
	})
	if retErr != nil {
		return []models.Category{}, retErr
	}

	return categories.([]models.Category), nil
}

func (w *DefaultWorkflow) Terpenes(dispensary, menuId, menuType string) ([]*models.Terpene, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return []*models.Terpene{}, fmt.Errorf("couldn't find dispensary by menu for terpenes. Dispensary=%s, MenuId=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("terpenes-%s-%s-%s", dispensary, menuId, menuType)
	terpenes, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetTerpenes(menuId)
	})
	if retErr != nil {
		return []*models.Terpene{}, retErr
	}

	return terpenes.([]*models.Terpene), nil
}

func (w *DefaultWorkflow) Cannabinoids(dispensary, menuId, menuType string) ([]*models.Cannabinoid, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return []*models.Cannabinoid{}, fmt.Errorf("couldn't find dispensary by menu for cannabinoids. Dispensary=%s, MenuId=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("cannabinoids-%s-%s-%s", dispensary, menuId, menuType)
	cannabinoids, retErr := w.C.GetOrRetrieve(queryKey, DayTTL, func() (any, error) {
		return repo.GetCannabinoids(menuId)
	})
	if retErr != nil {
		return []*models.Cannabinoid{}, retErr
	}

	return cannabinoids.([]*models.Cannabinoid), nil
}

func (w *DefaultWorkflow) Offers(dispensary, menuId, menuType string) ([]*models.Offer, error) {
	repo, err := w.F.FindByDispensaryMenu(dispensary, menuId, menuType)
	if err != nil {
		return []*models.Offer{}, fmt.Errorf("couldn't find dispensary by menu for offers. Dispensary=%s, MenuId=%s, MenuType=%s. Err: %v", dispensary, menuId, menuType, err)
	}
	queryKey := fmt.Sprintf("offers-%s-%s-%s", dispensary, menuId, menuType)
	offers, retErr := w.C.GetOrRetrieve(queryKey, HalfDayTTL, func() (any, error) {
		return repo.GetOffers(menuId)
	})
	if retErr != nil {
		return []*models.Offer{}, retErr
	}

	return offers.([]*models.Offer), nil
}

func (w *DefaultWorkflow) ProductsForSubCategory(_, _, _ string, products []*models.Product, subCategory string) ([]*models.Product, error) {
	return w.TF.SubCategory(subCategory, products), nil
}

func (w *DefaultWorkflow) ProductsExcludingBrands(_, _, _ string, products []*models.Product, brands []string) ([]*models.Product, error) {
	return w.TF.NotBrands(brands, products), nil
}

func (w *DefaultWorkflow) ProductsForBrands(_, _, _ string, products []*models.Product, brands []string) ([]*models.Product, error) {
	return w.TF.Brands(brands, products), nil
}

func (w *DefaultWorkflow) ProductsForVariants(_, _, _ string, products []*models.Product, variants []string) ([]*models.Product, error) {
	return w.TF.Variants(variants, products), nil
}

func (w *DefaultWorkflow) ProductsForPriceRange(_, _, _ string, products []*models.Product, min, max float64) ([]*models.Product, error) {
	return w.TF.Price(min, max, products), nil
}

func (w *DefaultWorkflow) ProductsIncludingTerms(_, _, _ string, products []*models.Product, includes []string) ([]*models.Product, error) {
	return w.TF.IncludingTerms(includes, products), nil
}

func (w *DefaultWorkflow) ProductsExcludingTerms(_, _, _ string, products []*models.Product, excludes []string) ([]*models.Product, error) {
	return w.TF.ExcludingTerms(excludes, products), nil
}

func (w *DefaultWorkflow) SortProductsByPriceAsc(_, _, _ string, products []*models.Product) {
	w.S.PriceAsc(products)
}

func (w *DefaultWorkflow) SortProductsByPriceDesc(_, _, _ string, products []*models.Product) {
	w.S.PriceDesc(products)
}

func (w *DefaultWorkflow) SortProductsByTHCAsc(_, _, _ string, products []*models.Product) {
	w.S.THCAsc(products)
}

func (w *DefaultWorkflow) SortProductsByTHCDesc(_, _, _ string, products []*models.Product) {
	w.S.THCDesc(products)
}

func (w *DefaultWorkflow) SortProductsByTop3Terps(_, _, _ string, products []*models.Product, terps [3]string) {
	w.S.Top3Terps(products, terps)
}

func (w *DefaultWorkflow) LogError(err error, context context.Context) {
	slog.InfoContext(context, err.Error())
}
