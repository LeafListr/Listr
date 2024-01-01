package curaleaf

import (
	"errors"
	"net/http"

	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/factory"

	"github.com/Linkinlog/LeafListr/internal/repository"
)

const (
	RepoNotFoundError = "repository not found"
	MenuNotFoundError = "menu not found"
	MenuTypeMismatch  = "menu type mismatch"
)

type DefaultRepositoryFactory struct {
	memCache cache.Cacher
}

func NewRepoFactory(mC cache.Cacher) factory.RepositoryFactory {
	return &DefaultRepositoryFactory{
		memCache: mC,
	}
}

func (rf *DefaultRepositoryFactory) FindByDispensary(dispensary, menuType string) (repository.Repository, error) {
	return findRepository(dispensary, menuType, rf.memCache)
}

func (rf *DefaultRepositoryFactory) FindByDispensaryMenu(dispensary, menuId, menuType string) (repository.Repository, error) {
	return findRepositoryForMenu(dispensary, menuId, menuType, rf.memCache)
}

func findRepositoryForMenu(dispensary string, menuId, menuType string, mc cache.Cacher) (repository.Repository, error) {
	repo, err := findRepository(dispensary, menuType, mc)
	if err != nil {
		return nil, err
	}

	menu, locationErr := repo.Location(menuId)
	if locationErr != nil {
		return nil, locationErr
	} else if menu == nil || menu.Id != menuId {
		return nil, errors.New(MenuNotFoundError)
	} else if repo == nil {
		return nil, errors.New(RepoNotFoundError)
	}

	return repo, nil
}

func findRepository(dispensary, menuType string, mc cache.Cacher) (repository.Repository, error) {
	var repo repository.Repository
	var err error

	switch dispensary {
	case "curaleaf", "Curaleaf":
		c := NewHTTPClient(
			GqlEndpoint,
			make(http.Header),
		)
		repo = NewRepository(c, NewClientTranslator(), mc, menuType)
	default:
		err = errors.New("unsupported dispensary")
	}
	return repo, err
}
