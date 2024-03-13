package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/transformation"
	"github.com/Linkinlog/LeafListr/internal/views/assets"
	"github.com/Linkinlog/LeafListr/internal/views/handlers"

	"github.com/Linkinlog/LeafListr/internal/workflow"

	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	_ "github.com/Linkinlog/LeafListr/docs"
	apiTranslator "github.com/Linkinlog/LeafListr/internal/api/translation"
	"github.com/Linkinlog/LeafListr/internal/translation"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Listr API
//	@version		0.2.1
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

const (
	Curaleaf = "Curaleaf"
	Beyond   = "Beyond-Hello"
)

type supportedDispensaryOptions string

const (
	Locations supportedDispensaryOptions = "locations"
)

type API struct {
	w                     workflow.Workflow
	t                     translation.APITranslatable
	h                     http.Handler
	supportedDispensaries []string
}

func ListenAndServe(addr string, timeout int) error {
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

func New() *API {
	router := chi.NewRouter()
	router.Use(corsMiddleware)

	bsp := honeycomb.NewBaggageSpanProcessor()

	_, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(bsp),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}

	cc := cache.NewCache()
	api := &API{
		w:                     workflow.NewWorkflow(cc),
		h:                     router,
		supportedDispensaries: []string{Curaleaf, Beyond},
		t:                     apiTranslator.NewAPITranslator(),
	}

	router.Use(RequestLogger)
	router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.Files()))))

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Route(appendVer("/api"), api.jsonRoutes)

	router.Get("/", api.handleLandingPage)
	router.Route("/views", api.htmlRoutes)

	api.h = router

	return api
}

func appendVer(path string) string {
	return fmt.Sprintf("%s/v%d", path, version)
}

func (a *API) jsonRoutes(r chi.Router) {
	r.Get("/dispensaries", a.handleDispensaryListing)
	r.Get("/dispensaries/{dispensaryId}", a.handleDispensary)

	r.Get("/dispensaries/{dispensaryId}/locations", a.handleLocationListing)
	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}", a.handleLocation)

	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}/products", a.handleProductListing)
	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}/products/{productId}", a.handleProduct)

	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}/offers", a.handleOfferListing)

	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}/terpenes", a.handleTerpeneListing)

	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}/cannabinoids", a.handleCannabinoidListing)

	r.Get("/dispensaries/{dispensaryId}/locations/{locationId}/categories", a.handleCategoryListing)
}

func (a *API) handleLandingPage(w http.ResponseWriter, r *http.Request) {
	h := handlers.NewHtmlHandler(a.supportedDispensaries, a.w)
	h.LandingPage(w, r)
}

func (a *API) htmlRoutes(r chi.Router) {
	r.Get("/locations", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Locations)
	r.Get("/nav", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Navigation)
	r.Get("/categories", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Categories)
	r.Get("/products", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Products)
	r.Get("/offers-and-search", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).OffersAndSearch)
	r.Get("/filters", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Filters)
	r.Get("/sorters", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Sorters)
	r.Get("/offers", handlers.NewHtmlHandler(a.supportedDispensaries, a.w).Offers)
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
		a.w.LogError(err, req.Context())
		return
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
// @Success		200	{array}	string	"List of supported dispensaries"	Enums(Curaleaf, Beyond-Hello)
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
// @Param			recreational	query	bool			true	"Recreational or medical"
// @Success		200				{array}	models.Location	"List of locations"
// @Router			/dispensaries/{dispensaryId}/locations [get].
func (a *API) handleLocationListing(r http.ResponseWriter, req *http.Request) {
	params, _ := params(req, "")
	locations, err := a.w.Locations(params)
	a.writeJson(r, req, a.t.TranslateLocations(locations), err)
}

// @Summary		Get product details
// @Description	Returns details of a specific product
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path		string			true	"Dispensary ID"
// @Param			locationId		path		string			true	"Location ID"
// @Param			recreational	query		bool			true	"Recreational or medical"
// @Param			productId		path		string			true	"Product ID"
// @Success		200				{object}	models.Product	"Product details"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/products/{productId} [get].
func (a *API) handleProduct(r http.ResponseWriter, req *http.Request) {
	params, productId := params(req, "productId")
	product, err := a.w.Product(params, productId)
	a.writeJson(r, req, a.t.TranslateProduct(product), err)
}

// @Summary		List products for a location
// @Description	Returns a list of products for a specific location
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			recreational	query	bool			true	"Recreational or medical"
// @Param			category		query	string			true	"Category"
// @Param			sub				query	string			false	"Sub Category"
// @Param			min_price_per_g	query	number			false	"Minimum price per gram"
// @Param			max_price_per_g	query	number			false	"Maximum price per gram"
// @Param			min_price		query	number			false	"Minimum price"
// @Param			max_price		query	number			false	"Maximum price"
// @Param			brands			query	string			false	"Brands to include"
// @Param			not_brands		query	string			false	"Brands to exclude"
// @Param			variants		query	string			false	"Variants to include"
// @Param			excludes		query	string			false	"Terms to exclude"
// @Param			includes		query	string			false	"Terms to include"
// @Param			price_sort		query	string			false	"Sort products by price"			Enums(asc, desc)
// @Param			thc_sort		query	string			false	"Sort products by THC"				Enums(asc, desc)
// @Param			terp_sort		query	string			false	"Sort products by Total Terpenes"	Enums(asc, desc)
// @Param			terp1			query	string			false	"Most important terpene"
// @Param			terp2			query	string			false	"Second most important terpene"
// @Param			terp3			query	string			false	"Third most important terpene"
// @Success		200				{array}	models.Product	"List of products"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/products [get].
func (a *API) handleProductListing(r http.ResponseWriter, req *http.Request) {
	params, _ := params(req, "")

	var products []*models.Product
	if category := req.URL.Query().Get("category"); category != "" {
		var err error
		products, err = a.w.ProductsInCategory(params, category)
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	} else {
		var err error
		products, err = a.w.Products(params)
		if err != nil {
			a.handleError(r, req, err)
			return
		}
	}

	fp := transformation.RequestToFilterParams(req)
	products, fErr := a.w.Filter(params, fp, products)
	if fErr != nil {
		a.handleError(r, req, fErr)
		return
	}

	sp := transformation.RequestToSortParams(req)
	sErr := a.w.Sort(params, sp, products)
	if sErr != nil {
		a.handleError(r, req, sErr)
		return
	}

	a.writeJson(r, req, a.t.TranslateProducts(products), nil)
}

// @Summary		List offers for a location
// @Description	Returns a list of offers for a specific location
// @Tags			offers
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			recreational	query	bool			true	"Recreational or medical"
// @Success		200				{array}	models.Offer	"List of offers"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/offers [get].
func (a *API) handleOfferListing(r http.ResponseWriter, req *http.Request) {
	params, _ := params(req, "")
	offers, err := a.w.Offers(params)
	a.writeJson(r, req, a.t.TranslateOffers(offers), err)
}

// @Summary		List categories for a location
// @Description	Returns a list of categories for a specific location
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string	true	"Dispensary ID"
// @Param			locationId		path	string	true	"Location ID"
// @Param			recreational	query	bool	true	"Recreational or medical"
// @Success		200				{array}	string	"List of categories"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/categories [get].
func (a *API) handleCategoryListing(r http.ResponseWriter, req *http.Request) {
	params, _ := params(req, "")
	categories, err := a.w.Categories(params)
	a.writeJson(r, req, a.t.TranslateCategories(categories), err)
}

// @Summary		List terpenes for a location
// @Description	Returns a list of terpenes for a specific location
// @Tags			terpenes
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string			true	"Dispensary ID"
// @Param			locationId		path	string			true	"Location ID"
// @Param			recreational	query	bool			true	"Recreational or medical"
// @Success		200				{array}	models.Terpene	"List of terpenes"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/terpenes [get].
func (a *API) handleTerpeneListing(r http.ResponseWriter, req *http.Request) {
	params, _ := params(req, "")
	terpenes, err := a.w.Terpenes(params)
	a.writeJson(r, req, a.t.TranslateTerpenes(terpenes), err)
}

// @Summary		List cannabinoids for a location
// @Description	Returns a list of cannabinoids for a specific location
// @Tags			cannabinoids
// @Accept			json
// @Produce		json
// @Param			dispensaryId	path	string				true	"Dispensary ID"
// @Param			locationId		path	string				true	"Location ID"
// @Param			recreational	query	bool				true	"Recreational or medical"
// @Success		200				{array}	models.Cannabinoid	"List of cannabinoids"
// @Router			/dispensaries/{dispensaryId}/locations/{locationId}/cannabinoids [get].
func (a *API) handleCannabinoidListing(r http.ResponseWriter, req *http.Request) {
	params, _ := params(req, "")
	cannabinoids, err := a.w.Cannabinoids(params)
	a.writeJson(r, req, a.t.TranslateCannabinoids(cannabinoids), err)
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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originRegex := regexp.MustCompile(`^https?:\/\/(localhost:[0-9]+|.*dahlton.org)\/?$`)
		origin := r.Header.Get("Origin")
		if originRegex.MatchString(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
			w.Header().Set("Access-Control-Max-Age", "300")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func params(req *http.Request, resource string) (wp workflow.WorkflowParams, resourceId string) {
	dispensaryId := chi.URLParam(req, "dispensaryId")
	locationId := chi.URLParam(req, "locationId")
	recreational := req.URL.Query().Get("recreational")

	if resource != "" {
		resourceId = chi.URLParam(req, resource)
	}

	return workflow.NewWorkflowParams(dispensaryId, locationId, recreational != ""), resourceId
}
