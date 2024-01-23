package workflow

import (
	"context"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Workflow
type Workflow interface {
	LocationManager
	ProductManager
	CategoryManager
	TerpeneManager
	CannabinoidManager
	OfferManager
	ErrorManager
	ProductFilter
	ProductSorter
}

type WorkflowParams struct {
	dispensary   string
	menuId       string
	recreational bool
}

func NewWorkflowParams(dispensary, menuId string, recreational bool) WorkflowParams {
	return WorkflowParams{
		dispensary:   dispensary,
		menuId:       menuId,
		recreational: recreational,
	}
}

type LocationManager interface {
	Location(w WorkflowParams) (*models.Location, error)
	Locations(w WorkflowParams) ([]*models.Location, error)
}

type ProductManager interface {
	Product(w WorkflowParams, productId string) (*models.Product, error)
	Products(w WorkflowParams) ([]*models.Product, error)
	ProductsInCategory(w WorkflowParams, category string) ([]*models.Product, error)
}

type ProductFilter interface {
	Filter(WorkflowParams, *transformation.FilterParams, []*models.Product) ([]*models.Product, error)
	ProductsForSubCategory(w WorkflowParams, products []*models.Product, subCategory string) ([]*models.Product, error)
	ProductsForBrands(w WorkflowParams, products []*models.Product, brands []string) ([]*models.Product, error)
	ProductsExcludingBrands(w WorkflowParams, products []*models.Product, brands []string) ([]*models.Product, error)
	ProductsForVariants(w WorkflowParams, products []*models.Product, variants []string) ([]*models.Product, error)
	ProductsIncludingTerms(w WorkflowParams, products []*models.Product, includes []string) ([]*models.Product, error)
	ProductsExcludingTerms(w WorkflowParams, products []*models.Product, excludes []string) ([]*models.Product, error)
	ProductsForPriceRange(w WorkflowParams, products []*models.Product, min, max float64) ([]*models.Product, error)
}

type ProductSorter interface {
	Sort(WorkflowParams, *transformation.SortParams, []*models.Product) error
	SortProductsByPriceAsc(w WorkflowParams, products []*models.Product)
	SortProductsByPriceDesc(w WorkflowParams, products []*models.Product)
	SortProductsByTHCAsc(w WorkflowParams, products []*models.Product)
	SortProductsByTHCDesc(w WorkflowParams, products []*models.Product)
	SortProductsByTop3Terps(w WorkflowParams, products []*models.Product, terps [3]string)
}

type CategoryManager interface {
	Categories(w WorkflowParams) ([]string, error)
}

type TerpeneManager interface {
	Terpenes(w WorkflowParams) ([]*models.Terpene, error)
}

type CannabinoidManager interface {
	Cannabinoids(w WorkflowParams) ([]*models.Cannabinoid, error)
}

type OfferManager interface {
	Offers(w WorkflowParams) ([]*models.Offer, error)
}

type ErrorManager interface {
	LogError(err error, context context.Context)
}
