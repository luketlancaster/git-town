package steps

import (
	"github.com/Originate/git-town/src/flows/gitflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/command"
	"github.com/Originate/git-town/src/tools/gittools"
)

// PreserveCheckoutHistoryStep does stuff
type PreserveCheckoutHistoryStep struct {
	NoOpStep
	InitialBranch                     string
	InitialPreviouslyCheckedOutBranch string
}

// Run executes this step.
func (step *PreserveCheckoutHistoryStep) Run() error {
	expectedPreviouslyCheckedOutBranch := gitflows.GetExpectedPreviouslyCheckedOutBranch(step.InitialPreviouslyCheckedOutBranch, step.InitialBranch)
	if expectedPreviouslyCheckedOutBranch != gittools.GetPreviouslyCheckedOutBranch() {
		currentBranch := gitlib.GetCurrentBranchName()
		command.New("git", "checkout", expectedPreviouslyCheckedOutBranch).Run()
		command.New("git", "checkout", currentBranch).Run()
	}
	return nil
}
