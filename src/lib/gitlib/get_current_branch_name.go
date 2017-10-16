package gitlib

import (
	"github.com/Originate/git-town/src/tools/command"
	"github.com/Originate/git-town/src/tools/dryrun"
	"github.com/Originate/git-town/src/tools/gittools"
)

// GetCurrentBranchName returns the name of the currently checked out branch
func GetCurrentBranchName() string {
	if dryrun.IsActive() {
		return dryrun.GetCurrentBranchName()
	}
	if gittools.IsRebaseInProgress() {
		return gittools.GetCurrentBranchNameDuringRebase()
	}
	return command.New("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
}
