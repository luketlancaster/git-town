package gitflows

import (
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/gittools"
)

// GetExpectedPreviouslyCheckedOutBranch returns what is the expected previously checked out branch
// given the inputs
func GetExpectedPreviouslyCheckedOutBranch(initialPreviouslyCheckedOutBranch, initialBranch string) string {
	if gitlib.HasLocalBranch(initialPreviouslyCheckedOutBranch) {
		if gitlib.GetCurrentBranchName() == initialBranch || !gitlib.HasLocalBranch(initialBranch) {
			return initialPreviouslyCheckedOutBranch
		}
		return initialBranch
	}
	return gittools.GetMainBranch()
}
