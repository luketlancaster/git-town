package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/tools/gittools"
)

// DeleteRemoteBranchStep deletes the current branch from the origin remote.
type DeleteRemoteBranchStep struct {
	NoOpStep
	BranchName string
	IsTracking bool
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *DeleteRemoteBranchStep) CreateUndoStepBeforeRun() Step {
	if step.IsTracking {
		return &CreateTrackingBranchStep{BranchName: step.BranchName}
	}
	sha := gittools.GetBranchSha(gittools.GetTrackingBranchName(step.BranchName))
	return &CreateRemoteBranchStep{BranchName: step.BranchName, Sha: sha}
}

// Run executes this step.
func (step *DeleteRemoteBranchStep) Run() error {
	return scriptflows.RunCommand("git", "push", "origin", ":"+step.BranchName)
}
