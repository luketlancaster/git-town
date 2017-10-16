package gitlib

import "github.com/Originate/git-town/src/tools/gittools"

// IsFeatureBranch returns whether the branch with the given name is
// a feature branch.
func IsFeatureBranch(branchName string) bool {
	return !gittools.IsMainBranch(branchName) && !IsPerennialBranch(branchName)
}
