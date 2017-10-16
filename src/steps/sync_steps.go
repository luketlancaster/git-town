package steps

import (
	"fmt"

	"github.com/Originate/git-town/src/lib/gitlib"
)

// GetSyncBranchSteps returns the steps to sync the branch with the given name.
func GetSyncBranchSteps(branchName string) (result StepList) {
	isFeature := gitlib.IsFeatureBranch(branchName)
	hasRemoteOrigin := gittools.HasRemote("origin")

	if !hasRemoteOrigin && !isFeature {
		return
	}

	result.Append(&CheckoutBranchStep{BranchName: branchName})
	if isFeature {
		result.Append(&MergeTrackingBranchStep{})
		result.Append(&MergeBranchStep{BranchName: gittools.GetParentBranch(branchName)})
	} else {
		if gittools.GetPullBranchStrategy() == "rebase" {
			result.Append(&RebaseTrackingBranchStep{})
		} else {
			result.Append(&MergeTrackingBranchStep{})
		}

		mainBranchName := gittools.GetMainBranch()
		if mainBranchName == branchName && gittools.HasRemote("upstream") {
			result.Append(&FetchUpstreamStep{})
			result.Append(&RebaseBranchStep{BranchName: fmt.Sprintf("upstream/%s", mainBranchName)})
		}
	}

	if hasRemoteOrigin && !gittools.IsOffline() {
		if gittools.HasTrackingBranch(branchName) {
			result.Append(&PushBranchStep{BranchName: branchName})
		} else {
			result.Append(&CreateTrackingBranchStep{BranchName: branchName})
		}
	}

	return
}
