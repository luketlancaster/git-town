package steps

import (
	"fmt"

	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/gittools"
)

// CommitOpenChangesStep commits all open changes as a new commit.
// It does not ask the user for a commit message, but chooses one automatically.
type CommitOpenChangesStep struct {
	NoOpStep
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *CommitOpenChangesStep) CreateUndoStepBeforeRun() Step {
	branchName := gitlib.GetCurrentBranchName()
	return &ResetToShaStep{Sha: gittools.GetBranchSha(branchName)}
}

// Run executes this step.
func (step *CommitOpenChangesStep) Run() error {
	err := scriptflows.RunCommand("git", "add", "-A")
	if err != nil {
		return err
	}
	return scriptflows.RunCommand("git", "commit", "-m", fmt.Sprintf("WIP on %s", gitlib.GetCurrentBranchName()))
}
