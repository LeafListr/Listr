package factory

import (
	"github.com/Linkinlog/LeafList/internal/repository"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . RepositoryFactory
type RepositoryFactory interface {
	FindByDispensary(dispensary string) (repository.Repository, error)
	FindByDispensaryMenu(dispensary, menuId string) (repository.Repository, error)
}
