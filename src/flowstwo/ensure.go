package flowstwo

import (
	"fmt"

	"github.com/Originate/git-town/src/flows"
	"github.com/Originate/git-town/src/flows/promptflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/gittools"
)

// Ensure asserts that the given condition is true.
// If not, it ends the application with the given error message.
func Ensure(condition bool, error string) {
	if !condition {
		flows.ExitWithErrorMessage(error)
	}
}

// EnsureBranchInSync enforces that a branch with the given name is in sync with its tracking branch
func EnsureBranchInSync(branchName, errorMessageSuffix string) {
	Ensure(gittools.IsBranchInSync(branchName), fmt.Sprintf("'%s' is not in sync with its tracking branch. %s", branchName, errorMessageSuffix))
}

// EnsureDoesNotHaveBranch enforces that a branch with the given name does not exist
func EnsureDoesNotHaveBranch(branchName string) {
	Ensure(!gittools.HasBranch(branchName), fmt.Sprintf("A branch named '%s' already exists", branchName))
}

// EnsureHasBranch enforces that a branch with the given name exists
func EnsureHasBranch(branchName string) {
	Ensure(gittools.HasBranch(branchName), fmt.Sprintf("There is no branch named '%s'", branchName))
}

// EnsureIsNotMainBranch enforces that a branch with the given name is not the main branch
func EnsureIsNotMainBranch(branchName, errorMessage string) {
	Ensure(!gittools.IsMainBranch(branchName), errorMessage)
}

// EnsureIsNotPerennialBranch enforces that a branch with the given name is not a perennial branch
func EnsureIsNotPerennialBranch(branchName, errorMessage string) {
	Ensure(!gitlib.IsPerennialBranch(branchName), errorMessage)
}

// EnsureIsPerennialBranch enforces that a branch with the given name is a perennial branch
func EnsureIsPerennialBranch(branchName, errorMessage string) {
	Ensure(gitlib.IsPerennialBranch(branchName), errorMessage)
}

// EnsureDoesNotHaveConflicts asserts that the workspace
// has no unresolved merge conflicts.
func EnsureDoesNotHaveConflicts() {
	Ensure(!gittools.HasConflicts(), "You must resolve the conflicts before continuing")
}

// EnsureDoesNotHaveOpenChanges assets that the workspace
// has no open changes
func EnsureDoesNotHaveOpenChanges(message string) {
	Ensure(!gittools.HasOpenChanges(), "You have uncommitted changes. "+message)
}

// EnsureVersionRequirementSatisfied asserts that Git is the needed version or higher
func EnsureVersionRequirementSatisfied() {
	Ensure(gittools.IsVersionRequirementSatisfied(), "Git Town requires Git 2.7.0 or higher")
}

// EnsureIsFeatureBranch asserts that the given branch is a feature branch.
func EnsureIsFeatureBranch(branchName, errorSuffix string) {
	Ensure(gitlib.IsFeatureBranch(branchName), fmt.Sprintf("The branch '%s' is not a feature branch. %s", branchName, errorSuffix))
}

// EnsureIsConfigured has the user to confgure the main branch and perennial branches if needed
func EnsureIsConfigured() {
	if gittools.GetMainBranch() == "" {
		promptflows.ConfigureMainBranch()
		promptflows.ConfigurePerennialBranches()
	}
}
