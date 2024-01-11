package factory

import (
	"errors"

	"github.com/Linkinlog/LeafListr/internal/beyondhello"
	"github.com/Linkinlog/LeafListr/internal/curaleaf"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

const (
	RepoNotFoundError = "repository not found"
	MenuNotFoundError = "menu not found"
	MenuTypeMismatch  = "menu type mismatch"
)

type DefaultRepositoryFactory struct{}

func NewRepoFactory() RepositoryFactory {
	return &DefaultRepositoryFactory{}
}

func (rf *DefaultRepositoryFactory) FindByDispensary(dispensary, menuType string) (repository.Repository, error) {
	return findRepository(dispensary, menuType)
}

func (rf *DefaultRepositoryFactory) FindByDispensaryMenu(dispensary, menuId, menuType string) (repository.Repository, error) {
	return findRepositoryForMenu(dispensary, menuId, menuType)
}

func findRepositoryForMenu(dispensary string, menuId, menuType string) (repository.Repository, error) {
	repo, err := findRepository(dispensary, menuType)
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

func findRepository(dispensary, menuType string) (repository.Repository, error) {
	var repo repository.Repository
	var err error

	switch dispensary {
	case "curaleaf", "Curaleaf":
		c := curaleaf.NewHTTPClient(
			curaleaf.GqlEndpoint,
		)
		repo = curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), menuType)
	case "beyond", "Beyond", "BeyondHello", "Beyond-Hello", "beyond-hello":
		repo = beyondhello.NewRepository(menuType)
	default:
		err = errors.New("unsupported dispensary")
	}
	return repo, err
}
