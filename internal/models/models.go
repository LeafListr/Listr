package models

type Price struct {
	Total           float64
	DiscountedTotal float64
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
	Name  string
	Price *Price
}

type Product struct {
	Id     string
	Name   string
	Ctg    Category
	Images []string
	V      []*Variant
	C      []*Cannabinoid
	T      []*Terpene
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
	Name    string
	Address string
	City    string
	State   string
	ZipCode string
}
