package curaleaf_test

import (
	"testing"

	"github.com/Linkinlog/LeafListr/internal/curaleaf"

	"github.com/stretchr/testify/assert"
)

func TestFindByDispensary(t *testing.T) {
	tests := map[string]struct {
		dispensary string
		menuType   string
		wantErr    bool
		errorMsg   string
	}{
		"invalid dispensary": {
			dispensary: "Delta9",
			menuType:   "medical",
			wantErr:    true,
			errorMsg:   "unsupported dispensary",
		},
		"valid dispensary": {
			dispensary: "Curaleaf",
			menuType:   "medical",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repoFactory := curaleaf.NewRepoFactory(curaleaf.NewCache())
			repo, err := repoFactory.FindByDispensary(tt.dispensary, tt.menuType)
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
		menuType   string
		wantErr    bool
		errorMsg   string
	}{
		"invalid dispensary": {
			dispensary: "Delta9",
			menuType:   "medical",
			wantErr:    true,
			errorMsg:   "unsupported dispensary",
		},
		"invalid location": {
			dispensary: "Curaleaf",
			location:   "Banana",
			menuType:   "medical",
			wantErr:    true,
			errorMsg:   "menu not found",
		},
		"valid dispensary": {
			dispensary: "Curaleaf",
			location:   "LMR070",
			menuType:   "medical",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repoFactory := curaleaf.NewRepoFactory(curaleaf.NewCache())
			repo, err := repoFactory.FindByDispensaryMenu(tt.dispensary, tt.location, tt.menuType)
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
