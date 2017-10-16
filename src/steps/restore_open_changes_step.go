package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// RestoreOpenChangesStep restores stashed away changes into the workspace.
type RestoreOpenChangesStep struct {
	NoOpStep
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *RestoreOpenChangesStep) CreateUndoStepBeforeRun() Step {
	return &StashOpenChangesStep{}
}

// Run executes this step.
func (step *RestoreOpenChangesStep) Run() error {
	return scriptflows.RunCommand("git", "stash", "pop")
}
