package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/tools/gittools"
)

// RebaseBranchStep rebases the current branch
// against the branch with the given name.
type RebaseBranchStep struct {
	NoOpStep
	BranchName string
}

// CreateAbortStep returns the abort step for this step.
func (step *RebaseBranchStep) CreateAbortStep() Step {
	return &AbortRebaseBranchStep{}
}

// CreateContinueStep returns the continue step for this step.
func (step *RebaseBranchStep) CreateContinueStep() Step {
	return &ContinueRebaseBranchStep{}
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *RebaseBranchStep) CreateUndoStepBeforeRun() Step {
	return &ResetToShaStep{Hard: true, Sha: gittools.GetCurrentSha()}
}

// Run executes this step.
func (step *RebaseBranchStep) Run() error {
	return scriptflows.RunCommand("git", "rebase", step.BranchName)
}
