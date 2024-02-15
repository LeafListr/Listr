package models

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Price struct {
	Total           float64
	DiscountedTotal float64
	IsDiscounted    bool
	PerGram         float64
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

type Product struct {
	Id     string
	Brand  string
	Name   string
	SubCtg string
	Images []string
	Weight string
	Ctg    string
	P      *Price
	C      []*Cannabinoid
	T      []*Terpene
}

func (p *Product) WeightInGrams() float64 {
	digitOrFractionPattern := regexp.MustCompile(`^(\d*\.?\d+|\d+\/\d+)\s*([mg|g|oz]+)$`)
	matches := digitOrFractionPattern.FindStringSubmatch(p.Weight)

	if matches == nil {
		return 0.0
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil && !isFraction(matches[1]) {
		return 0.0
	}

	unit := matches[2]
	switch unit {
	case "mg":
		return value / 1000
	case "g":
		return value
	case "oz":
		if isFraction(matches[1]) {
			gramsPerUnit, cErr := convertFractionToGrams(matches[1], unit)
			if cErr != nil {
				return 0.0
			}
			return gramsPerUnit
		} else {
			return value * 28
		}
	default:
		return 0.0
	}
}

func (p *Product) Price() float64 {
	if p.P.IsDiscounted {
		return p.P.DiscountedTotal
	}
	return p.P.Total
}

func (p *Product) PricePerGram() float64 {
	if p.WeightInGrams() == 0 {
		return 0
	}
	return p.Price() / p.WeightInGrams()
}

func isFraction(str string) bool {
	return strings.Contains(str, "/")
}

func convertFractionToGrams(frac string, unit string) (float64, error) {
	parts := strings.Split(frac, "/")
	numerator, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	denominator, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	switch unit {
	case "oz":
		return float64(numerator) / float64(denominator) * 28, nil
	default:
		return 0, fmt.Errorf("unsupported fractional unit: %s", unit)
	}
}

func (p *Product) THC() float64 {
	for _, c := range p.C {
		if strings.Contains(strings.ToLower(c.Name), "thc") {
			return c.Value
		}
	}

	return 0
}

func (p *Product) TotalTerps() float64 {
	sum := 0.0
	for _, t := range p.T {
		sum += t.Value
	}

	return sum
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
	Categories   []*string
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
