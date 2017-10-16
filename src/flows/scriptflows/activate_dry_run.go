package scriptflows

import (
	"github.com/Originate/git-town/src/exit"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/dryrun"
	"github.com/fatih/color"
)

// ActivateDryRun causes all commands to not be run
func ActivateDryRun() {
	_, err := color.New(color.FgBlue).Print(dryRunMessage)
	exit.On(err)
	dryrun.Activate(gitlib.GetCurrentBranchName())
}
