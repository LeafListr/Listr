package models

import "strings"

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
	Id     string
	Brand  string
	Name   string
	SubCtg string
	Images []string
	Weight string
	Ctg    Category
	Price  *Price
	C      []*Cannabinoid
	T      []*Terpene
}

func (p *Product) THC() float64 {
	for _, c := range p.C {
		if strings.Contains(strings.ToLower(c.Name), "thc") {
			return c.Value
		}
	}

	return 0
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
	Id            string
	LocationTypes []string
	Name          string
	Address       string
	City          string
	State         string
	ZipCode       string
}
