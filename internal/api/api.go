package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	`strings`
	"time"

	_ "github.com/Linkinlog/LeafListr/docs"
	apiTranslator "github.com/Linkinlog/LeafListr/internal/api/translation"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/manager"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
	"github.com/rs/cors"

	"github.com/Linkinlog/LeafListr/internal/workflow"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Listr API
//	@version		0.1.0
//	@description	This is the Listr server for dispensary management.
//	@BasePath		/api/v1

const version = 1

type supportedLocationOptions string

const (
	Products     supportedLocationOptions = "products"
	Offers       supportedLocationOptions = "offers"
	Terpenes     supportedLocationOptions = "terpenes"
	Cannabinoids supportedLocationOptions = "cannabinoids"
	Categories   supportedLocationOptions = "categories"
)

type supportedDispensary string

const (
	Curaleaf supportedDispensary = "curaleaf"
)

type supportedDispensaryOptions string

const (
	Locations supportedDispensaryOptions = "locations"
)

type API struct {
	w                     workflow.Manager
	t                     translation.APITranslatable
	h                     http.Handler
	supportedDispensaries []supportedDispensary
}

func New(manager workflow.Manager) *API {
	router := chi.NewRouter()

	api := &API{
		w:                     manager,
		h:                     router,
		supportedDispensaries: []supportedDispensary{Curaleaf},
		t:                     apiTranslator.NewAPITranslator(),
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://*.dahlton.org", "http://localhost:*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	})
	router.Use(c.Handler)
	router.Use(RequestLogger)

	router.Get("/swagger/*", httpSwagger.WrapHandler)

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
	slog.Info("Listening on " + addr)
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

// @Summary		Get dispensary details
// @Description	Returns details of a specific dispensary
// @Tags			dispensaries
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string						true	"Dispensary ID"
// @Success		200				{array}	supportedDispensaryOptions	"Dispensary details"
// @Router			/dispensaries/{dispensaryId} [get].
func (a *API) handleDispensary(r http.ResponseWriter, _ *http.Request) {
	a.writeJson(r, nil, []supportedDispensaryOptions{Locations}, nil)
}

// @Summary		List supported dispensaries
// @Description	Returns a list of supported dispensaries
// @Tags			dispensaries
// @Accept			json
// @Produce		json
// @Success		200	{array}	supportedDispensary	"List of supported dispensaries"
// @Router			/dispensaries [get].
func (a *API) handleDispensaryListing(r http.ResponseWriter, _ *http.Request) {
	a.writeJson(r, nil, a.supportedDispensaries, nil)
}

// @Summary		Get location details
// @Description	Returns details of a specific location
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string	true	"Dispensary ID"
// @Param			locationId		path	string	true	"Location ID"
// @Success		200				{array}	string	"Location details"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId} [get].
func (a *API) handleLocation(res http.ResponseWriter, req *http.Request) {
	a.writeJson(res, req, []supportedLocationOptions{Products, Offers, Terpenes, Cannabinoids, Categories}, nil)
}

// @Summary		List locations for a dispensary
// @Description	Returns a list of locations for a specific dispensary
// @Tags			locations
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Success		200				{array}	models.Location	"List of locations"
// @Router			/dispensaries/{dispensaryId}/locations [get].
func (a *API) handleLocationListing(r http.ResponseWriter, req *http.Request) {
	dispensary, _, _ := params(req, "")
	locations, err := a.w.Locations(dispensary)
	a.writeJson(r, req, a.t.TranslateAPILocations(locations), err)
}

// @Summary		Get product details
// @Description	Returns details of a specific product
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path		string			true	"Dispensary ID"
// @Param			locationId		path		string			true	"Location ID"
// @Param			productId		path		string			true	"Product ID"
// @Success		200				{object}	models.Product	"Product details"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/products/{productId} [get].
func (a *API) handleProduct(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, productId := params(req, "productId")
	product, err := a.w.Product(dispensary, locationId, productId)
	a.writeJson(r, req, a.t.TranslateAPIProduct(product), err)
}

// @Summary		List products for a location
// @Description	Returns a list of products for a specific location
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Success		200				{array}	models.Product	"List of products"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/products [get].
func (a *API) handleProductListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	if category := req.URL.Query().Get("category"); category != "" {
		products, err := a.w.ProductsForCategory(dispensary, locationId, models.Category(category))
		a.writeJson(r, req, a.t.TranslateAPIProducts(products), err)
		return
	}
	products, err := a.w.Products(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPIProducts(products), err)
}

// @Summary		List offers for a location
// @Description	Returns a list of offers for a specific location
// @Tags			offers
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Success		200				{array}	models.Offer	"List of offers"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/offers [get].
func (a *API) handleOfferListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	offers, err := a.w.Offers(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPIOffers(offers), err)
}

// @Summary		List categories for a location
// @Description	Returns a list of categories for a specific location
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Success		200				{array}	models.Category	"List of categories"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/categories [get].
func (a *API) handleCategoryListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	categories, err := a.w.Categories(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPICategories(categories), err)
}

// @Summary		List terpenes for a location
// @Description	Returns a list of terpenes for a specific location
// @Tags			terpenes
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Success		200				{array}	models.Terpene	"List of terpenes"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/terpenes [get].
func (a *API) handleTerpeneListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	terpenes, err := a.w.Terpenes(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPITerpenes(terpenes), err)
}

// @Summary		List cannabinoids for a location
// @Description	Returns a list of cannabinoids for a specific location
// @Tags			cannabinoids
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string				true	"Dispensary ID"
// @Param			locationId		path	string				true	"Location ID"
// @Success		200				{array}	models.Cannabinoid	"List of cannabinoids"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/cannabinoids [get].
func (a *API) handleCannabinoidListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, _ := params(req, "")
	cannabinoids, err := a.w.Cannabinoids(dispensary, locationId)
	a.writeJson(r, req, a.t.TranslateAPICannabinoids(cannabinoids), err)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := make([]slog.Attr, len(r.Header))
		idx := 0
		for k, v := range r.Header {
			headers[idx] = slog.String(k, strings.Join(v, ","))
			idx = idx + 1
		}
		ctx := r.Context()
		slog.InfoContext(ctx, "request received",
			slog.Group("req",
				slog.String("URL", r.URL.String()),
				slog.String("ClientIp", r.RemoteAddr),
				slog.String("HTTPMethod", r.Method),
				slog.String("Referer", r.Referer()),
				slog.Attr{Key: "Header", Value: slog.GroupValue(headers...)},
			),
		)

		next.ServeHTTP(w, r)
	})
}

func params(req *http.Request, resource string) (dispensary, locationId, resourceId string) {
	dispensary = chi.URLParam(req, "dispensaryId")
	locationId = chi.URLParam(req, "locationId")
	if resource != "" {
		resourceId = chi.URLParam(req, resource)
	}
	return dispensary, locationId, resourceId
}
