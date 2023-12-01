package transformation_test

import (
	"reflect"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/transformation"
	"github.com/Linkinlog/LeafListr/internal/models"
)

func TestPriceAsc(t *testing.T) {
	tests := map[string]struct {
		products []*models.Product
		want     []*models.Product
	}{
		"empty": {
			products: []*models.Product{},
			want:     []*models.Product{},
		},
		"multiple": {
			products: []*models.Product{
				{
					Price: &models.Price{Total: 10, DiscountedTotal: 5},
				},
				{
					Price: &models.Price{Total: 25, DiscountedTotal: 2},
				},
				{
					Price: &models.Price{Total: 20, DiscountedTotal: 10},
				},
			},
			want: []*models.Product{
				{
					Price: &models.Price{Total: 25, DiscountedTotal: 2},
				},
				{
					Price: &models.Price{Total: 10, DiscountedTotal: 5},
				},
				{
					Price: &models.Price{Total: 20, DiscountedTotal: 10},
				},
			},
		},
		"multiple, some without discounts": {
			products: []*models.Product{
				{
					Price: &models.Price{Total: 10, DiscountedTotal: 5},
				},
				{
					Price: &models.Price{Total: 25, DiscountedTotal: 0},
				},
				{
					Price: &models.Price{Total: 20, DiscountedTotal: 10},
				},
			},
			want: []*models.Product{
				{
					Price: &models.Price{Total: 10, DiscountedTotal: 5},
				},
				{
					Price: &models.Price{Total: 20, DiscountedTotal: 10},
				},
				{
					Price: &models.Price{Total: 25, DiscountedTotal: 0},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			transformation.NewSorterer().PriceAsc(tc.products)
			if !reflect.DeepEqual(tc.products, tc.want) {
				t.Fatalf("got: %v, want: %v", tc.products, tc.want)
			}
		})
	}
}

func TestPriceDesc(t *testing.T) {
	tests := map[string]struct {
		products []*models.Product
		want     []*models.Product
	}{
		"empty": {
			products: []*models.Product{},
			want:     []*models.Product{},
		},
		"multiple": {
			products: []*models.Product{
				{
					Price: &models.Price{Total: 10, DiscountedTotal: 5},
				},
				{
					Price: &models.Price{Total: 25, DiscountedTotal: 2},
				},
				{
					Price: &models.Price{Total: 20, DiscountedTotal: 10},
				},
			},
			want: []*models.Product{
				{
					Price: &models.Price{Total: 20, DiscountedTotal: 10},
				},
				{
					Price: &models.Price{Total: 10, DiscountedTotal: 5},
				},
				{
					Price: &models.Price{Total: 25, DiscountedTotal: 2},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			transformation.NewSorterer().PriceDesc(tc.products)
			if !reflect.DeepEqual(tc.products, tc.want) {
				for _, p := range tc.products {
					t.Logf("got: %+v", *p.Price)
				}
				for _, p := range tc.want {
					t.Logf("wanted: %+v", *p.Price)
				}
				t.Fatalf("got: %+v, want: %+v", tc.products, tc.want)
			}
		})
	}
}
