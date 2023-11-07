package factory

import (
	"errors"
	"github.com/Linkinlog/LeafList/internal/client"
	curarepo "github.com/Linkinlog/LeafList/internal/repository/curaleaf"

	"github.com/Linkinlog/LeafList/internal/repository"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

const RepoNotFoundError = "repository not found"

//counterfeiter:generate . RepositoryFactory
type RepositoryFactory interface {
	Find(dispensary, menuId string) (repository.Repository, error)
}

type DefaultRepositoryFactory struct{}

func NewRepoFactory() *DefaultRepositoryFactory {
	return &DefaultRepositoryFactory{}
}

func (rf *DefaultRepositoryFactory) Find(dispensary, menuId string) (repository.Repository, error) {
	var repo repository.Repository

	switch dispensary {
	case "curaleaf", "Curaleaf":
		c := client.NewHTTPClient(
			curarepo.GqlEndpoint,
			curarepo.Headers,
		)
		repo = curarepo.NewRepository(c)
	}

	if repo != nil {
		return repo, nil
	}
	return nil, errors.New(RepoNotFoundError)
}

//func curaSwitch(menuId string) repository.Repository {
//	var repo repository.Repository
//
//	switch menuId {
//	case curaleaf.GbgId, "LMR094":
//		c := client.NewHTTPClient(
//			curaleaf.GqlEndpoint,
//			curaleaf.Headers,
//		)
//		repo = curarepo.NewRepository(c)
//	}
//
//	return repo
//}
