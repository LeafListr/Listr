package beyondhello

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Linkinlog/LeafListr/internal/client"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

const BeyondEndpoint = "https://vfm4x0n23a-dsn.algolia.net/1/indexes/menu-products-production/query"

var Headers = map[string][]string{
	"x-algolia-api-key":        {"b499e29eb7542dc373ec0254e007205d"},
	"x-algolia-application-id": {"VFM4X0N23A"},
}

func NewRepository() repository.Repository {
	return &Repository{
		C: NewHTTPClient(BeyondEndpoint, Headers),
	}
}

type Repository struct {
	C client.Client
}

func (r *Repository) Location(menuId string) (*models.Location, error) {
	return nil, nil
}

func (r *Repository) Locations(longitude, latitude float64) ([]*models.Location, error) {
	return nil, nil
}

func (r *Repository) GetProduct(menuId, productId string) (*models.Product, error) {
	return nil, nil
}

func (r *Repository) GetProducts(menuId string) ([]*models.Product, error) {
	mId, err := strconv.Atoi(menuId)
	if err != nil {
		return nil, err
	}
	query := MenuQuery(mId)
	return r.getProducts(query)
}

func (r *Repository) GetProductsForCategory(menuId string, category models.Category) ([]*models.Product, error) {
	return nil, nil
}

func (r *Repository) GetCategories(menuId string) ([]*models.Category, error) {
	return nil, nil
}

func (r *Repository) GetTerpenes(menuId string) ([]*models.Terpene, error) {
	return nil, nil
}

func (r *Repository) GetCannabinoids(menuId string) ([]*models.Cannabinoid, error) {
	return nil, nil
}

func (r *Repository) GetOffers(menuId string) ([]*models.Offer, error) {
	return nil, nil
}

func (r *Repository) getProducts(query string) ([]*models.Product, error) {
	resp, err := r.getResponse(query)
	if err != nil {
		return nil, err
	}

	return resp.translateProducts()
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
