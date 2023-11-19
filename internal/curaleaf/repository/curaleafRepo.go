package repository

import (
	"context"
	"encoding/json"
	"fmt"

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
	C client.Client
	T translation.ClientTranslatable
}

func NewRepository(c client.Client, translatable translation.ClientTranslatable) repository.Repository {
	return &Repository{
		C: c,
		T: translatable,
	}
}

func (r *Repository) Location(menuId string) (*models.Location, error) {
	return r.getMenu(menuId)
}

func (r *Repository) Locations(longitude, latitude float64) ([]*models.Location, error) {
	return r.getLocations(longitude, latitude)
}

func (r *Repository) GetProduct(menuId, productId string) (*models.Product, error) {
	return r.getProduct(menuId, productId)
}

func (r *Repository) GetProducts(menuId string) ([]*models.Product, error) {
	return r.getProducts(menuId)
}

func (r *Repository) GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error) {
	return r.getProductsForCategory(menuId, category)
}

func (r *Repository) GetCategories(menuId string) ([]*models.Category, error) {
	return r.getCategories(menuId)
}

func (r *Repository) GetTerpenes(menuId string) ([]*models.Terpene, error) {
	return r.getTerpenes(menuId)
}

func (r *Repository) GetCannabinoids(menuId string) ([]*models.Cannabinoid, error) {
	return r.getCannabinoids(menuId)
}

func (r *Repository) GetOffers(menuId string) ([]*models.Offer, error) {
	return r.getOffers(menuId)
}

func (r *Repository) getMenu(menuId string) (*models.Location, error) {
	var location *models.Location
	locs, err := r.getLocations(0, 0)
	if err != nil {
		return nil, err
	}
	for _, l := range locs {
		if l.Name == menuId {
			location = l
		}
	}
	return location, nil
}

func (r *Repository) getLocations(longitude, latitude float64) ([]*models.Location, error) {
	locationResp, err := r.C.Query(context.Background(), curaClient.AllLocationsQuery(longitude, latitude), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}
	err = json.Unmarshal(locationResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateClientLocations(cResp.Data.Dispensaries), nil
}

func (r *Repository) getProduct(menuId, productId string) (*models.Product, error) {
	productResp, err := r.C.Query(context.Background(), curaClient.ProductQuery(menuId, productId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}
	err = json.Unmarshal(productResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateClientProduct(cResp.Data.Product.Product), nil
}

func (r *Repository) getProducts(menuId string) ([]*models.Product, error) {
	allProdResp, err := r.C.Query(context.Background(), curaClient.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}
	err = json.Unmarshal(allProdResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateClientProducts(cResp.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getProductsForCategory(menuId string, category models.Category) ([]*models.Product, error) {
	allProdForCatResp, err := r.C.Query(context.Background(), curaClient.AllProductForCategoryQuery(menuId, "MEDICAL", string(category)), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}
	err = json.Unmarshal(allProdForCatResp, &cResp)
	if err != nil {
		return nil, fmt.Errorf("could not get products for category: %s, menu=%s. Err=%v", category, menuId, err)
	}

	return r.T.TranslateClientProducts(cResp.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getCategories(menuId string) ([]*models.Category, error) {
	allCatsResp, err := r.C.Query(context.Background(), curaClient.AllCategoriesQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}
	err = json.Unmarshal(allCatsResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateClientCategories(cResp.Data.DispensaryMenu.AllFilters.Categories), nil
}

func (r *Repository) getTerpenes(menuId string) ([]*models.Terpene, error) {
	allProdResp, err := r.C.Query(context.Background(), curaClient.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}

	err = json.Unmarshal(allProdResp, &cResp)
	if err != nil {
		return nil, err
	}

	terpeneMap := make(map[string]*models.Terpene)
	for _, product := range cResp.Data.DispensaryMenu.Products {
		for _, terpene := range product.LabResults.Terpenes {
			if _, exists := terpeneMap[terpene.Terpene.Name]; !exists {
				terpeneMap[terpene.Terpene.Name] = r.T.TranslateClientTerpene(&terpene)
			}
		}
	}

	var ts []*models.Terpene
	for _, t := range terpeneMap {
		ts = append(ts, t)
	}

	return ts, nil
}

func (r *Repository) getCannabinoids(menuId string) ([]*models.Cannabinoid, error) {
	allProdResp, err := r.C.Query(context.Background(), curaClient.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}

	err = json.Unmarshal(allProdResp, &cResp)
	if err != nil {
		return nil, err
	}

	cannabinoidMap := make(map[string]*models.Cannabinoid)
	for _, product := range cResp.Data.DispensaryMenu.Products {
		for _, cannabinoid := range product.LabResults.Cannabinoids {
			if _, exists := cannabinoidMap[cannabinoid.Cannabinoid.Name]; !exists {
				cannabinoidMap[cannabinoid.Cannabinoid.Name] = r.T.TranslateClientCannabinoid(&cannabinoid)
			}
		}
	}

	var cbs []*models.Cannabinoid
	for _, cb := range cannabinoidMap {
		cbs = append(cbs, cb)
	}
	return cbs, nil
}

func (r *Repository) getOffers(menuId string) ([]*models.Offer, error) {
	allOfferResp, err := r.C.Query(context.Background(), curaClient.AllOffersQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaClient.Response{}

	err = json.Unmarshal(allOfferResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateClientOffers(cResp.Data.DispensaryMenu.Offers), nil
}
