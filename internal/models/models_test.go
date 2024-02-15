package models_test

import (
	"math"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/models"
)

func TestProduct_Weight(t *testing.T) {
	tests := map[string]struct {
		weight string
		want   float64
	}{
		"bad data": {
			weight: "30bites",
			want:   0.0,
		},
		"30mg": {
			weight: "30mg",
			want:   0.03,
		},
		".3g": {
			weight: ".3g",
			want:   0.3,
		},
		"1g": {
			weight: "1g",
			want:   1.0,
		},
		"2g": {
			weight: "2g",
			want:   2.0,
		},
		"3.5g": {
			weight: "3.5g",
			want:   3.5,
		},
		"1/8oz": {
			weight: "1/8oz",
			want:   3.5,
		},
		"1/4oz": {
			weight: "1/4oz",
			want:   7.0,
		},
		"1/2oz": {
			weight: "1/2oz",
			want:   14.0,
		},
		"1oz": {
			weight: "1oz",
			want:   28.0,
		},
		"2oz": {
			weight: "2oz",
			want:   56.0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := &models.Product{Weight: tc.weight}
			got := p.WeightInGrams()
			if got != tc.want {
				t.Errorf("expected %f, got %f", tc.want, got)
			}
		})
	}
}

func TestProduct_PricePerGram(t *testing.T) {
	tests := map[string]struct {
		price  *models.Price
		weight string
		want   float64
	}{
		"bad data": {
			price:  &models.Price{Total: 0.0},
			weight: "30bites",
			want:   0.0,
		},
		".3g for $3": {
			price:  &models.Price{Total: 3.0},
			weight: ".3g",
			want:   10.0,
		},
		"1g for $20": {
			price:  &models.Price{Total: 20.0},
			weight: "1g",
			want:   20.0,
		},
		"1g for $20, discounted to $10": {
			price:  &models.Price{Total: 20.0, IsDiscounted: true, DiscountedTotal: 10.0},
			weight: "1g",
			want:   10.0,
		},
		"1g for $10": {
			price:  &models.Price{Total: 10.0},
			weight: "1g",
			want:   10.0,
		},
		"3.5g for $60": {
			price:  &models.Price{Total: 60.0},
			weight: "3.5g",
			want:   17.14,
		},
		"3.5g for $35": {
			price:  &models.Price{Total: 35.0},
			weight: "3.5g",
			want:   10.0,
		},
		"1/8oz for $35": {
			price:  &models.Price{Total: 35.0},
			weight: "1/8oz",
			want:   10.0,
		},
		"1/4oz for $70": {
			price:  &models.Price{Total: 70.0},
			weight: "1/4oz",
			want:   10.0,
		},
		"1/2oz for $140": {
			price:  &models.Price{Total: 140.0},
			weight: "1/2oz",
			want:   10.0,
		},
		"1oz for $280": {
			price:  &models.Price{Total: 280.0},
			weight: "1oz",
			want:   10.0,
		},
		"2oz for $560": {
			price:  &models.Price{Total: 560.0},
			weight: "2oz",
			want:   10.0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := &models.Product{P: tc.price, Weight: tc.weight}
			got := p.PricePerGram()
			if math.Round(got*100)/100 != tc.want {
				t.Errorf("expected %f, got %f", tc.want, got)
			}
		})
	}
}
