package flowstwo

import (
	"fmt"

	"github.com/Originate/git-town/src/lib/gitlib"
)

func removePerennialBranch(branchName string) {
	EnsureIsPerennialBranch(branchName, fmt.Sprintf("'%s' is not a perennial branch", branchName))
	gitlib.RemoveFromPerennialBranches(branchName)
}
