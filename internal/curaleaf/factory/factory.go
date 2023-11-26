package factory

import (
	"errors"

	"github.com/Linkinlog/LeafListr/internal/curaleaf/client"
	curarepo "github.com/Linkinlog/LeafListr/internal/curaleaf/repository"
	"github.com/Linkinlog/LeafListr/internal/curaleaf/translation"
	"github.com/Linkinlog/LeafListr/internal/factory"

	"github.com/Linkinlog/LeafListr/internal/repository"
)

const (
	RepoNotFoundError = "repository not found"
	MenuNotFoundError = "menu not found"
)

type DefaultRepositoryFactory struct{}

func NewRepoFactory() factory.RepositoryFactory {
	return &DefaultRepositoryFactory{}
}

func (rf *DefaultRepositoryFactory) FindByDispensary(dispensary string) (repository.Repository, error) {
	return findRepository(dispensary)
}

func (rf *DefaultRepositoryFactory) FindByDispensaryMenu(dispensary, menuId string) (repository.Repository, error) {
	return findMenu(dispensary, menuId)
}

func findMenu(dispensary string, menuId string) (repository.Repository, error) {
	repo, err := findRepository(dispensary)
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

func findRepository(dispensary string) (repository.Repository, error) {
	var repo repository.Repository
	var err error

	switch dispensary {
	case "curaleaf", "Curaleaf":
		c := client.NewHTTPClient(
			curarepo.GqlEndpoint,
			curarepo.Headers,
		)
		repo = curarepo.NewRepository(c, translation.NewClientTranslator())
	default:
		err = errors.New("unsupported dispensary")
	}
	return repo, err
}
