package curaleaf

import (
	"context"
	"encoding/json"
	"github.com/Linkinlog/LeafList/internal/client"
	"github.com/Linkinlog/LeafList/internal/client/curaleaf"
	"github.com/Linkinlog/LeafList/internal/repository/models"
)

type Repository struct {
	C client.Client
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

func (r *Repository) GetOffers(menuId string) ([]*models.Offer, error) {
	return r.getOffers(menuId)
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

	return translateTerpenes(os.Data.DispensaryMenu.Products), nil
}

func translateTerpenes(products []curaleaf.Product) []*models.Terpene {
	var mTerpenes []*models.Terpene
	for _, product := range products {
		for _, terp := range product.LabResults.Terpenes {
			tempTerp := &models.Terpene{
				Name:        terp.Terpene.Name,
				Description: terp.Terpene.Description,
				Value:       terp.Value,
			}
			mTerpenes = append(mTerpenes, tempTerp)
		}
	}
	return mTerpenes
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

	return translateOffers(os.Data.DispensaryMenu.Offers), nil
}

func translateOffers(offers []curaleaf.Offer) []*models.Offer {
	var mOffers []*models.Offer
	for _, o := range offers {
		tempOffer := &models.Offer{
			Id:          o.Id,
			Description: o.Title,
		}
		mOffers = append(mOffers, tempOffer)
	}
	return mOffers
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

	return translateCategories(cs.Data.DispensaryMenu.AllFilters.Categories), nil
}

func translateCategories(cs []curaleaf.Category) []models.Category {
	var mCs []models.Category
	for _, c := range cs {
		tempC := models.Category(c.Key)
		mCs = append(mCs, tempC)
	}
	return mCs
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

	return translateProducts(ps.Data.DispensaryMenu.Products), nil
}

func translateProducts(ps []curaleaf.Product) []*models.Product {
	if len(ps) <= 0 {
		return nil
	}
	var mPs []*models.Product
	for _, p := range ps {
		tempProd := &models.Product{
			Id:  p.ID,
			Ctg: models.Category(p.Category.DisplayName),
		}
		for _, v := range p.Variants {
			tempPrice := &models.Price{
				Variant:         v.Option,
				Total:           v.Price,
				DiscountedTotal: v.SpecialPrice,
			}
			tempProd.P = append(tempProd.P, tempPrice)
		}
		for _, t := range p.LabResults.Terpenes {
			tempTerp := &models.Terpene{
				Name:        t.Terpene.Name,
				Description: t.Terpene.Description,
				Value:       t.Value,
			}
			tempProd.T = append(tempProd.T, tempTerp)
		}
		for _, c := range p.LabResults.Cannabinoids {
			tempCanna := &models.Cannabinoid{
				Name:        c.Cannabinoid.Name,
				Description: c.Cannabinoid.Description,
				Value:       c.Value,
			}
			tempProd.C = append(tempProd.C, tempCanna)
		}
		mPs = append(mPs, tempProd)
	}
	return mPs
}
