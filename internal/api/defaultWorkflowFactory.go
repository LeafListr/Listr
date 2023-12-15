package api

import (
	"errors"

	"github.com/Linkinlog/LeafListr/internal/beyondhello"

	"github.com/Linkinlog/LeafListr/internal/curaleaf"
	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/workflow"
)

type DefaultWorkflowFactory struct{}

func NewWorkflowFactory() factory.WorkflowFactory {
	return &DefaultWorkflowFactory{}
}

func (d *DefaultWorkflowFactory) FindByDispensary(dispensary string) (workflow.Workflow, error) {
	switch dispensary {
	case "curaleaf", "Curaleaf":
		return curaleaf.NewWorkflow(), nil
	case "beyond", "beyond-hello", "Beyond", "Beyond-Hello":
		return beyondhello.NewWorkflow(), nil
	default:
		return nil, errors.New("invalid dispensary")
	}
}
