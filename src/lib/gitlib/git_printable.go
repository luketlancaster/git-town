package gitlib

import (
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/Originate/git-town/src/tools/stringtools"
)

// GetPrintableBranchTree returns a user printable branch tree
func GetPrintableBranchTree(branchName string) (result string) {
	result += branchName
	for _, childBranch := range gittools.GetChildBranches(branchName) {
		result += "\n" + stringtools.Indent(GetPrintableBranchTree(childBranch), 1)
	}
	return
}
