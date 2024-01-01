package factory

import (
	"github.com/Linkinlog/LeafListr/internal/workflow"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . WorkflowFactory
type WorkflowFactory interface {
	FindByDispensary(dispensary string) (workflow.Workflow, error)
}
