package factory

import (
	"github.com/Linkinlog/LeafList/internal/repository"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . RepositoryFactory
type RepositoryFactory interface {
	FindRepository(dispensary string) (repository.Repository, error)
	FindMenu(dispensary, menuId string) (repository.Repository, error)
}
