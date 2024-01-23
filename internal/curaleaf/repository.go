package curaleaf

import (
	"context"
	"encoding/json"
	"strings"
	"sync"

	"github.com/Linkinlog/LeafListr/internal/client"

	"github.com/Linkinlog/LeafListr/internal/models"

	"github.com/Linkinlog/LeafListr/internal/repository"
)

const (
	GqlEndpoint = "https://graph.curaleaf.com/api/curaql"
)

type Repository struct {
	C        client.Client
	T        *ClientTranslator
	menuType string
	menuId   string
}

func NewRepository(c client.Client, translator *ClientTranslator, menuId string, recreational bool) repository.Repository {
	menuType := "MEDICAL"
	if recreational {
		menuType = "RECREATIONAL"
	}
	return &Repository{
		C:        c,
		T:        translator,
		menuType: menuType,
		menuId:   menuId,
	}
}

func (r *Repository) Location() (*models.Location, error) {
	query := AllLocationsQuery(0, 0)
	return r.getLocation(query, r.menuId)
}

func (r *Repository) Locations(longitude, latitude float64) ([]*models.Location, error) {
	query := AllLocationsQuery(longitude, latitude)
	return r.getLocations(query)
}

func (r *Repository) GetProduct(productId string) (*models.Product, error) {
	query := ProductQuery(r.menuId, productId, r.menuType)
	return r.getProduct(query)
}

func (r *Repository) GetProducts() ([]*models.Product, error) {
	query := AllProductQuery(r.menuId, r.menuType)
	return r.getProducts(query)
}

func (r *Repository) GetProductsForCategory(category string) ([]*models.Product, error) {
	query := AllProductForCategoryQuery(r.menuId, r.menuType, string(category))
	return r.getProductsForCategory(query)
}

func (r *Repository) GetCategories() ([]string, error) {
	query := AllCategoriesQuery(r.menuId, r.menuType)
	return r.getCategories(query)
}

func (r *Repository) GetTerpenes() ([]*models.Terpene, error) {
	query := AllProductQuery(r.menuId, r.menuType)
	return r.getTerpenes(query)
}

func (r *Repository) GetCannabinoids() ([]*models.Cannabinoid, error) {
	query := AllProductQuery(r.menuId, r.menuType)
	return r.getCannabinoids(query)
}

func (r *Repository) GetOffers() ([]*models.Offer, error) {
	query := AllOffersQuery(r.menuId, r.menuType)
	return r.getOffers(query)
}

func (r *Repository) getLocation(query, locationId string) (*models.Location, error) {
	location := &models.Location{}
	locs, err := r.getLocations(query)
	if err != nil {
		return location, err
	}
	for _, l := range locs {
		if l.Id == locationId && contains(l.LocationTypes, r.menuType) {
			location = l
		}
	}
	return location, nil
}

func (r *Repository) getLocations(query string) ([]*models.Location, error) {
	locs := make([]*models.Location, 0)
	locationResp, err := r.getResponse(query)
	if err != nil {
		return []*models.Location{}, err
	}

	for _, l := range r.T.TranslateClientLocations(locationResp.Data.Dispensaries) {
		if contains(l.LocationTypes, r.menuType) {
			locs = append(locs, l)
		}
	}

	return locs, nil
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

func (r *Repository) getCategories(query string) ([]string, error) {
	allCatsResp, err := r.getResponse(query)
	if err != nil {
		return []string{}, err
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

func (r *Repository) getResponse(query string) (Response, error) {
	resp, err := r.C.Query(context.Background(), query, "POST")
	if err != nil {
		return Response{}, err
	}

	if !json.Valid(resp) {
		return Response{}, repository.InvalidJSONError
	}

	cResp := Response{}
	err = json.Unmarshal(resp, &cResp)
	if err != nil {
		return Response{}, err
	}

	return cResp, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}
