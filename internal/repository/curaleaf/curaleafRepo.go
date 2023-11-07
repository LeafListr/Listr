package curaleaf

import (
	"context"
	"encoding/json"
	"github.com/Linkinlog/LeafList/internal/client"
	"github.com/Linkinlog/LeafList/internal/client/curaleaf"
	"github.com/Linkinlog/LeafList/internal/models"
	"github.com/Linkinlog/LeafList/internal/translation"
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

func NewRepository(c client.Client) *Repository {
	return &Repository{
		C: c,
		T: translation.NewTranslator(),
	}
}

func (r *Repository) GetLocations(longitude, latitude float64) ([]*models.Brand, error) {
	return r.getLocations(longitude, latitude)
}

func (r *Repository) GetProduct(menuId, productId string) (*models.Product, error) {
	return r.getProduct(menuId, productId)
}

func (r *Repository) GetProducts(menuId string) ([]*models.Product, error) {
	return r.getProducts(menuId, "")
}

func (r *Repository) GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error) {
	return r.getProducts(menuId, category)
}

func (r *Repository) GetCategories(menuId string) ([]models.Category, error) {
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

func (r *Repository) getLocations(longitude, latitude float64) ([]*models.Brand, error) {
	lResp, err := r.C.Query(context.Background(), curaleaf.AllLocationsQuery(longitude, latitude), "POST")

	if err != nil {
		return nil, err
	}

	ls := curaleaf.LocationResponse{}
	err = json.Unmarshal(lResp, &ls)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateLocations(ls.Data.Dispensaries), nil
}

func (r *Repository) getProduct(menuId, productId string) (*models.Product, error) {
	pResp, err := r.C.Query(context.Background(), curaleaf.ProductQuery(menuId, productId, "MEDICAL"), "POST")

	if err != nil {
		return nil, err
	}

	p := curaleaf.ProductResponse{}
	err = json.Unmarshal(pResp, &p)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateProduct(p.Data.Product.Product), nil
}

func (r *Repository) getProducts(menuId string, category models.Category) ([]*models.Product, error) {
	var query string
	if category != "" {
		query = curaleaf.AllProductForCategoryQuery(menuId, "MEDICAL", string(category))
	} else {
		query = curaleaf.AllProductQuery(menuId, "MEDICAL")
	}
	pResp, err := r.C.Query(context.Background(), query, "POST")

	if err != nil {
		return nil, err
	}

	ps := curaleaf.AllProductsResponse{}
	err = json.Unmarshal(pResp, &ps)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateProducts(ps.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getCategories(menuId string) ([]models.Category, error) {
	cResp, err := r.C.Query(context.Background(), curaleaf.AllCategoriesQuery(menuId, "MEDICAL"), "POST")

	if err != nil {
		return nil, err
	}

	cs := curaleaf.AllCategoriesResponse{}
	err = json.Unmarshal(cResp, &cs)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateCategories(cs.Data.DispensaryMenu.AllFilters.Categories), nil
}

func (r *Repository) getTerpenes(menuId string) ([]*models.Terpene, error) {
	tResp, err := r.C.Query(context.Background(), curaleaf.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	os := curaleaf.AllProductsResponse{}

	err = json.Unmarshal(tResp, &os)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateTerpenes(os.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getCannabinoids(menuId string) ([]*models.Cannabinoid, error) {
	tResp, err := r.C.Query(context.Background(), curaleaf.AllProductQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	os := curaleaf.AllProductsResponse{}

	err = json.Unmarshal(tResp, &os)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateCannabinoids(os.Data.DispensaryMenu.Products), nil
}

func (r *Repository) getOffers(menuId string) ([]*models.Offer, error) {
	oResp, err := r.C.Query(context.Background(), curaleaf.AllOffersQuery(menuId, "MEDICAL"), "POST")
	if err != nil {
		return nil, err
	}

	os := curaleaf.AllOffersResponse{}

	err = json.Unmarshal(oResp, &os)
	if err != nil {
		return nil, err
	}

	return r.T.TranslateOffers(os.Data.DispensaryMenu.Offers), nil
}
