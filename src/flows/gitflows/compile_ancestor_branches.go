package gitflows

import (
	"errors"

	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/Originate/git-town/src/tools/stringtools"
)

// CompileAncestorBranches calculates and returns the list of ancestor branches
// of the given branch based off the "git-town-branch.XXX.parent" configuration values.
func CompileAncestorBranches(branchName string) (result []string) {
	current := branchName
	for {
		if gittools.IsMainBranch(current) || gitlib.IsPerennialBranch(current) {
			return
		}
		parent := gittools.GetParentBranch(current)
		if parent == "" {
			return
		}
		result = append([]string{parent}, result...)
		current = parent
	}
}

// ValidateIsOnline asserts that Git Town is not in offline mode
func ValidateIsOnline() error {
	if gitlib.IsOffline() {
		return errors.New("This command requires an active internet connection")
	}
	return nil
}

// IsAncestorBranch returns whether the given branch is an ancestor of the other given branch.
func IsAncestorBranch(branchName, ancestorBranchName string) bool {
	ancestorBranches := CompileAncestorBranches(branchName)
	return stringtools.DoesStringArrayContain(ancestorBranches, ancestorBranchName)
}
