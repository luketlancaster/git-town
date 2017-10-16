package cmd

import (
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/steps"
	"github.com/spf13/cobra"
)

var pruneBranchesCommand = &cobra.Command{
	Use:   "prune-branches",
	Short: "Deletes local branches whose tracking branch no longer exists",
	Long: `Deletes local branches whose tracking branch no longer exists

Deletes branches whose tracking branch no longer exists from the local repository.
This usually means the branch was shipped or killed on another machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		steps.Run(steps.RunOptions{
			CanSkip:              func() bool { return false },
			Command:              "prune-branches",
			IsAbort:              false,
			IsContinue:           false,
			IsSkip:               false,
			IsUndo:               undoFlag,
			SkipMessageGenerator: func() string { return "" },
			StepListGenerator: func() steps.StepList {
				checkPruneBranchesPreconditions()
				return getPruneBranchesStepList()
			},
		})
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errortools.FirstError(
			validateMaxArgsFunc(args, 0),
			gittools.ValidateIsRepository,
			validateIsConfigured,
			gittools.ValidateIsOnline,
		)
	},
}

func checkPruneBranchesPreconditions() {
	if gittools.HasRemote("origin") {
		scriptlib.Fetch()
	}
}

func getPruneBranchesStepList() (result steps.StepList) {
	initialBranchName := gitlib.GetCurrentBranchName()
	for _, branchName := range gittools.GetLocalBranchesWithDeletedTrackingBranches() {
		if initialBranchName == branchName {
			result.Append(&steps.CheckoutBranchStep{BranchName: gittools.GetMainBranch()})
		}

		parent := gittools.GetParentBranch(branchName)
		if parent != "" {
			for _, child := range gittools.GetChildBranches(branchName) {
				result.Append(&steps.SetParentBranchStep{BranchName: child, ParentBranchName: parent})
			}
			result.Append(&steps.DeleteParentBranchStep{BranchName: branchName})
			result.Append(&steps.DeleteAncestorBranchesStep{})
		}

		result.Append(&steps.DeleteLocalBranchStep{BranchName: branchName})
	}
	result.Wrap(steps.WrapOptions{RunInGitRoot: false, StashOpenChanges: false})
	return
}

func init() {
	pruneBranchesCommand.Flags().BoolVar(&undoFlag, "undo", false, undoFlagDescription)
	RootCmd.AddCommand(pruneBranchesCommand)
}
