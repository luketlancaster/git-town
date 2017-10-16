package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/tools/gittools"
)

// MergeBranchStep merges the branch with the given name into the current branch
type MergeBranchStep struct {
	NoOpStep
	BranchName string
}

// CreateAbortStep returns the abort step for this step.
func (step *MergeBranchStep) CreateAbortStep() Step {
	return &AbortMergeBranchStep{}
}

// CreateContinueStep returns the continue step for this step.
func (step *MergeBranchStep) CreateContinueStep() Step {
	return &ContinueMergeBranchStep{}
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *MergeBranchStep) CreateUndoStepBeforeRun() Step {
	return &ResetToShaStep{Hard: true, Sha: gittools.GetCurrentSha()}
}

// Run executes this step.
func (step *MergeBranchStep) Run() error {
	return scriptflows.RunCommand("git", "merge", "--no-edit", step.BranchName)
}
