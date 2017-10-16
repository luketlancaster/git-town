package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// PushTagsStep pushes newly created Git tags to the remote.
type PushTagsStep struct {
	NoOpStep
}

// Run executes this step.
func (step *PushTagsStep) Run() error {
	return scriptflows.RunCommand("git", "push", "--tags")
}
