package models

type Price struct {
	Total           float64
	DiscountedTotal float64 `json:"discountedTotal,omitempty"`
}

type Cannabinoid struct {
	Name        string
	Description string
	Value       float64
}

type Terpene struct {
	Name        string
	Description string
	Value       float64
}

type Category string

type Variant struct {
	Name  string `json:"Name,omitempty"`
	Price *Price `json:"Price,omitempty"`
}

type Product struct {
	Id   string         `json:"id,omitempty"`
	Name string         `json:"name,omitempty"`
	Ctg  Category       `json:"category,omitempty"`
	V    []*Variant     `json:"variant,omitempty"`
	C    []*Cannabinoid `json:"cannabinoids,omitempty"`
	T    []*Terpene     `json:"terpenes,omitempty"`
}

type Offer struct {
	Id          string
	Description string
}

type Dispensary struct {
	Name      string      `json:"name"`
	Locations []*Location `json:"locations"`
}

type Location struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type Menu struct {
	Products     []*Product     `json:"products,omitempty"`
	Offers       []*Offer       `json:"offers,omitempty"`
	Categories   []*Category    `json:"categories,omitempty"`
	Terpenes     []*Terpene     `json:"terpenes,omitempty"`
	Cannabinoids []*Cannabinoid `json:"cannabinoids,omitempty"`
}
