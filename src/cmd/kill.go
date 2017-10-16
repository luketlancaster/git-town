package cmd

import (
	"fmt"
	"os"

	"github.com/Originate/git-town/src/flows/gitflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/steps"
	"github.com/spf13/cobra"
)

type killConfig struct {
	InitialBranch       string
	IsTargetBranchLocal bool
	TargetBranch        string
}

var killCommand = &cobra.Command{
	Use:   "kill [<branch>]",
	Short: "Removes an obsolete feature branch",
	Long: `Removes an obsolete feature branch

Deletes the current branch, or "<branch_name>" if given,
from the local and remote repositories.

Does not delete perennial branches nor the main branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		steps.Run(steps.RunOptions{
			CanSkip:              func() bool { return false },
			Command:              "kill",
			IsAbort:              false,
			IsContinue:           false,
			IsSkip:               false,
			IsUndo:               undoFlag,
			SkipMessageGenerator: func() string { return "" },
			StepListGenerator: func() steps.StepList {
				return getKillStepList(getKillConfig(args))
			},
		})
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errortools.FirstError(
			validateMaxArgsFunc(args, 1),
			gittools.ValidateIsRepository,
			validateIsConfigured,
		)
	},
}

func getKillConfig(args []string) (result killConfig) {
	result.InitialBranch = gitlib.GetCurrentBranchName()

	if len(args) == 0 {
		result.TargetBranch = result.InitialBranch
	} else {
		result.TargetBranch = args[0]
	}

	workflows.EnsureIsFeatureBranch(result.TargetBranch, "You can only kill feature branches.")

	result.IsTargetBranchLocal = gittools.HasLocalBranch(result.TargetBranch)
	if result.IsTargetBranchLocal {
		gitflows.EnsureKnowsParentBranches([]string{result.TargetBranch})
	}

	if gittools.HasRemote("origin") && !gittools.IsOffline() {
		scriptlib.Fetch()
	}

	if result.InitialBranch != result.TargetBranch {
		workflows.EnsureHasBranch(result.TargetBranch)
	}

	return
}

func getKillStepList(config killConfig) (result steps.StepList) {
	if config.IsTargetBranchLocal {
		targetBranchParent := gittools.GetParentBranch(config.TargetBranch)
		if gittools.HasTrackingBranch(config.TargetBranch) && !gittools.IsOffline() {
			result.Append(&steps.DeleteRemoteBranchStep{BranchName: config.TargetBranch, IsTracking: true})
		}
		if config.InitialBranch == config.TargetBranch {
			if gittools.HasOpenChanges() {
				result.Append(&steps.CommitOpenChangesStep{})
			}
			result.Append(&steps.CheckoutBranchStep{BranchName: targetBranchParent})
		}
		result.Append(&steps.DeleteLocalBranchStep{BranchName: config.TargetBranch, Force: true})
		for _, child := range gittools.GetChildBranches(config.TargetBranch) {
			result.Append(&steps.SetParentBranchStep{BranchName: child, ParentBranchName: targetBranchParent})
		}
		result.Append(&steps.DeleteParentBranchStep{BranchName: config.TargetBranch})
		result.Append(&steps.DeleteAncestorBranchesStep{})
	} else if !gittools.IsOffline() {
		result.Append(&steps.DeleteRemoteBranchStep{BranchName: config.TargetBranch, IsTracking: false})
	} else {
		fmt.Printf("Cannot delete remote branch '%s' in offline mode", config.TargetBranch)
		os.Exit(1)
	}
	result.Wrap(steps.WrapOptions{
		RunInGitRoot:     true,
		StashOpenChanges: config.InitialBranch != config.TargetBranch && config.TargetBranch == gittools.GetPreviouslyCheckedOutBranch(),
	})
	return
}

func init() {
	killCommand.Flags().BoolVar(&undoFlag, "undo", false, undoFlagDescription)
	RootCmd.AddCommand(killCommand)
}
