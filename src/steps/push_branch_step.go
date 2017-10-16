package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/dryrun"
	"github.com/Originate/git-town/src/tools/gittools"
)

// PushBranchStep pushes the branch with the given name to the origin remote.
// Optionally with force.
type PushBranchStep struct {
	NoOpStep
	BranchName string
	Force      bool
	Undoable   bool
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *PushBranchStep) CreateUndoStepBeforeRun() Step {
	if step.Undoable {
		return &PushBranchAfterCurrentBranchSteps{}
	}
	return &SkipCurrentBranchSteps{}
}

// Run executes this step.
func (step *PushBranchStep) Run() error {
	if !gittools.ShouldBranchBePushed(step.BranchName) && !dryrun.IsActive() {
		return nil
	}
	if step.Force {
		return scriptflows.RunCommand("git", "push", "-f", "origin", step.BranchName)
	}
	if gitlib.GetCurrentBranchName() == step.BranchName {
		return scriptflows.RunCommand("git", "push")
	}
	return scriptflows.RunCommand("git", "push", "origin", step.BranchName)
}
