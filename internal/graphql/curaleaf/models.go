package curaleaf

const (
	Endpoint  = "https://graph.curaleaf.com/api/curaql"
	Authority = "graph.curaleaf.com"
	GbgId     = "LMR124"
	MenuType  = "MEDICAL"
)

type Image struct {
	URL      string `json:"url"`
	TypeName string `json:"__typename,omitempty"`
}

type Brand struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Image       Image  `json:"image"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	TypeName    string `json:"__typename,omitempty"`
}

type Category struct {
	DisplayName string `json:"displayName"`
	Key         string `json:"key"`
	TypeName    string `json:"__typename,omitempty"`
}

type Effect struct {
	DisplayName string `json:"displayName"`
}

type Cannabinoid struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Terpene struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type THC struct {
	Formatted string `json:"formatted"`
	Range     string `json:"range"`
	TypeName  string `json:"__typename,omitempty"`
}

type LabResult struct {
	Cannabinoids []struct {
		Cannabinoid Cannabinoid `json:"cannabinoid"`
		Unit        string      `json:"unit"`
		Value       float64     `json:"value"`
	} `json:"cannabinoids"`
	Terpenes []struct {
		Terpene    Terpene `json:"terpene"`
		UnitSymbol string  `json:"unitSymbol"`
		Value      float64 `json:"value"`
	} `json:"terpenes"`
	THC      THC    `json:"thc"`
	TypeName string `json:"__typename,omitempty"`
}

type Offer struct {
	Description string `json:"description,omitempty"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	TypeName    string `json:"__typename,omitempty"`
}

type Strain struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
	TypeName    string `json:"__typename,omitempty"`
}

type Subcategory struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
	TypeName    string `json:"__typename,omitempty"`
}

type Variant struct {
	FlowerEquivalent struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"flowerEquivalent,omitempty"`
	ID           string  `json:"id"`
	IsSpecial    bool    `json:"isSpecial"`
	Option       string  `json:"option"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	SpecialPrice float64 `json:"specialPrice,omitempty"`
	TypeName     string  `json:"__typename,omitempty"`
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
	TypeName        string      `json:"__typename,omitempty"`
	CardDescription string      `json:"cardDescription,omitempty"`
}

type AllProductResponse struct {
	Data struct {
		Product struct {
			Product Product `json:"product"`
		} `json:"product"`
	} `json:"data"`
}

type AllProductForCategoryResponse struct {
	Data struct {
		DispensaryMenu struct {
			Offers   []Offer   `json:"offers"`
			Products []Product `json:"products"`
			TypeName string    `json:"__typename"`
		} `json:"dispensaryMenu"`
	} `json:"data"`
}
