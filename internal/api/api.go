package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/transformation"

	"github.com/Linkinlog/LeafListr/internal/workflow"

	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	_ "github.com/Linkinlog/LeafListr/docs"
	apiTranslator "github.com/Linkinlog/LeafListr/internal/api/translation"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/translation"
	"github.com/rs/cors"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Listr API
//	@version		0.1.2
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
	Curaleaf supportedDispensary = "Curaleaf"
	Beyond   supportedDispensary = "Beyond-Hello"
)

type supportedDispensaryOptions string

const (
	Locations supportedDispensaryOptions = "locations"
)

type API struct {
	w                     workflow.Workflow
	t                     translation.APITranslatable
	h                     http.Handler
	supportedDispensaries []supportedDispensary
}

func New() *API {
	router := chi.NewRouter()

	bsp := honeycomb.NewBaggageSpanProcessor()

	_, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(bsp),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}

	f := factory.NewRepoFactory()
	tf := transformation.NewFilterer()
	s := transformation.NewSorterer()
	cc := cache.NewCache()
	api := &API{
		w:                     workflow.NewWorkflow(f, tf, s, cc),
		h:                     router,
		supportedDispensaries: []supportedDispensary{Curaleaf, Beyond},
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
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.Parse()

	logLevel := slog.LevelInfo
	if verbose {
		logLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	slog.SetDefault(logger)

	a := New()
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
// @Param			menu_type		query	string			true	"Menu type"
// @Success		200				{array}	models.Location	"List of locations"
// @Router			/dispensaries/{dispensaryId}/locations [get].
func (a *API) handleLocationListing(r http.ResponseWriter, req *http.Request) {
	dispensary, _, menuType, _ := params(req, "")
	locations, err := a.w.Locations(dispensary, menuType)
	a.writeJson(r, req, a.t.TranslateAPILocations(locations), err)
}

// @Summary		Get product details
// @Description	Returns details of a specific product
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path		string			true	"Dispensary ID"
// @Param			locationId		path		string			true	"Location ID"
// @Param			menu_type		query		string			true	"Menu type"
// @Param			productId		path		string			true	"Product ID"
// @Success		200				{object}	models.Product	"Product details"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/products/{productId} [get].
func (a *API) handleProduct(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, menuType, productId := params(req, "productId")
	product, err := a.w.Product(dispensary, locationId, menuType, productId)
	a.writeJson(r, req, a.t.TranslateAPIProduct(product), err)
}

// @Summary		List products for a location
// @Description	Returns a list of products for a specific location
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			menu_type		query	string			true	"Menu type"
// @Param			category		query	string			true	"Category"
// @Param			sub				query	string			false	"Sub Category"
// @Param			min_price		query	number			false	"Minimum price"
// @Param			max_price		query	number			false	"Maximum price"
// @Param			brands			query	string			false	"Brands to include"
// @Param			not_brands		query	string			false	"Brands to exclude"
// @Param			variants		query	string			false	"Variants to include"
// @Param			sort			query	string			false	"Sort products"	Enums(price_asc, price_desc)
// @Param			terp1			query	string			false	"Most important terpene"
// @Param			terp2			query	string			false	"Second most important terpene"
// @Param			terp3			query	string			false	"Third most important terpene"
// @Success		200				{array}	models.Product	"List of products"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/products [get].
func (a *API) handleProductListing(r http.ResponseWriter, req *http.Request) {
	var products []*models.Product
	var err error
	dispensary, locationId, menuType, _ := params(req, "")
	if category := req.URL.Query().Get("category"); category != "" {
		products, err = a.w.ProductsForCategory(dispensary, locationId, menuType, models.Category(category))
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}
	if subCategory := subCategoryFilter(req); subCategory != "" {
		products, err = a.w.ProductsForSubCategory(dispensary, locationId, menuType, products, subCategory)
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}
	if hasPriceFilters(req) {
		minPrice, maxPrice := priceFilters(req)
		products, err = a.w.ProductsForPriceRange(dispensary, locationId, menuType, products, minPrice, maxPrice)
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}
	if brands := brandFilters(req); brands != "" {
		products, err = a.w.ProductsForBrands(dispensary, locationId, menuType, products, strings.Split(brands, ","))
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}
	if notBrands := notBrandFilters(req); notBrands != "" {
		products, err = a.w.ProductsExcludingBrands(dispensary, locationId, menuType, products, strings.Split(notBrands, ","))
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}
	if variants := variantFilters(req); variants != "" {
		products, err = a.w.ProductsForVariants(dispensary, locationId, menuType, products, strings.Split(variants, ","))
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}

	if products == nil {
		products, err = a.w.Products(dispensary, locationId, menuType)
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}

	if sortTop3Terps(req) {
		slog.Debug("Sorting products by top 3 terps")
		a.w.SortProductsByTop3Terps(dispensary, locationId, menuType, products, top3Terps(req))
	}

	if sortPriceAsc(req) {
		slog.Debug("Sorting products by price asc")
		a.w.SortProductsByPriceAsc(dispensary, locationId, menuType, products)
	}
	if sortPriceDesc(req) {
		slog.Debug("Sorting products by price desc")
		a.w.SortProductsByPriceDesc(dispensary, locationId, menuType, products)
	}

	a.writeJson(r, req, a.t.TranslateAPIProducts(products), err)
}

// @Summary		List offers for a location
// @Description	Returns a list of offers for a specific location
// @Tags			offers
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			menu_type		query	string			true	"Menu type"
// @Success		200				{array}	models.Offer	"List of offers"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/offers [get].
func (a *API) handleOfferListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, menuType, _ := params(req, "")
	offers, err := a.w.Offers(dispensary, locationId, menuType)
	a.writeJson(r, req, a.t.TranslateAPIOffers(offers), err)
}

// @Summary		List categories for a location
// @Description	Returns a list of categories for a specific location
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			menu_type		query	string			true	"Menu type"
// @Success		200				{array}	models.Category	"List of categories"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/categories [get].
func (a *API) handleCategoryListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, menuType, _ := params(req, "")
	categories, err := a.w.Categories(dispensary, locationId, menuType)
	a.writeJson(r, req, a.t.TranslateAPICategories(categories), err)
}

// @Summary		List terpenes for a location
// @Description	Returns a list of terpenes for a specific location
// @Tags			terpenes
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			menu_type		query	string			true	"Menu type"
// @Success		200				{array}	models.Terpene	"List of terpenes"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/terpenes [get].
func (a *API) handleTerpeneListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, menuType, _ := params(req, "")
	terpenes, err := a.w.Terpenes(dispensary, locationId, menuType)
	a.writeJson(r, req, a.t.TranslateAPITerpenes(terpenes), err)
}

// @Summary		List cannabinoids for a location
// @Description	Returns a list of cannabinoids for a specific location
// @Tags			cannabinoids
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string				true	"Dispensary ID"
// @Param			locationId		path	string				true	"Location ID"
// @Param			menu_type		query	string				true	"Menu type"
// @Success		200				{array}	models.Cannabinoid	"List of cannabinoids"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/cannabinoids [get].
func (a *API) handleCannabinoidListing(r http.ResponseWriter, req *http.Request) {
	dispensary, locationId, menuType, _ := params(req, "")
	cannabinoids, err := a.w.Cannabinoids(dispensary, locationId, menuType)
	a.writeJson(r, req, a.t.TranslateAPICannabinoids(cannabinoids), err)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tr := otel.Tracer("LeafListr-API")
		ctx, span := tr.Start(ctx, "HTTPRequest",
			trace.WithAttributes(
				attribute.String("url", r.URL.String()),
				attribute.String("client_ip", r.RemoteAddr),
				attribute.String("http_method", r.Method),
				attribute.String("referer", r.Referer()),
			),
		)
		defer span.End()

		headers := make([]slog.Attr, len(r.Header))
		idx := 0
		for k, v := range r.Header {
			if len(v) > 0 {
				span.SetAttributes(attribute.String(fmt.Sprintf("http.header.%s", k), strings.Join(v, ", ")))
				headers[idx] = slog.String(k, strings.Join(v, ","))
				idx = idx + 1
			}
		}
		slog.DebugContext(ctx, "request received",
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

func params(req *http.Request, resource string) (dispensary, locationId, menuType, resourceId string) {
	dispensary = chi.URLParam(req, "dispensaryId")
	locationId = chi.URLParam(req, "locationId")
	menuType = req.URL.Query().Get("menu_type")
	if resource != "" {
		resourceId = chi.URLParam(req, resource)
	}
	return dispensary, locationId, menuType, resourceId
}

func subCategoryFilter(req *http.Request) string {
	return req.URL.Query().Get("sub")
}

func hasPriceFilters(req *http.Request) bool {
	return req.URL.Query().Get("min_price") != "" || req.URL.Query().Get("max_price") != ""
}

func priceFilters(req *http.Request) (float64, float64) {
	var minP, maxP float64
	var err error
	if minStr := req.URL.Query().Get("min_price"); minStr != "" {
		minP, err = strconv.ParseFloat(minStr, 64)
		if err != nil {
			slog.Debug("Error parsing min price", err)
			return 0, 0
		}
	}
	if maxStr := req.URL.Query().Get("max_price"); maxStr != "" {
		maxP, err = strconv.ParseFloat(maxStr, 64)
		if err != nil {
			slog.Debug("Error parsing max price", err)
			return 0, 0
		}
	}
	return minP, maxP
}

func brandFilters(req *http.Request) string {
	return req.URL.Query().Get("brands")
}

func notBrandFilters(req *http.Request) string {
	return req.URL.Query().Get("not_brands")
}

func variantFilters(req *http.Request) string {
	return req.URL.Query().Get("variants")
}

func sortPriceAsc(req *http.Request) bool {
	return req.URL.Query().Get("sort") == "price_asc"
}

func sortPriceDesc(req *http.Request) bool {
	return req.URL.Query().Get("sort") == "price_desc"
}

func top3Terps(req *http.Request) [3]string {
	var terps [3]string
	if terp1 := req.URL.Query().Get("terp1"); terp1 != "" {
		terps[0] = terp1
	}
	if terp2 := req.URL.Query().Get("terp2"); terp2 != "" {
		terps[1] = terp2
	}
	if terp3 := req.URL.Query().Get("terp3"); terp3 != "" {
		terps[2] = terp3
	}
	return terps
}

func sortTop3Terps(req *http.Request) bool {
	return top3Terps(req) != [3]string{}
}
