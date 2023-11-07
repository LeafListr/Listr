package workflow

import (
	"github.com/Linkinlog/LeafList/internal/factory/factoryfakes"
	"github.com/Linkinlog/LeafList/internal/models"
	"github.com/Linkinlog/LeafList/internal/repository"
	"reflect"
	"testing"
)

func TestWorkflow_AllProducts(t *testing.T) {
	type fields struct {
		r repository.Repository
	}
	type args struct {
		menuId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factoryfakes.FakeRepositoryFactory{}
			f.FindReturns(tt.fields.r, nil)
			w := &Workflow{
				f: &f,
			}
			got, err := w.AllProducts("", tt.args.menuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllProducts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_AllProductsForCategory(t *testing.T) {
	type fields struct {
		r repository.Repository
	}
	type args struct {
		menuId   string
		category models.Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factoryfakes.FakeRepositoryFactory{}
			f.FindReturns(tt.fields.r, nil)
			w := &Workflow{
				f: &f,
			}
			got, err := w.AllProductsForCategory("", tt.args.menuId, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllProductsForCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllProductsForCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_Categories(t *testing.T) {
	type fields struct {
		r repository.Repository
	}
	type args struct {
		menuId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Category
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factoryfakes.FakeRepositoryFactory{}
			f.FindReturns(tt.fields.r, nil)
			w := &Workflow{
				f: &f,
			}
			got, err := w.Categories("", tt.args.menuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Categories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Categories() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_Offers(t *testing.T) {
	type fields struct {
		r repository.Repository
	}
	type args struct {
		menuId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Offer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factoryfakes.FakeRepositoryFactory{}
			f.FindReturns(tt.fields.r, nil)
			w := &Workflow{
				f: &f,
			}
			got, err := w.Offers("", tt.args.menuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Offers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Offers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_Product(t *testing.T) {
	type fields struct {
		r repository.Repository
	}
	type args struct {
		menuId    string
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factoryfakes.FakeRepositoryFactory{}
			f.FindReturns(tt.fields.r, nil)
			w := &Workflow{
				f: &f,
			}
			got, err := w.Product("", tt.args.menuId, tt.args.productId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Product() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Product() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_Terpenes(t *testing.T) {
	type fields struct {
		r repository.Repository
	}
	type args struct {
		menuId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Terpene
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factoryfakes.FakeRepositoryFactory{}
			f.FindReturns(tt.fields.r, nil)
			w := &Workflow{
				f: &f,
			}
			got, err := w.Terpenes("", tt.args.menuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Terpenes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Terpenes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
