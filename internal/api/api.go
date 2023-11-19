package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	apiTranslator "github.com/Linkinlog/LeafListr/internal/api/translation"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/manager"

	"github.com/Linkinlog/LeafListr/internal/translation"

	"github.com/Linkinlog/LeafListr/internal/workflow"
	"github.com/go-chi/chi/v5"
)

const version = 1

type API struct {
	w                     workflow.Manager
	t                     translation.APITranslatable
	h                     http.Handler
	supportedDispensaries []string
}

func New(manager workflow.Manager) *API {
	router := chi.NewRouter()

	api := &API{
		w:                     manager,
		h:                     router,
		supportedDispensaries: []string{"curaleaf"},
		t:                     apiTranslator.NewAPITranslator(),
	}

	router.Get(v("/dispensaries"), api.handleDispensaryListing)
	router.Get(v("/dispensaries/{dispensaryId}"), api.handleDispensary)

	router.Get(v("/dispensaries/{dispensaryId}/locations"), api.handleLocationListing)
	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}"), api.handleLocation)

	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}/products"), api.handleProductListing)
	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}/products/{productId}"), api.handleProduct)

	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}/offers"), api.handleOfferListing)

	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}/terpenes"), api.handleTerpeneListing)

	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}/cannabinoids"), api.handleCannabinoidListing)

	router.Get(v("/dispensaries/{dispensaryId}/locations/{locationId}/categories"), api.handleCategoryListing)

	api.h = router

	return api
}

func v(path string) string {
	return fmt.Sprintf("/api/v%d%s", version, path)
}

func ListenAndServe(addr string, timeout int64) error {
	a := New(manager.NewWorkflowManager())
	h := http.Server{
		Addr:              addr,
		Handler:           a.h,
		ReadHeaderTimeout: time.Duration(timeout) * time.Second,
	}
	fmt.Println("Listening on " + addr)
	return h.ListenAndServe()
}

func (a *API) writeJson(r http.ResponseWriter, req *http.Request, data any, err error) {
	if err != nil {
		a.handleError(r, req, err)
		return
	}
	resp, marshalErr := json.Marshal(&data)
	if marshalErr != nil {
		a.handleError(r, req, marshalErr)
	}

	r.Header().Set("Content-Type", "application/json")

	_, err = r.Write(resp)
	if err != nil {
		a.handleError(r, req, err)
	}
}

func (a *API) handleError(r http.ResponseWriter, req *http.Request, err error) {
	a.w.LogError(err, req.Context())
	r.WriteHeader(http.StatusInternalServerError)
}

func (a *API) handleDispensary(r http.ResponseWriter, _ *http.Request) {
	supportedPages := []string{
		"locations",
	}
	a.writeJson(r, nil, supportedPages, nil)
}

func (a *API) handleDispensaryListing(r http.ResponseWriter, _ *http.Request) {
	a.writeJson(r, nil, a.supportedDispensaries, nil)
}

func (a *API) handleLocation(res http.ResponseWriter, req *http.Request) {
	supportedLocationOptions := []string{ // TODO do we need to abstract this out? might be something to look into
		"products",
		"offers",
		"terpenes",
		"cannabinoids",
		"categories",
	}
	a.writeJson(res, req, supportedLocationOptions, nil)
}

func (a *API) handleLocationListing(r http.ResponseWriter, req *http.Request) {
	dispensary, _, _ := params(req, "")
	locations, err := a.w.Locations(dispensary)
	a.writeJson(r, req, a.t.TranslateAPILocations(locations), err)
}

func (a *API) handleProduct(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, productId := params(req, "productId")
	product, err := a.w.Product(dispensary, locationId, productId)
	a.writeJson(r, req, a.t.TranslateAPIProduct(product), err)
}

func (a *API) handleProductListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	products, err := a.w.Products(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPIProducts(products), err)
}

func (a *API) handleOfferListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	offers, err := a.w.Offers(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPIOffers(offers), err)
}

func (a *API) handleCategoryListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	categories, err := a.w.Categories(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPICategories(categories), err)
}

func (a *API) handleTerpeneListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	terpenes, err := a.w.Terpenes(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPITerpenes(terpenes), err)
}

func (a *API) handleCannabinoidListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	cannabinoids, err := a.w.Cannabinoids(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPICannabinoids(cannabinoids), err)
}

func params(req *http.Request, resource string) (dispensary, locationId, resourceId string) {
	dispensary = chi.URLParam(req, "dispensaryId")
	locationId = chi.URLParam(req, "locationId")
	if resource != "" {
		resourceId = chi.URLParam(req, resource)
	}
	return dispensary, locationId, resourceId
}
