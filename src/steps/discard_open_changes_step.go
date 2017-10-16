package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// DiscardOpenChangesStep resets the branch to the last commit, discarding uncommitted changes.
type DiscardOpenChangesStep struct {
	NoOpStep
}

// Run executes this step.
func (step *DiscardOpenChangesStep) Run() error {
	return scriptflows.RunCommand("git", "reset", "--hard")
}
