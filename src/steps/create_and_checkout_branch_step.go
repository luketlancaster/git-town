package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/tools/gittools"
)

// CreateAndCheckoutBranchStep creates a new branch and makes it the current one.
type CreateAndCheckoutBranchStep struct {
	NoOpStep
	BranchName       string
	ParentBranchName string
}

// Run executes this step.
func (step *CreateAndCheckoutBranchStep) Run() error {
	gittools.SetParentBranch(step.BranchName, step.ParentBranchName)
	return scriptflows.RunCommand("git", "checkout", "-b", step.BranchName, step.ParentBranchName)
}
