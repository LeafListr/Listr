package transformation

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func RequestToFilterParams(req *http.Request) *FilterParams {
	return &FilterParams{
		SubCategoryName:    subCategoryFilter(req),
		MinPricePerG:       minPricePerGFilter(req),
		MaxPricePerG:       maxPricePerGFilter(req),
		MinPrice:           minPriceFilter(req),
		MaxPrice:           maxPriceFilter(req),
		IncludedBrandNames: brandFilters(req),
		ExcludedBrandNames: notBrandFilters(req),
		Variants:           variantFilters(req),
		IncludedTerms:      includedTerms(req),
		ExcludedTerms:      excludedTerms(req),
	}
}

func RequestToSortParams(request *http.Request) *SortParams {
	return &SortParams{
		Top3Terps:     top3Terps(request),
		PriceAsc:      sortPriceAsc(request),
		PriceDesc:     sortPriceDesc(request),
		THCAsc:        sortTHCAsc(request),
		THCDesc:       sortTHCDesc(request),
		TerpAsc:       sortTerpAsc(request),
		TerpDesc:      sortTerpDesc(request),
		GramPriceAsc:  sortGramPriceAsc(request),
		GramPriceDesc: sortGramPriceDesc(request),
	}
}

func subCategoryFilter(req *http.Request) string {
	return req.URL.Query().Get("sub")
}

func maxPricePerGFilter(req *http.Request) float64 {
	var maxP float64
	var err error
	if maxStr := req.URL.Query().Get("max_price_per_g"); maxStr != "" {
		maxP, err = strconv.ParseFloat(maxStr, 64)
		if err != nil {
			slog.Debug("Error parsing max price", err)
			return 0
		}
	}
	return maxP
}

func minPricePerGFilter(req *http.Request) float64 {
	var minP float64
	var err error
	if minStr := req.URL.Query().Get("min_price_per_g"); minStr != "" {
		minP, err = strconv.ParseFloat(minStr, 64)
		if err != nil {
			slog.Debug("Error parsing min price", err)
			return 0
		}
	}
	return minP
}

func maxPriceFilter(req *http.Request) float64 {
	var maxP float64
	var err error
	if maxStr := req.URL.Query().Get("max_price"); maxStr != "" {
		maxP, err = strconv.ParseFloat(maxStr, 64)
		if err != nil {
			slog.Debug("Error parsing max price", err)
			return 0
		}
	}
	return maxP
}

func minPriceFilter(req *http.Request) float64 {
	var minP float64
	var err error
	if minStr := req.URL.Query().Get("min_price"); minStr != "" {
		minP, err = strconv.ParseFloat(minStr, 64)
		if err != nil {
			slog.Debug("Error parsing min price", err)
			return 0
		}
	}
	return minP
}

func brandFilters(req *http.Request) []string {
	return strings.Split(req.URL.Query().Get("brands"), ",")
}

func notBrandFilters(req *http.Request) []string {
	return strings.Split(req.URL.Query().Get("not_brands"), ",")
}

func variantFilters(req *http.Request) []string {
	return strings.Split(req.URL.Query().Get("variants"), ",")
}

func excludedTerms(req *http.Request) []string {
	return strings.Split(req.URL.Query().Get("exclude"), ",")
}

func includedTerms(req *http.Request) []string {
	return strings.Split(req.URL.Query().Get("include"), ",")
}

func sortPriceAsc(req *http.Request) bool {
	return req.URL.Query().Get("price_sort") == "asc"
}

func sortPriceDesc(req *http.Request) bool {
	return req.URL.Query().Get("price_sort") == "desc"
}

func sortTHCAsc(req *http.Request) bool {
	return req.URL.Query().Get("thc_sort") == "asc"
}

func sortTHCDesc(req *http.Request) bool {
	return req.URL.Query().Get("thc_sort") == "desc"
}

func sortTerpAsc(req *http.Request) bool {
	return req.URL.Query().Get("terp_sort") == "asc"
}

func sortTerpDesc(req *http.Request) bool {
	return req.URL.Query().Get("terp_sort") == "desc"
}

func sortGramPriceAsc(req *http.Request) bool {
	return req.URL.Query().Get("gram_price_sort") == "asc"
}

func sortGramPriceDesc(req *http.Request) bool {
	return req.URL.Query().Get("gram_price_sort") == "desc"
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
