package factory

import (
	"github.com/Linkinlog/LeafListr/internal/repository"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . RepositoryFactory
type RepositoryFactory interface {
	FindByDispensary(dispensary, menuType string) (repository.Repository, error)
	FindByDispensaryMenu(dispensary, menuId, menuType string) (repository.Repository, error)
}
