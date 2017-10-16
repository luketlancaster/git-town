/*
This file contains functionality around storing configuration settings
inside Git's metadata storage for the repository.
*/

package gitlib

import (
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/Originate/git-town/src/tools/stringtools"
)

// IsPerennialBranch returns whether the branch with the given name is
// a perennial branch.
func IsPerennialBranch(branchName string) bool {
	perennialBranches := gittools.GetPerennialBranches()
	return stringtools.DoesStringArrayContain(perennialBranches, branchName)
}

// RemoveFromPerennialBranches removes the given branch as a perennial branch
func RemoveFromPerennialBranches(branchName string) {
	gittools.SetPerennialBranches(stringtools.RemoveStringFromSlice(gittools.GetPerennialBranches(), branchName))
}

// ShouldHackPush returns whether the current repository is configured to push
// freshly created branches up to the origin remote.
func ShouldHackPush() bool {
	return stringtools.StringToBool(gittools.GetLocalConfigurationValueWithDefault("git-town.hack-push-flag", "false"))
}
