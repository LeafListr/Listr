package cache_test

import (
	"testing"
	"time"

	"github.com/Linkinlog/LeafListr/internal/cache"
)

func TestSet(t *testing.T) {
	t.Parallel()
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
			c := cache.NewCache()
			err := c.Set(tc.key, time.Hour, tc.value)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result, getErr := c.Get(tc.key)
			if getErr != nil {
				t.Fatalf("unexpected error: %v", getErr)
			} else if result != tc.value {
				t.Fatalf("expected %v, got %v", tc.value, result)
			}
		})
	}
}
