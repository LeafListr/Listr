package translation_test

import (
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/translation"
	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestClientTranslator_TranslateClientLocation(t *testing.T) {
	tests := map[string]struct {
		input    client.Location
		expected *models.Location
	}{
		"Translate Valid Location": {
			input: client.Location{
				UniqueId: "TestLocation",
				Location: client.LocationDetails{
					Address: "123 Test St",
					City:    "TestCity",
					State:   "TestState",
					ZipCode: "12345",
				},
			},
			expected: &models.Location{
				Name:    "TestLocation",
				Address: "123 Test St",
				City:    "TestCity",
				State:   "TestState",
				ZipCode: "12345",
			},
		},
	}

	translator := &translation.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientLocation(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientTranslator_TranslateClientLocations(t *testing.T) {
	tests := map[string]struct {
		input    []client.Location
		expected []*models.Location
	}{
		"Translate Multiple Locations": {
			input: []client.Location{
				{
					UniqueId: "Location1",
					Location: client.LocationDetails{
						Address: "123 Location St",
						City:    "LocationCity",
						State:   "LocationState",
						ZipCode: "54321",
					},
				},
				{
					UniqueId: "Location2",
					Location: client.LocationDetails{
						Address: "456 Location Ave",
						City:    "AnotherCity",
						State:   "AnotherState",
						ZipCode: "98765",
					},
				},
			},
			expected: []*models.Location{
				{
					Name:    "Location1",
					Address: "123 Location St",
					City:    "LocationCity",
					State:   "LocationState",
					ZipCode: "54321",
				},
				{
					Name:    "Location2",
					Address: "456 Location Ave",
					City:    "AnotherCity",
					State:   "AnotherState",
					ZipCode: "98765",
				},
			},
		},
	}

	translator := &translation.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientLocations(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientTranslator_TranslateClientProduct(t *testing.T) {
	tests := map[string]struct {
		input    client.Product
		expected *models.Product
	}{
		"Translate Valid Product": {
			input: client.Product{
				ID:   "Product1",
				Name: "Test Product",
			},
			expected: &models.Product{
				Id:   "Product1",
				Name: "Test Product",
			},
		},
	}

	translator := &translation.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientProduct(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientTranslator_TranslateClientProducts(t *testing.T) {
	tests := map[string]struct {
		input    []client.Product
		expected []*models.Product
	}{
		"Translate Multiple Products": {
			input: []client.Product{
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

	translator := &translation.ClientTranslator{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := translator.TranslateClientProducts(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
