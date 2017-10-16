package steps

import "github.com/Originate/git-town/src/tools/gittools"

// DeleteAncestorBranchesStep removes all ancestor information
// for the current branch
// from the Git Town configuration.
type DeleteAncestorBranchesStep struct {
	NoOpStep
}

// Run executes this step.
func (step *DeleteAncestorBranchesStep) Run() error {
	gittools.DeleteAllAncestorBranches()
	return nil
}
