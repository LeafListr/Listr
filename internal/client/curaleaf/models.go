package curaleaf

const (
	Endpoint  = "https://graph.curaleaf.com/api/curaql"
	Authority = "graph.curaleaf.com"
	GbgId     = "LMR124"
	MenuType  = "MEDICAL"
)

type Image struct {
	URL string `json:"url"`
}

type Brand struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Image       Image  `json:"image"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
}

type Category struct {
	DisplayName string `json:"displayName"`
	Key         string `json:"key"`
}

type Effect struct {
	DisplayName string `json:"displayName"`
}

type CannabinoidObj struct {
	Cannabinoid Cannabinoid `json:"cannabinoid"`
	Unit        string      `json:"unit"`
	Value       float64     `json:"value"`
}

type Cannabinoid struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type TerpeneObj struct {
	Terpene    Terpene `json:"terpene"`
	UnitSymbol string  `json:"unitSymbol"`
	Value      float64 `json:"value"`
}

type Terpene struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type THC struct {
	Formatted string    `json:"formatted"`
	Range     []float64 `json:"range"`
}

type LabResult struct {
	Cannabinoids []CannabinoidObj `json:"cannabinoids"`
	Terpenes     []TerpeneObj     `json:"terpenes"`
	THC          THC              `json:"thc"`
}

type Offer struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id"`
	Title       string `json:"title"`
}

type Strain struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
}

type Subcategory struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
}

type Variant struct {
	FlowerEquivalent struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"flowerEquivalent,omitempty"`
	Id           string  `json:"id"`
	IsSpecial    bool    `json:"isSpecial"`
	Option       string  `json:"option"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	SpecialPrice float64 `json:"specialPrice,omitempty"`
}

type Product struct {
	Brand           Brand       `json:"brand"`
	Category        Category    `json:"category"`
	DescriptionHtml string      `json:"descriptionHtml,omitempty"`
	Effects         []Effect    `json:"effects,omitempty"`
	ID              string      `json:"id"`
	Images          []Image     `json:"images"`
	LabResults      LabResult   `json:"labResults"`
	Name            string      `json:"name"`
	Offers          []Offer     `json:"offers,omitempty"`
	Strain          Strain      `json:"strain,omitempty"`
	Subcategory     Subcategory `json:"subcategory,omitempty"`
	Variants        []Variant   `json:"variants"`
	CardDescription string      `json:"cardDescription,omitempty"`
}

type AllOffersResponse struct {
	Data struct {
		DispensaryMenu struct {
			Offers []Offer `json:"offers"`
		} `json:"dispensaryMenu"`
	} `json:"data"`
}

func NewAllOffersResponse(offers []Offer) AllOffersResponse {
	return AllOffersResponse{
		Data: struct {
			DispensaryMenu struct {
				Offers []Offer `json:"offers"`
			} `json:"dispensaryMenu"`
		}{
			DispensaryMenu: struct {
				Offers []Offer `json:"offers"`
			}{
				Offers: offers,
			},
		},
	}
}

type AllCategoriesResponse struct {
	Data struct {
		DispensaryMenu struct {
			AllFilters struct {
				Categories []Category `json:"categories"`
			} `json:"allFilters"`
		} `json:"dispensaryMenu"`
	} `json:"data"`
}

func NewAllCategoriesResponse(categories []Category) AllCategoriesResponse {
	return AllCategoriesResponse{
		Data: struct {
			DispensaryMenu struct {
				AllFilters struct {
					Categories []Category `json:"categories"`
				} `json:"allFilters"`
			} `json:"dispensaryMenu"`
		}{
			DispensaryMenu: struct {
				AllFilters struct {
					Categories []Category `json:"categories"`
				} `json:"allFilters"`
			}{
				AllFilters: struct {
					Categories []Category `json:"categories"`
				}{
					Categories: categories,
				},
			},
		},
	}
}

type AllProductsResponse struct {
	Data struct {
		DispensaryMenu struct {
			Offers   []Offer   `json:"offers"`
			Products []Product `json:"products"`
		} `json:"dispensaryMenu"`
	} `json:"data"`
}
