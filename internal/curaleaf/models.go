package curaleaf

const (
	ProductCategoryVape = "VAPORIZERS"
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

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LocationDetails struct {
	Coordinates       Coordinates `json:"coordinates"`
	Address           string      `json:"address"`
	City              string      `json:"city"`
	Distance          float64     `json:"distance"`
	DistanceFormatted string      `json:"distanceFormatted"`
	State             string      `json:"state"`
	StateAbbreviation string      `json:"stateAbbreviation"`
	StateSlug         string      `json:"stateSlug"`
	ZipCode           string      `json:"zipCode"`
}

type Location struct {
	UniqueId         string          `json:"uniqueId"`
	Name             string          `json:"name"`
	Slug             string          `json:"slug"`
	OrderTypes       []string        `json:"orderTypes"`
	MenuTypes        []string        `json:"menuTypes"`
	IsOpened         bool            `json:"isOpened"`
	Location         LocationDetails `json:"location"`
	NextTime         string          `json:"nextTime"`
	ValidForDelivery bool            `json:"validForDelivery"`
}

type Response struct {
	DataObj
	ErrorObj
}

type ErrorObj struct {
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Code      string `json:"code"`
			Exception struct {
				Stacktrace []string `json:"stacktrace"`
			} `json:"exception"`
		} `json:"extensions"`
	} `json:"errors,omitempty"`
}

type DataObj struct {
	Data struct {
		Dispensaries []Location `json:"dispensaries,omitempty"`
		DispensaryMenuObj
		ProductObj
	} `json:"data,omitempty"`
}

type ProductObj struct {
	Product struct {
		Product Product `json:"product"`
	} `json:"product,omitempty"`
}

type DispensaryMenuObj struct {
	DispensaryMenu struct {
		Offers   []Offer   `json:"offers"`
		Products []Product `json:"products"`
		AllFiltersObj
	} `json:"dispensaryMenu,omitempty"`
}

type AllFiltersObj struct {
	AllFilters struct {
		Categories []Category `json:"categories"`
	} `json:"allFilters,omitempty"`
}

func NewResponse(products []Product, offers []Offer, categories []Category, locations []Location) *Response {
	var product Product
	if len(products) >= 1 {
		product = products[0]
	}
	return &Response{
		DataObj: DataObj{
			Data: struct {
				Dispensaries []Location `json:"dispensaries,omitempty"`
				DispensaryMenuObj
				ProductObj
			}{
				Dispensaries: locations,
				DispensaryMenuObj: DispensaryMenuObj{
					DispensaryMenu: struct {
						Offers   []Offer   `json:"offers"`
						Products []Product `json:"products"`
						AllFiltersObj
					}{
						Offers:   offers,
						Products: products,
						AllFiltersObj: AllFiltersObj{
							AllFilters: struct {
								Categories []Category `json:"categories"`
							}{
								Categories: categories,
							},
						},
					},
				},
				ProductObj: ProductObj{
					Product: struct {
						Product Product `json:"product"`
					}{
						Product: product,
					},
				},
			},
		},
		ErrorObj: ErrorObj{
			Errors: []struct {
				Message   string `json:"message"`
				Locations []struct {
					Line   int `json:"line"`
					Column int `json:"column"`
				} `json:"locations"`
				Path       []string `json:"path"`
				Extensions struct {
					Code      string `json:"code"`
					Exception struct {
						Stacktrace []string `json:"stacktrace"`
					} `json:"exception"`
				} `json:"extensions"`
			}{},
		},
	}
}
