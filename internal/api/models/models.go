package models

type Price struct {
	Total           float64 `json:"total"`
	DiscountedTotal float64 `json:"discountedTotal"`
	IsDiscounted    bool    `json:"isDiscounted"`
	PerGram         float64 `json:"perGram"`
}

type Cannabinoid struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type Terpene struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type Product struct {
	Id        string         `json:"id"`
	Brand     string         `json:"brand"`
	Name      string         `json:"name"`
	Ctg       string         `json:"category"`
	SubCtg    string         `json:"subcategory"`
	Weight    string         `json:"variant,omitempty"`
	Images    []string       `json:"images"`
	P         *Price         `json:"price,omitempty"`
	C         []*Cannabinoid `json:"cannabinoids,omitempty"`
	T         []*Terpene     `json:"terpenes,omitempty"`
	TT        float64        `json:"terpeneTotal,omitempty"`
	Permalink string         `json:"permalink"`
}

func (p *Product) Image() string {
	if len(p.Images) == 0 {
		return ""
	}
	return p.Images[0]
}

type Offer struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type Dispensary struct {
	Name      string      `json:"name"`
	Locations []*Location `json:"locations,omitempty"`
}

type Location struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}
