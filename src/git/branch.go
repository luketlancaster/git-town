package git

import (
	"github.com/Originate/git-town/src/command"
	"github.com/Originate/git-town/src/dryrun"
)

// GetCurrentBranchName returns the name of the currently checked out branch
func GetCurrentBranchName() string {
	if dryrun.IsActive() {
		return dryrun.GetCurrentBranchName()
	}
	if IsRebaseInProgress() {
		return getCurrentBranchNameDuringRebase()
	}
	return command.New("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
}
