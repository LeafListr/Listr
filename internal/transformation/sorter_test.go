package transformation_test

import (
	"reflect"
	"testing"

	"github.com/Linkinlog/LeafListr/internal/transformation"

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
			transformation.NewSorterer(nil).PriceAsc(tc.products)
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
			transformation.NewSorterer(nil).PriceDesc(tc.products)
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

func TestTop3Terps(t *testing.T) {
	tests := map[string]struct {
		products []*models.Product
		terps    [3]string
		want     []*models.Product
	}{
		"empty": {
			products: []*models.Product{},
			want:     []*models.Product{},
		},
		"multiple": {
			products: []*models.Product{
				{
					Id: "1",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 0.1},
						{Name: "Terpene 2", Value: 0.2},
						{Name: "Terpene 3", Value: 0.3},
						{Name: "Terpene 4", Value: 0.4},
					},
				},
				{
					Id: "2",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 2.5},
						{Name: "Terpene 2", Value: 0.6},
						{Name: "Terpene 3", Value: 3.7},
						{Name: "Terpene 4", Value: 0.7},
					},
				},
				{
					Id: "3",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 1.5},
						{Name: "Terpene 2", Value: 2.6},
						{Name: "Terpene 3", Value: 2.7},
						{Name: "Terpene 4", Value: 0.7},
					},
				},
			},
			terps: [3]string{"Terpene 2", "Terpene 1", "Terpene 3"},
			want: []*models.Product{
				{
					Id: "3",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 1.5},
						{Name: "Terpene 2", Value: 2.6},
						{Name: "Terpene 3", Value: 2.7},
						{Name: "Terpene 4", Value: 0.7},
					},
				},
				{
					Id: "2",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 2.5},
						{Name: "Terpene 2", Value: 0.6},
						{Name: "Terpene 3", Value: 3.7},
						{Name: "Terpene 4", Value: 0.7},
					},
				},
				{
					Id: "1",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 0.1},
						{Name: "Terpene 2", Value: 0.2},
						{Name: "Terpene 3", Value: 0.3},
						{Name: "Terpene 4", Value: 0.4},
					},
				},
			},
		},
		"terpene2 outweighs terpene1": {
			products: []*models.Product{
				{
					Id: "A",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 1.0},
						{Name: "Terpene 2", Value: 2.0},
					},
				},
				{
					Id: "B",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 2.0},
						{Name: "Terpene 2", Value: 3.5},
					},
				},
			},
			terps: [3]string{"Terpene 1", "Terpene 2", "Terpene 3"},
			want: []*models.Product{
				{Id: "B"},
				{Id: "A"},
			},
		},
		"terpene3 outweighs terpene1 and terpene2": {
			products: []*models.Product{
				{
					Id: "C",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 1.0},
						{Name: "Terpene 2", Value: 1.0},
						{Name: "Terpene 3", Value: 4.0},
					},
				},
				{
					Id: "D",
					T: []*models.Terpene{
						{Name: "Terpene 1", Value: 0.5},
						{Name: "Terpene 2", Value: 0.5},
						{Name: "Terpene 3", Value: 2.0},
					},
				},
			},
			terps: [3]string{"Terpene 1", "Terpene 2", "Terpene 3"},
			want: []*models.Product{
				{Id: "C"},
				{Id: "D"},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			transformation.NewSorterer(nil).Top3Terps(tc.products, tc.terps)
			for i, p := range tc.products {
				if p.Id != tc.want[i].Id {
					t.Fatalf("got: %+v, want: %+v", p.Id, tc.want[i].Id)
				}
			}
		})
	}
}

func TestTHCAsc(t *testing.T) {
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
					Id: "1",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 10,
						},
					},
				},
				{
					Id: "2",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 25,
						},
					},
				},
				{
					Id: "3",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 20,
						},
					},
				},
			},
			want: []*models.Product{
				{
					Id: "1",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 10,
						},
					},
				},
				{
					Id: "3",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 20,
						},
					},
				},
				{
					Id: "2",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 25,
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			transformation.NewSorterer(nil).THCAsc(tc.products)
			if !reflect.DeepEqual(tc.products, tc.want) {
				t.Fatalf("got: %v, want: %v", tc.products, tc.want)
			}
		})
	}
}

func TestTHCDesc(t *testing.T) {
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
					Id: "1",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 10,
						},
					},
				},
				{
					Id: "2",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 25,
						},
					},
				},
				{
					Id: "3",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 20,
						},
					},
				},
			},
			want: []*models.Product{
				{
					Id: "2",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 25,
						},
					},
				},
				{
					Id: "3",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 20,
						},
					},
				},
				{
					Id: "1",
					C: []*models.Cannabinoid{
						{
							Name:  "THC",
							Value: 10,
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			transformation.NewSorterer(nil).THCDesc(tc.products)
			if !reflect.DeepEqual(tc.products, tc.want) {
				t.Fatalf("got: %v, want: %v", tc.products, tc.want)
			}
		})
	}
}
