package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// CreateTrackingBranchStep pushes the current branch up to origin
// and marks it as tracking the current branch.
type CreateTrackingBranchStep struct {
	NoOpStep
	BranchName string
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *CreateTrackingBranchStep) CreateUndoStepBeforeRun() Step {
	return &DeleteRemoteBranchStep{BranchName: step.BranchName}
}

// Run executes this step.
func (step *CreateTrackingBranchStep) Run() error {
	return scriptflows.RunCommand("git", "push", "-u", "origin", step.BranchName)
}
