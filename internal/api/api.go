package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Linkinlog/LeafList/internal/workflow"
	"github.com/go-chi/chi/v5"
)

const version = 1

type API struct {
	w workflow.Manager
	h http.Handler
}

func New(w workflow.Manager) *API {
	r := chi.NewRouter()

	a := &API{
		w: w,
	}

	r.Get(v("/dispensaries"), a.handleDispensaryListing)
	r.Get(v("/dispensaries/{dispensaryId}"), a.handleDispensary)

	r.Get(v("/dispensaries/{dispensaryId}/menus"), a.handleMenuListing)
	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}"), a.handleMenu)

	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}/products"), a.handleProductListing)
	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}/products/{productId}"), a.handleProduct)

	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}/offers"), a.handleOfferListing)

	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}/terpenes"), a.handleTerpeneListing)

	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}/cannabinoids"), a.handleCannabinoidListing)

	r.Get(v("/dispensaries/{dispensaryId}/menus/{menuId}/categories"), a.handleCategoryListing)

	a.h = r

	return a
}

func v(path string) string {
	return fmt.Sprintf("/api/v%d%s", version, path)
}

func ListenAndServe(addr string) error {
	a := New(workflow.NewWorkflowManager())
	fmt.Println("Listening on " + addr)
	return http.ListenAndServe(addr, a.h)
}

func (a *API) writeJson(r http.ResponseWriter, data any, err error) {
	if err != nil {
		a.handleError(r, err)
		return
	}
	resp, marshalErr := json.Marshal(&data)
	if marshalErr != nil {
		a.handleError(r, marshalErr)
	}
	r.Header().Set("Content-Type", "application/json")
	_, err = r.Write(resp)
	if err != nil {
		a.handleError(r, err)
	}
}

func (a *API) handleError(r http.ResponseWriter, err error) {
	a.w.LogError(err)
	r.WriteHeader(http.StatusInternalServerError)
}

func (a *API) handleDispensary(r http.ResponseWriter, _ *http.Request) {
	supportedPages := []string{
		"menus",
	}
	a.writeJson(r, supportedPages, nil)
}

func (a *API) handleDispensaryListing(r http.ResponseWriter, _ *http.Request) {
	supportedDispensaries := []string{
		"curaleaf",
	}
	a.writeJson(r, supportedDispensaries, nil)
}

func (a *API) handleMenu(r http.ResponseWriter, _ *http.Request) {
	// do we need to abstract this out? might be something to look into
	supportedMenuOptions := []string{
		"products",
		"offers",
		"terpenes",
		"cannabinoids",
		"categories",
	}
	a.writeJson(r, supportedMenuOptions, nil)
}

func (a *API) handleMenuListing(r http.ResponseWriter, req *http.Request) {
	dispensary, _, _ := params(req, "")
	menus, err := a.w.Menus(dispensary)
	a.writeJson(r, menus, err)
}

func (a *API) handleProduct(r http.ResponseWriter, req *http.Request) {
	dispensary, menuId, productId := params(req, "productId")
	product, err := a.w.Product(dispensary, menuId, productId)
	a.writeJson(r, product, err)
}

func (a *API) handleProductListing(r http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	products, err := a.w.Products(dispensary, menuId)
	a.writeJson(r, products, err)
}

func (a *API) handleOfferListing(r http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	offers, err := a.w.Offers(dispensary, menuId)
	a.writeJson(r, offers, err)
}

func (a *API) handleCategoryListing(r http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	categories, err := a.w.Categories(dispensary, menuId)
	a.writeJson(r, categories, err)
}

func (a *API) handleTerpeneListing(r http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	terpenes, err := a.w.Terpenes(dispensary, menuId)
	a.writeJson(r, terpenes, err)
}

func (a *API) handleCannabinoidListing(r http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	cannabinoids, err := a.w.Cannabinoids(dispensary, menuId)
	a.writeJson(r, cannabinoids, err)
}

func params(req *http.Request, resource string) (dispensary, menuId, resourceId string) {
	dispensary = chi.URLParam(req, "dispensaryId")
	menuId = chi.URLParam(req, "menuId")
	if resource != "" {
		resourceId = chi.URLParam(req, resource)
	}
	return dispensary, menuId, resourceId
}
