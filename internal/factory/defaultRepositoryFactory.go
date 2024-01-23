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

type DefaultRepositoryFactory struct {
	dispensary   string
	menuId       string
	recreational bool
}

func NewRepoFactory(dispensary, menuId string, recreational bool) RepositoryFactory {
	return &DefaultRepositoryFactory{
		dispensary:   dispensary,
		menuId:       menuId,
		recreational: recreational,
	}
}

func (rf *DefaultRepositoryFactory) FindByDispensary() (repository.Repository, error) {
	return findRepository(rf.dispensary, "", rf.recreational)
}

func (rf *DefaultRepositoryFactory) FindByDispensaryMenu() (repository.Repository, error) {
	return findRepositoryForMenu(rf.dispensary, rf.menuId, rf.recreational)
}

func findRepositoryForMenu(dispensary, menuId string, recreational bool) (repository.Repository, error) {
	repo, err := findRepository(dispensary, menuId, recreational)
	if err != nil {
		return nil, err
	}

	menu, locationErr := repo.Location()
	if locationErr != nil {
		return nil, locationErr
	} else if menu == nil || menu.Id != menuId {
		return nil, errors.New(MenuNotFoundError)
	} else if repo == nil {
		return nil, errors.New(RepoNotFoundError)
	}

	return repo, nil
}

func findRepository(dispensary, menuId string, recreational bool) (repository.Repository, error) {
	var repo repository.Repository
	var err error
	switch dispensary {
	case "curaleaf", "Curaleaf":
		c := curaleaf.NewHTTPClient(
			curaleaf.GqlEndpoint,
		)
		repo = curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), menuId, recreational)
	case "beyond", "Beyond", "BeyondHello", "Beyond-Hello", "beyond-hello":
		repo = beyondhello.NewRepository(menuId, recreational)
	default:
		err = errors.New("unsupported dispensary")
	}
	return repo, err
}
