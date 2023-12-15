package curaleaf_test

import (
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestClientTranslator_TranslateClientLocation(t *testing.T) {
	tests := map[string]struct {
		input    curaleaf.Location
		expected *models.Location
	}{
		"Translate Valid Location": {
			input: curaleaf.Location{
				UniqueId: "TestLocation",
				Name:     "TestName",
				Location: curaleaf.LocationDetails{
					Address: "123 Test St",
					City:    "TestCity",
					State:   "TestState",
					ZipCode: "12345",
				},
			},
			expected: &models.Location{
				Name:    "TestName",
				Id:      "TestLocation",
				Address: "123 Test St",
				City:    "TestCity",
				State:   "TestState",
				ZipCode: "12345",
			},
		},
	}

	translator := &curaleaf.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientLocation(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientTranslator_TranslateClientLocations(t *testing.T) {
	tests := map[string]struct {
		input    []curaleaf.Location
		expected []*models.Location
	}{
		"Translate Multiple Locations": {
			input: []curaleaf.Location{
				{
					Name:     "Name1",
					UniqueId: "Location1",
					Location: curaleaf.LocationDetails{
						Address: "123 Location St",
						City:    "LocationCity",
						State:   "LocationState",
						ZipCode: "54321",
					},
				},
				{
					Name:     "Name2",
					UniqueId: "Location2",
					Location: curaleaf.LocationDetails{
						Address: "456 Location Ave",
						City:    "AnotherCity",
						State:   "AnotherState",
						ZipCode: "98765",
					},
				},
			},
			expected: []*models.Location{
				{
					Name:    "Name1",
					Id:      "Location1",
					Address: "123 Location St",
					City:    "LocationCity",
					State:   "LocationState",
					ZipCode: "54321",
				},
				{
					Name:    "Name2",
					Id:      "Location2",
					Address: "456 Location Ave",
					City:    "AnotherCity",
					State:   "AnotherState",
					ZipCode: "98765",
				},
			},
		},
	}

	translator := &curaleaf.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientLocations(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientTranslator_TranslateClientProduct(t *testing.T) {
	tests := map[string]struct {
		input    curaleaf.Product
		expected *models.Product
	}{
		"Translate Valid Product": {
			input: curaleaf.Product{
				ID:   "Product1",
				Name: "Test Product",
			},
			expected: &models.Product{
				Id:   "Product1",
				Name: "Test Product",
			},
		},
	}

	translator := &curaleaf.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientProduct(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientTranslator_TranslateClientProducts(t *testing.T) {
	tests := map[string]struct {
		input    []curaleaf.Product
		expected []*models.Product
	}{
		"Translate Multiple Products": {
			input: []curaleaf.Product{
				{
					ID:   "Product1",
					Name: "Test Product 1",
				},
				{
					ID:   "Product2",
					Name: "Test Product 2",
				},
			},
			expected: []*models.Product{
				{
					Id:   "Product1",
					Name: "Test Product 1",
				},
				{
					Id:   "Product2",
					Name: "Test Product 2",
				},
			},
		},
	}

	translator := &curaleaf.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientProducts(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
