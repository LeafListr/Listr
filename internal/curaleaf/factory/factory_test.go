package factory_test

import (
	"github.com/Linkinlog/LeafListr/internal/curaleaf/factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByDispensary(t *testing.T) {
	tests := map[string]struct {
		dispensary string
		wantErr    bool
		errorMsg   string
	}{
		"invalid dispensary": {
			dispensary: "Delta9",
			wantErr:    true,
			errorMsg:   "unsupported dispensary",
		},
		"valid dispensary": {
			dispensary: "Curaleaf",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repoFactory := factory.NewRepoFactory()
			repo, err := repoFactory.FindByDispensary(tt.dispensary)
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
		wantErr    bool
		errorMsg   string
	}{
		"invalid dispensary": {
			dispensary: "Delta9",
			wantErr:    true,
			errorMsg:   "unsupported dispensary",
		},
		"invalid location": {
			dispensary: "Curaleaf",
			location:   "Banana",
			wantErr:    true,
			errorMsg:   "menu not found",
		},
		"valid dispensary": {
			dispensary: "Curaleaf",
			location:   "LMR070",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repoFactory := factory.NewRepoFactory()
			repo, err := repoFactory.FindByDispensaryMenu(tt.dispensary, tt.location)
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
