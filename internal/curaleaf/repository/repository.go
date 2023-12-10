package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/models"

	curaClient "github.com/Linkinlog/LeafListr/internal/curaleaf/client"

	"github.com/Linkinlog/LeafListr/internal/client"
	"github.com/Linkinlog/LeafListr/internal/repository"
	"github.com/Linkinlog/LeafListr/internal/translation"
)

const (
	GqlEndpoint = "https://graph.curaleaf.com/api/curaql"
	GbgId       = "LMR124"
	Authority   = "graph.curaleaf.com"
	MenuType    = "MEDICAL"
)

var Headers = map[string][]string{"authority": {Authority}}

type Repository struct {
	C  client.Client
	T  translation.ClientTranslatable
	MC cache.Cacher
}

func NewRepository(c client.Client, translatable translation.ClientTranslatable, mc cache.Cacher) repository.Repository {
	return &Repository{
		C:  c,
		T:  translatable,
		MC: mc,
	}
}

func (r *Repository) Location(menuId string) (*models.Location, error) {
	queryKey := fmt.Sprintf("location-%s", menuId)
	location, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getLocation(menuId)
	})
	if err != nil {
		return &models.Location{}, err
	}

	return location.(*models.Location), nil
}

func (r *Repository) Locations(longitude, latitude float64) ([]*models.Location, error) {
	query := curaClient.AllLocationsQuery(longitude, latitude)
	queryKey := fmt.Sprintf("locations-%f-%f", longitude, latitude)

	locations, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getLocations(query)
	})
	if err != nil {
		return []*models.Location{}, err
	}

	return locations.([]*models.Location), nil
}

func (r *Repository) GetProduct(menuId, productId string) (*models.Product, error) {
	query := curaClient.ProductQuery(menuId, productId, "MEDICAL")
	queryKey := fmt.Sprintf("product-%s-%s", menuId, productId)

	product, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getProduct(query)
	})
	if err != nil {
		return &models.Product{}, err
	}

	return product.(*models.Product), nil
}

func (r *Repository) GetProducts(menuId string) ([]*models.Product, error) {
	query := curaClient.AllProductQuery(menuId, "MEDICAL")
	queryKey := fmt.Sprintf("products-%s", menuId)

	products, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getProducts(query)
	})
	if err != nil {
		return []*models.Product{}, err
	}

	return products.([]*models.Product), nil
}

func (r *Repository) GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error) {
	query := curaClient.AllProductForCategoryQuery(menuId, "MEDICAL", string(category))
	queryKey := fmt.Sprintf("products-%s-%s", menuId, string(category))

	products, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getProductsForCategory(query)
	})
	if err != nil {
		return []*models.Product{}, err
	}

	return products.([]*models.Product), nil
}

func (r *Repository) GetCategories(menuId string) ([]*models.Category, error) {
	query := curaClient.AllCategoriesQuery(menuId, "MEDICAL")
	queryKey := fmt.Sprintf("categories-%s", menuId)

	categories, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getCategories(query)
	})
	if err != nil {
		return []*models.Category{}, err
	}

	return categories.([]*models.Category), nil
}

func (r *Repository) GetTerpenes(menuId string) ([]*models.Terpene, error) {
	query := curaClient.AllProductQuery(menuId, "MEDICAL")
	queryKey := fmt.Sprintf("terpenes-%s", menuId)

	terpenes, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getTerpenes(query)
	})
	if err != nil {
		return []*models.Terpene{}, err
	}

	return terpenes.([]*models.Terpene), nil
}

func (r *Repository) GetCannabinoids(menuId string) ([]*models.Cannabinoid, error) {
	query := curaClient.AllProductQuery(menuId, "MEDICAL")
	queryKey := fmt.Sprintf("cannabinoids-%s", menuId)

	cannabinoids, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getCannabinoids(query)
	})
	if err != nil {
		return []*models.Cannabinoid{}, err
	}

	return cannabinoids.([]*models.Cannabinoid), nil
}

func (r *Repository) GetOffers(menuId string) ([]*models.Offer, error) {
	query := curaClient.AllOffersQuery(menuId, "MEDICAL")
	queryKey := fmt.Sprintf("offers-%s", menuId)

	offers, err := r.MC.GetOrRetrieve(queryKey, func() (any, error) {
		return r.getOffers(query)
	})
	if err != nil {
		return []*models.Offer{}, err
	}

	return offers.([]*models.Offer), nil
}

func (r *Repository) getLocation(locationId string) (*models.Location, error) {
	location := &models.Location{}
	locs, err := r.getLocations(curaClient.AllLocationsQuery(0, 0))
	if err != nil {
		return location, err
	}
	for _, l := range locs {
		if l.Id == locationId {
			location = l
		}
	}
	return location, nil
}

func (r *Repository) getLocations(query string) ([]*models.Location, error) {
	locationResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Location{}, err
	}

	return r.T.TranslateClientLocations(locationResp.Data.Dispensaries), nil
}

func (r *Repository) getProduct(query string) (*models.Product, error) {
	productResp, err := r.getResponse(query)
	if err != nil {
		return &models.Product{}, err
	}

	if productResp.Data.Product.Product.ID == "" {
		return &models.Product{}, repository.ResourceNotFound
	}

	return r.T.TranslateClientProduct(productResp.Data.Product.Product), nil
}

func (r *Repository) getProducts(query string) ([]*models.Product, error) {
	allProdResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Product{}, err
	}

	return r.T.TranslateClientProducts(allProdResp.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getProductsForCategory(query string) ([]*models.Product, error) {
	allProdForCatResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Product{}, err
	}
	return r.T.TranslateClientProducts(allProdForCatResp.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getCategories(query string) ([]*models.Category, error) {
	allCatsResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Category{}, err
	}
	return r.T.TranslateClientCategories(allCatsResp.Data.DispensaryMenu.AllFilters.Categories), nil
}

func (r *Repository) getTerpenes(query string) ([]*models.Terpene, error) {
	allProdResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Terpene{}, err
	}
	var mu sync.Mutex
	terpeneMap := make(map[string]*models.Terpene)
	for _, product := range allProdResp.Data.DispensaryMenu.Products {
		for _, terpene := range product.LabResults.Terpenes {
			mu.Lock()
			if _, exists := terpeneMap[terpene.Terpene.Name]; !exists {
				terpene.Value = 0
				terpeneMap[terpene.Terpene.Name] = r.T.TranslateClientTerpene(terpene)
			}
			mu.Unlock()
		}
	}

	ts := make([]*models.Terpene, 0)
	for _, t := range terpeneMap {
		ts = append(ts, t)
	}

	return ts, nil
}

func (r *Repository) getCannabinoids(query string) ([]*models.Cannabinoid, error) {
	allProdResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Cannabinoid{}, err
	}
	var mu sync.Mutex

	cannabinoidMap := make(map[string]*models.Cannabinoid)
	for _, product := range allProdResp.Data.DispensaryMenu.Products {
		for _, cannabinoid := range product.LabResults.Cannabinoids {
			mu.Lock()
			if _, exists := cannabinoidMap[cannabinoid.Cannabinoid.Name]; !exists {
				cannabinoidMap[cannabinoid.Cannabinoid.Name] = r.T.TranslateClientCannabinoid(cannabinoid)
			}
			mu.Unlock()
		}
	}

	cbs := make([]*models.Cannabinoid, 0)
	for _, cb := range cannabinoidMap {
		cbs = append(cbs, cb)
	}
	return cbs, nil
}

func (r *Repository) getOffers(query string) ([]*models.Offer, error) {
	allOfferResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Offer{}, err
	}
	return r.T.TranslateClientOffers(allOfferResp.Data.DispensaryMenu.Offers), nil
}

func (r *Repository) getResponse(query string) (curaClient.Response, error) {
	resp, err := r.C.Query(context.Background(), query, "POST")
	if err != nil {
		return curaClient.Response{}, err
	}

	if !json.Valid(resp) {
		return curaClient.Response{}, repository.InvalidJSONError
	}

	cResp := curaClient.Response{}
	err = json.Unmarshal(resp, &cResp)
	if err != nil {
		return curaClient.Response{}, err
	}

	return cResp, nil
}
