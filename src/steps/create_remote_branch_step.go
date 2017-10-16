package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// CreateRemoteBranchStep pushes the current branch up to origin.
type CreateRemoteBranchStep struct {
	NoOpStep
	BranchName string
	Sha        string
}

// Run executes this step.
func (step *CreateRemoteBranchStep) Run() error {
	return scriptflows.RunCommand("git", "push", "origin", step.Sha+":refs/heads/"+step.BranchName)
}
