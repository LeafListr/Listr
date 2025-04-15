package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Linkinlog/LeafListr/internal/api/translation"
	"github.com/Linkinlog/LeafListr/internal/transformation"
	"github.com/Linkinlog/LeafListr/internal/views/components"
	"github.com/Linkinlog/LeafListr/internal/workflow"
)

type HtmlHandler struct {
	supportedDispensaries []string
	w                     *workflow.Workflow
}

func NewHtmlHandler(supportedDispensaries []string, w *workflow.Workflow) *HtmlHandler {
	return &HtmlHandler{
		supportedDispensaries: supportedDispensaries,
		w:                     w,
	}
}

func (h *HtmlHandler) LandingPage(r http.ResponseWriter, req *http.Request) {
	err := components.Index("LeafListr", h.supportedDispensaries).Render(req.Context(), r)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
	}
}

func (h *HtmlHandler) Locations(r http.ResponseWriter, req *http.Request) {
	selectedDispenary := req.URL.Query().Get("dispensary")
	params := params(req)
	repo, err := h.w.RepoFromFactory(params)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	locs, locErr := h.w.Locations(params, repo)
	if locErr != nil {
		http.Error(r, locErr.Error(), http.StatusInternalServerError)
		return
	}
	transLocs := translation.NewAPITranslator().TranslateLocations(locs)
	rErr := components.Locations(selectedDispenary, transLocs).Render(req.Context(), r)
	if rErr != nil {
		http.Error(r, rErr.Error(), http.StatusInternalServerError)
	}
}

func (h *HtmlHandler) Navigation(r http.ResponseWriter, req *http.Request) {
	err := components.NavigationForm(h.supportedDispensaries).Render(req.Context(), r)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
	}
}

func (h *HtmlHandler) Categories(r http.ResponseWriter, req *http.Request) {
	params := params(req)
	repo, err := h.w.RepoFromFactory(params)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, cErr := h.w.Categories(params, repo)
	if cErr != nil {
		http.Error(r, cErr.Error(), http.StatusInternalServerError)
		return
	}
	rErr := components.Categories(categories).Render(req.Context(), r)
	if rErr != nil {
		http.Error(r, rErr.Error(), http.StatusInternalServerError)
	}
}

func (h *HtmlHandler) OffersAndSearch(r http.ResponseWriter, req *http.Request) {
	params := params(req)
	repo, rErr := h.w.RepoFromFactory(params)
	if rErr != nil {
		http.Error(r, rErr.Error(), http.StatusInternalServerError)
		return
	}

	terps, terpErr := h.w.Terpenes(params, repo)
	if terpErr != nil {
		h.w.LogError(terpErr, req.Context())
	}
	transTerps := translation.NewAPITranslator().TranslateTerpenes(terps)

	offers, offerErr := h.w.Offers(params, repo)
	if offerErr != nil {
		http.Error(r, offerErr.Error(), http.StatusInternalServerError)
		return
	}
	transOffers := translation.NewAPITranslator().TranslateOffers(offers)

	categories, cErr := h.w.Categories(params, repo)
	if cErr != nil {
		http.Error(r, cErr.Error(), http.StatusInternalServerError)
		return
	}

	err := components.OffersAndSearch(categories, transOffers, transTerps).Render(req.Context(), r)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HtmlHandler) Sorters(r http.ResponseWriter, req *http.Request) {
	params := params(req)
	repo, rErr := h.w.RepoFromFactory(params)
	if rErr != nil {
		http.Error(r, rErr.Error(), http.StatusInternalServerError)
		return
	}

	terps, terpErr := h.w.Terpenes(params, repo)
	if terpErr != nil {
		h.w.LogError(terpErr, req.Context())
	}
	transTerps := translation.NewAPITranslator().TranslateTerpenes(terps)

	err := components.Sorters(transTerps).Render(req.Context(), r)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HtmlHandler) Filters(r http.ResponseWriter, req *http.Request) {
	err := components.Filters().Render(req.Context(), r)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
	}
}

func (h *HtmlHandler) Offers(r http.ResponseWriter, req *http.Request) {
	params := params(req)
	repo, rErr := h.w.RepoFromFactory(params)
	if rErr != nil {
		http.Error(r, rErr.Error(), http.StatusInternalServerError)
		return
	}

	offers, offerErr := h.w.Offers(params, repo)
	if offerErr != nil {
		http.Error(r, offerErr.Error(), http.StatusInternalServerError)
		return
	}
	transOffers := translation.NewAPITranslator().TranslateOffers(offers)

	err := components.Offers(transOffers).Render(req.Context(), r)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
	}
}

func (h *HtmlHandler) Products(r http.ResponseWriter, req *http.Request) {
	selectedCategory := req.URL.Query().Get("category")
	params := params(req)
	repo, err := h.w.RepoFromFactory(params)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	products, pErr := h.w.ProductsInCategory(params, selectedCategory, repo)
	if pErr != nil {
		http.Error(r, pErr.Error(), http.StatusInternalServerError)
		return
	}
	fp := transformation.RequestToFilterParams(req)
	fProducts, fErr := h.w.Filter(fp, products)
	if fErr != nil {
		h.w.LogError(fErr, req.Context())
		return
	}

	sp := transformation.RequestToSortParams(req)
	sErr := h.w.Sort(sp, fProducts)
	if sErr != nil {
		h.w.LogError(sErr, req.Context())
		return
	}
	transProds := translation.NewAPITranslator().TranslateProducts(fProducts)

	if req.URL.Query().Get("json") == "true" {
		r.Header().Set("Content-Type", "application/json")
		r.WriteHeader(http.StatusOK)
		json.NewEncoder(r).Encode(transProds)
		return
	}

	rErr := components.Products(transProds).Render(req.Context(), r)
	if rErr != nil {
		http.Error(r, rErr.Error(), http.StatusInternalServerError)
	}
}

func params(r *http.Request) (wp workflow.WorkflowParams) {
	selectedDispensary := r.URL.Query().Get("dispensary")
	selectedLocation := r.URL.Query().Get("location")
	recreationalString := r.URL.Query().Get("recreational")
	recreational := false
	if recreationalString == "on" {
		recreational = true
	}

	return workflow.NewWorkflowParams(selectedDispensary, selectedLocation, recreational)
}
