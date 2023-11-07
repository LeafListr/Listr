package models

type Price struct {
	Variant         string
	Total           float64
	DiscountedTotal float64 `json:"DiscountedTotal,omitempty"`
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

type Product struct {
	Id   string         `json:"Id,omitempty"`
	Name string         `json:"Name,omitempty"`
	Ctg  Category       `json:"Category,omitempty"`
	P    []*Price       `json:"Price,omitempty"`
	C    []*Cannabinoid `json:"Cannabinoids,omitempty"`
	T    []*Terpene     `json:"Terpenes,omitempty"`
}

type Offer struct {
	Id          string
	Description string
}

type Brand struct {
	UniqueId   string   `json:"Id"`
	Name       string   `json:"name"`
	OrderTypes []string `json:"orderTypes"`
	MenuTypes  []string `json:"menuTypes"`
	IsOpened   bool     `json:"isOpened"`
	Location   Location `json:"location"`
}

type Location struct {
	Coordinates Coordinates `json:"coordinates"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	ZipCode     string      `json:"zipCode"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
