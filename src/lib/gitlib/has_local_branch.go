package gitlib

import (
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/Originate/git-town/src/tools/stringtools"
)

// HasLocalBranch returns whether the local repository contains
// a branch with the given name.
func HasLocalBranch(branchName string) bool {
	return stringtools.DoesStringArrayContain(gittools.GetLocalBranches(), branchName)
}
