package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/tools/gittools"
)

// DeleteLocalBranchStep deletes the branch with the given name,
// optionally in a safe or unsafe way.
type DeleteLocalBranchStep struct {
	NoOpStep
	BranchName string
	Force      bool
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *DeleteLocalBranchStep) CreateUndoStepBeforeRun() Step {
	sha := gittools.GetBranchSha(step.BranchName)
	return &CreateBranchStep{BranchName: step.BranchName, StartingPoint: sha}
}

// Run executes this step.
func (step *DeleteLocalBranchStep) Run() error {
	op := "-d"
	if step.Force || gittools.DoesBranchHaveUnmergedCommits(step.BranchName) {
		op = "-D"
	}
	return scriptflows.RunCommand("git", "branch", op, step.BranchName)
}
