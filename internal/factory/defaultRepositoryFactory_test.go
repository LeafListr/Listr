package factory_test

import (
	"testing"

	"github.com/Linkinlog/LeafListr/internal/factory"

	"github.com/stretchr/testify/assert"
)

func TestFindByDispensary(t *testing.T) {
	tests := map[string]struct {
		dispensary string
		rec        bool
		wantErr    bool
		errorMsg   string
	}{
		"invalid dispensary": {
			dispensary: "Delta9",
			rec:        false,
			wantErr:    true,
			errorMsg:   "unsupported dispensary",
		},
		"valid dispensary": {
			dispensary: "Curaleaf",
			rec:        false,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repoFactory := factory.NewRepoFactory(tt.dispensary, "", tt.rec)
			repo, err := repoFactory.FindByDispensary()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, repo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo)
			}
		})
	}
}

func TestFindByDispensaryMenu(t *testing.T) {
	tests := map[string]struct {
		dispensary string
		location   string
		rec        bool
		wantErr    bool
		errorMsg   string
	}{
		"invalid dispensary": {
			dispensary: "Delta9",
			rec:        false,
			wantErr:    true,
			errorMsg:   "unsupported dispensary",
		},
		"invalid location": {
			dispensary: "Curaleaf",
			location:   "Banana",
			rec:        false,
			wantErr:    true,
			errorMsg:   "menu not found",
		},
		"valid dispensary": {
			dispensary: "Curaleaf",
			location:   "LMR070",
			rec:        false,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repoFactory := factory.NewRepoFactory(tt.dispensary, tt.location, tt.rec)
			repo, err := repoFactory.FindByDispensaryMenu()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, repo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo)
			}
		})
	}
}
