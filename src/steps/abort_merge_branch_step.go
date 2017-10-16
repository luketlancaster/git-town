package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// AbortMergeBranchStep aborts the current merge conflict.
type AbortMergeBranchStep struct {
	NoOpStep
}

// Run executes this step.
func (step *AbortMergeBranchStep) Run() error {
	return scriptflows.RunCommand("git", "merge", "--abort")
}
