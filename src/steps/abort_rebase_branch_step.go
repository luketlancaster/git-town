package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// AbortRebaseBranchStep represents aborting on ongoing merge conflict.
// This step is used in the abort scripts for Git Town commands.
type AbortRebaseBranchStep struct {
	NoOpStep
}

// Run executes this step.
func (step *AbortRebaseBranchStep) Run() error {
	return scriptflows.RunCommand("git", "rebase", "--abort")
}
