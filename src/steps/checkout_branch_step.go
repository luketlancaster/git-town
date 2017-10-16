package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/lib/gitlib"
)

// CheckoutBranchStep checks out a new branch.
type CheckoutBranchStep struct {
	NoOpStep
	BranchName string
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *CheckoutBranchStep) CreateUndoStepBeforeRun() Step {
	return &CheckoutBranchStep{BranchName: gitlib.GetCurrentBranchName()}
}

// Run executes this step.
func (step *CheckoutBranchStep) Run() error {
	if gitlib.GetCurrentBranchName() != step.BranchName {
		return scriptflows.RunCommand("git", "checkout", step.BranchName)
	}
	return nil
}
