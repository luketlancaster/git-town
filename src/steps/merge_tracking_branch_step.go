package steps

import (
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/gittools"
)

// MergeTrackingBranchStep merges the tracking branch of the current branch
// into the current branch.
type MergeTrackingBranchStep struct {
	NoOpStep
}

// CreateAbortStep returns the abort step for this step.
func (step *MergeTrackingBranchStep) CreateAbortStep() Step {
	return &AbortMergeBranchStep{}
}

// CreateContinueStep returns the continue step for this step.
func (step *MergeTrackingBranchStep) CreateContinueStep() Step {
	return &ContinueMergeBranchStep{}
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *MergeTrackingBranchStep) CreateUndoStepBeforeRun() Step {
	return &ResetToShaStep{Hard: true, Sha: gittools.GetCurrentSha()}
}

// Run executes this step.
func (step *MergeTrackingBranchStep) Run() error {
	branchName := gitlib.GetCurrentBranchName()
	if gittools.HasTrackingBranch(branchName) {
		return (&MergeBranchStep{BranchName: gittools.GetTrackingBranchName(branchName)}).Run()
	}
	return nil
}
