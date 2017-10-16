package promptlib

import (
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/fatih/color"
)

// PrintNumberedBranches prints the given branches to the console,
// preceded by numbers
func PrintNumberedBranches(branchNames []string) {
	boldFmt := color.New(color.Bold)
	for index, branchName := range branchNames {
		cfmt.Printf("  %s: %s\n", boldFmt.Sprintf("%d", index+1), branchName)
	}
}
