package models

type Price struct {
	Total           float64
	DiscountedTotal float64
	IsDiscounted    bool
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
	Id      string
	Brand   string
	Name    string
	SubCtg  string
	Images  []string
	Variant string
	Ctg     Category
	Price   *Price
	C       []*Cannabinoid
	T       []*Terpene
}

type Offer struct {
	Id          string
	Description string
}

type Dispensary struct {
	Name      string
	Locations []*Location
}

type Menu struct {
	Products     []*Product
	Offers       []*Offer
	Categories   []*Category
	Terpenes     []*Terpene
	Cannabinoids []*Cannabinoid
}

type Location struct {
	Id      string
	Name    string
	Address string
	City    string
	State   string
	ZipCode string
}
