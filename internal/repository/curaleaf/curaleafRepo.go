package curaleaf

import (
	"context"
	"encoding/json"

	"github.com/Linkinlog/LeafListr/internal/client"
	"github.com/Linkinlog/LeafListr/internal/client/curaleaf"
	"github.com/Linkinlog/LeafListr/internal/models"
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
	T translation.Translatable
}

func NewRepository(c client.Client) repository.Repository {
	return &Repository{
		C: c,
		T: translation.NewTranslator(),
	}
}

func (r *Repository) GetMenu(menuId string) (*models.Dispensary, error) {
	return r.getMenu(menuId)
}

func (r *Repository) GetMenus(longitude, latitude float64) ([]*models.Dispensary, error) {
	return r.getMenus(longitude, latitude)
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

func (r *Repository) getMenu(menuId string) (*models.Dispensary, error) {
	var disp *models.Dispensary
	disps, err := r.getMenus(0, 0)
	if err != nil {
		return nil, err
	}
	for _, l := range disps {
		if l.UniqueId == menuId {
			disp = l
		}
	}
	return disp, nil
}

func (r *Repository) getMenus(longitude, latitude float64) ([]*models.Dispensary, error) {
	locationResp, err := r.C.Query(context.Background(), curaleaf.AllLocationsQuery(longitude, latitude), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}
	err = json.Unmarshal(locationResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateLocations(cResp.Data.Dispensaries), nil
}

func (r *Repository) getProduct(menuId, productId string) (*models.Product, error) {
	productResp, err := r.C.Query(context.Background(), curaleaf.ProductQuery(menuId, productId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}
	err = json.Unmarshal(productResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateProduct(cResp.Data.Product.Product), nil
}

func (r *Repository) getProducts(menuId string) ([]*models.Product, error) {
	allProdResp, err := r.C.Query(context.Background(), curaleaf.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}
	err = json.Unmarshal(allProdResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateProducts(cResp.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getProductsForCategory(menuId string, category models.Category) ([]*models.Product, error) {
	allProdForCatResp, err := r.C.Query(context.Background(), curaleaf.AllProductForCategoryQuery(menuId, "MEDICAL", string(category)), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}
	err = json.Unmarshal(allProdForCatResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateProducts(cResp.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getCategories(menuId string) ([]*models.Category, error) {
	allCatsResp, err := r.C.Query(context.Background(), curaleaf.AllCategoriesQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}
	err = json.Unmarshal(allCatsResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateCategories(cResp.Data.DispensaryMenu.AllFilters.Categories), nil
}

func (r *Repository) getTerpenes(menuId string) ([]*models.Terpene, error) {
	allProdResp, err := r.C.Query(context.Background(), curaleaf.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}

	err = json.Unmarshal(allProdResp, &cResp)
	if err != nil {
		return nil, err
	}

	terpeneMap := make(map[string]*models.Terpene)
	for _, product := range cResp.Data.DispensaryMenu.Products {
		for _, terpene := range product.LabResults.Terpenes {
			if _, exists := terpeneMap[terpene.Terpene.Name]; !exists {
				terpeneMap[terpene.Terpene.Name] = r.T.TranslateTerpene(&terpene)
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
	allProdResp, err := r.C.Query(context.Background(), curaleaf.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}

	err = json.Unmarshal(allProdResp, &cResp)
	if err != nil {
		return nil, err
	}

	cannabinoidMap := make(map[string]*models.Cannabinoid)
	for _, product := range cResp.Data.DispensaryMenu.Products {
		for _, cannabinoid := range product.LabResults.Cannabinoids {
			if _, exists := cannabinoidMap[cannabinoid.Cannabinoid.Name]; !exists {
				cannabinoidMap[cannabinoid.Cannabinoid.Name] = r.T.TranslateCannabinoid(&cannabinoid)
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
	allOfferResp, err := r.C.Query(context.Background(), curaleaf.AllOffersQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	cResp := curaleaf.Response{}

	err = json.Unmarshal(allOfferResp, &cResp)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateOffers(cResp.Data.DispensaryMenu.Offers), nil
}
