package workflow

import (
	"github.com/Linkinlog/LeafListr/internal/cache"
	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/transformation"
)

type defaultWorkflow struct {
	rf factory.RepositoryFactory
	c  cache.Cacher
	f  transformation.Filterer
}
