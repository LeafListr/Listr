package factory

import (
	"errors"

	"github.com/Linkinlog/LeafListr/internal/beyondhello"
	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/client"
	"github.com/Linkinlog/LeafListr/internal/curaleaf"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

const (
	RepoNotFoundError = "repository not found"
	MenuNotFoundError = "menu not found"
	MenuTypeMismatch  = "menu type mismatch"
)

func NewRepoFactory(dispensary, menuId string, recreational bool) *RepositoryFactory {
	return &RepositoryFactory{
		dispensary:   dispensary,
		menuId:       menuId,
		recreational: recreational,
	}
}

type RepositoryFactory struct {
	dispensary   string
	menuId       string
	recreational bool
}

func (rf *RepositoryFactory) FindByDispensary() (repository.Repository, error) {
	return findRepository(rf.dispensary, "", rf.recreational)
}

func (rf *RepositoryFactory) FindByDispensaryMenu() (repository.Repository, error) {
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
	c := cache.NewCache()

	switch dispensary {
	case "curaleaf", "Curaleaf":
		c := client.NewHTTPClient(
			curaleaf.GqlEndpoint,
			curaleaf.Headers,
			c,
		)
		repo = curaleaf.NewRepository(c, curaleaf.NewClientTranslator(), menuId, recreational)
	case "beyond", "Beyond", "BeyondHello", "Beyond-Hello", "beyond-hello":
		c := client.NewHTTPClient(
			beyondhello.BeyondEndpoint,
			beyondhello.Headers,
			c,
		)
		repo = beyondhello.NewRepository(menuId, recreational, c)
	default:
		err = errors.New("unsupported dispensary")
	}
	return repo, err
}
