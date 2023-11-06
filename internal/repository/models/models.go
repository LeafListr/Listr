package models

type Price struct {
	Variant         string
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

type Product struct {
	Id  string
	P   []*Price
	C   []*Cannabinoid
	T   []*Terpene
	Ctg Category
}

type Offer struct {
	Id          string
	Description string
}
