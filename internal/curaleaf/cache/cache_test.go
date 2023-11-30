package cache_test

import (
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/cache"
)

func TestSet(t *testing.T) {
	tests := map[string]struct {
		key   string
		value any
	}{
		"valid set": {
			key:   "key",
			value: "value",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache := cache.NewCache()
			err := cache.Set(tc.key, tc.value)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, err := cache.Get(tc.key)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			} else if result != tc.value {
				t.Fatalf("expected %v, got %v", tc.value, result)
			}
		})
	}
}
