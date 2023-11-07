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

//func ListenAndServeTLS(addr string) error {
//	// todo
//	fmt.Println("TODO")
//	return nil
//}

func (a *API) handleError(r http.ResponseWriter, err error) {
	a.w.LogError(err)
	r.WriteHeader(http.StatusInternalServerError)
}

func (a *API) writeJson(r http.ResponseWriter, j any) {
	resp, err := json.Marshal(&j)
	if err != nil {
		a.handleError(r, err)
	}
	r.Header().Set("Content-Type", "application/json")
	_, err = r.Write(resp)
	if err != nil {
		a.handleError(r, err)
	}
}

func (a *API) handleDispensary(respw http.ResponseWriter, _ *http.Request) {
	supportedPages := []string{
		"menus",
	}
	a.writeJson(respw, supportedPages)
}

func (a *API) handleDispensaryListing(respw http.ResponseWriter, _ *http.Request) {
	supportedDispensaries := []string{
		"curaleaf",
	}
	a.writeJson(respw, supportedDispensaries)
}

func (a *API) handleMenu(respw http.ResponseWriter, _ *http.Request) {
	supportedDispensaries := []string{
		"products",
		"offers",
		"terpenes",
		"cannabinoids",
		"categories",
	}
	a.writeJson(respw, supportedDispensaries)
}

func (a *API) handleMenuListing(respw http.ResponseWriter, req *http.Request) {
	dispensary, _, _ := params(req, "")
	locations, err := a.w.Locations(dispensary, 0, 0)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, locations)
}

func (a *API) handleProduct(respw http.ResponseWriter, req *http.Request) {
	dispensary, menuId, productId := params(req, "productId")
	product, err := a.w.Product(dispensary, menuId, productId)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, product)
}

func (a *API) handleProductListing(respw http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	products, err := a.w.AllProducts(dispensary, menuId)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, products)
}

func (a *API) handleOfferListing(respw http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	offers, err := a.w.Offers(dispensary, menuId)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, offers)
}

func (a *API) handleCategoryListing(respw http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	categories, err := a.w.Categories(dispensary, menuId)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, categories)
}

func (a *API) handleTerpeneListing(respw http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	terpenes, err := a.w.Terpenes(dispensary, menuId)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, terpenes)
}

func (a *API) handleCannabinoidListing(respw http.ResponseWriter, req *http.Request) {
	dispensary, menuId, _ := params(req, "")
	cannabinoids, err := a.w.Cannabinoids(dispensary, menuId)
	if err != nil {
		a.handleError(respw, err)
	}
	a.writeJson(respw, cannabinoids)
}

func params(req *http.Request, resource string) (dispensary, menuId, resourceId string) {
	dispensary = chi.URLParam(req, "dispensaryId")
	menuId = chi.URLParam(req, "menuId")
	if resource != "" {
		resourceId = chi.URLParam(req, resource)
	}
	return dispensary, menuId, resourceId
}
