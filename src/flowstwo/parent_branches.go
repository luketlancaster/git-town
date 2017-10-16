package flowstwo

import (
	"fmt"

	"github.com/Originate/git-town/src/flows/gitflows"
	"github.com/Originate/git-town/src/flows/promptflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/fatih/color"
)

// EnsureKnowsParentBranches asserts that the entire ancestry for all given branches
// is known to Git Town.
// Missing ancestry information is queried from the user.
func EnsureKnowsParentBranches(branchNames []string) {
	for _, branchName := range branchNames {
		if gittools.IsMainBranch(branchName) || gitlib.IsPerennialBranch(branchName) || gittools.HasCompiledAncestorBranches(branchName) {
			continue
		}
		askForBranchAncestry(branchName)
		ancestors := gitflows.CompileAncestorBranches(branchName)
		gittools.SetAncestorBranches(branchName, ancestors)

		if parentBranchHeaderShown {
			fmt.Println()
		}
	}
}

// Helpers

var parentBranchHeaderShown = false
var parentBranchHeaderTemplate = `
Feature branches can be branched directly off
%s or from other feature branches.

The former allows to develop and ship features completely independent of each other.
The latter allows to build on top of currently unshipped features.

`
var parentBranchPromptTemplate = "Please specify the parent branch of %s by name or number (default: %s): "

func askForBranchAncestry(branchName string) {
	current := branchName
	for {
		parent := gittools.GetParentBranch(current)
		if parent == "" {
			printParentBranchHeader()
			parent = promptflows.AskForBranch(promptflows.BranchPromptConfig{
				branchNames:       gittools.GetLocalBranchesWithMainBranchFirst(),
				defaultBranchName: gittools.GetMainBranch(),
				prompt:            getParentBranchPrompt(current),
				validate: func(branchName string) error {
					return validateParentBranch(current, branchName)
				},
			})
			gittools.SetParentBranch(current, parent)
		}
		if parent == gittools.GetMainBranch() || workflows.IsPerennialBranch(parent) {
			break
		}
		current = parent
	}
}

func validateParentBranch(branchName string, parent string) error {
	if branchName == parent {
		return fmt.Errorf("'%s' cannot be the parent of itself", parent)
	}
	if branchName != "" && gittools.IsAncestorBranch(parent, branchName) {
		return fmt.Errorf("Nested branch loop detected: '%s' is an ancestor of '%s'", branchName, parent)
	}
	return nil
}

func printParentBranchHeader() {
	if !parentBranchHeaderShown {
		parentBranchHeaderShown = true
		cfmt.Printf(parentBranchHeaderTemplate, gittools.GetMainBranch())
		printNumberedBranches(gittools.GetLocalBranchesWithMainBranchFirst())
		fmt.Println()
	}
}

func getParentBranchPrompt(branchName string) string {
	coloredBranchName := color.New(color.Bold).Add(color.FgCyan).Sprintf(branchName)
	return fmt.Sprintf(parentBranchPromptTemplate, coloredBranchName, gittools.GetMainBranch())
}
