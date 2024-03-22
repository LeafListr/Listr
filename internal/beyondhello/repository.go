package beyondhello

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

const (
	BeyondEndpoint          = "https://vfm4x0n23a-dsn.algolia.net/1/indexes/menu-products-production/query"
	LocationListingEndpoint = "https://beyond-hello.com/wp-json/wp/v2/online_menu?per_page=100"
)

var Headers = map[string][]string{
	"x-algolia-api-key":        {"b499e29eb7542dc373ec0254e007205d"},
	"x-algolia-application-id": {"VFM4X0N23A"},
	"Content-Type":             {"application/json"},
}

func NewRepository(menuId string, recreational bool) repository.Repository {
	menuType := "MEDICAL"
	if recreational {
		menuType = "RECREATIONAL"
	}
	return &Repository{
		menuType: menuType,
		menuId:   menuId,
	}
}

type Repository struct {
	menuType string
	menuId   string
}

func (r *Repository) Location() (*models.Location, error) {
	locs, err := r.Locations(0, 0)
	if err != nil {
		return nil, err
	}
	for _, l := range locs {
		if l.Id == r.menuId {
			return l, nil
		}
	}
	return nil, errors.New("no location found")
}

func (r *Repository) Locations(longitude, latitude float64) ([]*models.Location, error) {
	resp, err := r.locationResponse()
	if err != nil {
		return nil, err
	}

	translated, tErr := resp.translateLocations()
	if tErr != nil {
		return nil, tErr
	}

	locs := make([]*models.Location, 0)

	for _, l := range translated {
		for _, t := range l.LocationTypes {
			if strings.EqualFold(t, r.menuType) {
				locs = append(locs, l)
			}
		}
	}

	return locs, nil
}

func (r *Repository) GetProduct(productId string) (*models.Product, error) {
	mId, err := strconv.Atoi(r.menuId)
	if err != nil {
		return nil, err
	}
	pId, convErr := strconv.Atoi(productId)
	if convErr != nil {
		return nil, convErr
	}
	query := productQuery(mId, pId)
	resp, respErr := r.productResponse(query)
	if respErr != nil {
		return nil, respErr
	}

	prods, transErr := resp.translateProducts()
	if transErr != nil {
		return nil, transErr
	}
	if len(prods) > 0 {
		return prods[0], nil
	}

	return nil, errors.New("no product found")
}

func (r *Repository) GetProducts() ([]*models.Product, error) {
	mId, err := strconv.Atoi(r.menuId)
	if err != nil {
		return nil, err
	}
	query := menuQuery(mId)
	return r.getProducts(query)
}

func (r *Repository) GetProductsForCategory(category string) ([]*models.Product, error) {
	categoryValid := false
	validCategories, catErr := r.GetCategories()
	if catErr != nil {
		return nil, catErr
	}
	for _, c := range validCategories {
		if strings.EqualFold(string(c), string(category)) {
			categoryValid = true
		}
	}

	if !categoryValid {
		return nil, repository.InvalidCategoryError
	}

	mId, err := strconv.Atoi(r.menuId)
	if err != nil {
		return nil, err
	}
	query := menuQueryWithCategory(mId, string(category))
	return r.getProducts(query)
}

func (r *Repository) GetSubcategories(category string) ([]string, error) {
	return []string{
		"",
	}, nil
}

func (r *Repository) GetCategories() ([]string, error) {
	return []string{
		"flower",
		"vape",
		"extract",
		"edible",
		"tincture",
		"gear",
		"topical",
	}, nil
}

func (r *Repository) GetTerpenes() ([]*models.Terpene, error) {
	terps := map[string]struct{}{}
	products, err := r.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		for _, t := range p.T {
			terps[t.Name] = struct{}{}
		}
	}
	terpSlice := make([]*models.Terpene, 0, len(terps))
	for t := range terps {
		terpSlice = append(terpSlice, &models.Terpene{
			Name: t,
		})
	}
	return terpSlice, nil
}

func (r *Repository) GetCannabinoids() ([]*models.Cannabinoid, error) {
	canns := map[string]struct{}{}
	products, err := r.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		for _, c := range p.C {
			canns[c.Name] = struct{}{}
		}
	}
	cannSlice := make([]*models.Cannabinoid, 0, len(canns))
	for c := range canns {
		cannSlice = append(cannSlice, &models.Cannabinoid{
			Name: c,
		})
	}
	return cannSlice, nil
}

func (r *Repository) GetOffers() ([]*models.Offer, error) {
	// TODO this is awful
	mId, err := strconv.Atoi(r.menuId)
	if err != nil {
		return nil, err
	}
	resp, err := r.productResponse(menuQuery(mId))
	if err != nil {
		return nil, err
	}

	offerMap := map[string]struct{}{}
	for _, p := range resp.Hits {
		if p.SpecialTitle != "" {
			offerMap[p.SpecialTitle] = struct{}{}
		}
	}
	offers := make([]*models.Offer, 0)
	for o := range offerMap {
		offers = append(offers, &models.Offer{
			Id:          o,
			Description: o,
		})
	}
	return offers, nil
}

func (r *Repository) getProducts(query string) ([]*models.Product, error) {
	resp, err := r.productResponse(query)
	if err != nil {
		return nil, err
	}

	return resp.translateProducts()
}

func (r *Repository) productResponse(query string) (ProductResponse, error) {
	c := NewHTTPClient(Endpoint(BeyondEndpoint), Headers)
	resp, err := c.Query(context.Background(), query, "POST")
	if err != nil {
		return ProductResponse{}, err
	}

	if !json.Valid(resp) {
		return ProductResponse{}, repository.InvalidJSONError
	}

	cResp := ProductResponse{}
	err = json.Unmarshal(resp, &cResp)
	if err != nil {
		return ProductResponse{}, err
	}

	return cResp, nil
}

func (r *Repository) locationResponse() (LocationResponse, error) {
	c := NewHTTPClient(Endpoint(LocationListingEndpoint), Headers)
	resp, err := c.Query(context.Background(), "", "GET")
	if err != nil {
		return LocationResponse{}, err
	}

	if !json.Valid(resp) {
		return LocationResponse{}, repository.InvalidJSONError
	}

	lResp := LocationResponse{}
	err = json.Unmarshal(resp, &lResp)
	if err != nil {
		return LocationResponse{}, err
	}

	return lResp, nil
}
