package cmd

import (
	"errors"

	"github.com/Originate/git-town/src/flows/gitflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/steps"

	"github.com/spf13/cobra"
)

type prependConfig struct {
	InitialBranch string
	ParentBranch  string
	TargetBranch  string
}

var prependCommand = &cobra.Command{
	Use:   "prepend <branch>",
	Short: "Creates a new feature branch as the parent of the current branch",
	Long: `Creates a new feature branch as the parent of the current branch

Syncs the parent branch
forks a new feature branch with the given name off the parent branch,
makes the new branch the parent of the current branch,
pushes the new feature branch to the remote repository,
and brings over all uncommitted changes to the new feature branch.

Additionally, when there is a remote upstream,
the main branch is synced with its upstream counterpart.
This can be disabled by toggling the "hack-push-flag" configuration:

	git town hack-push-flag false`,
	Run: func(cmd *cobra.Command, args []string) {
		steps.Run(steps.RunOptions{
			CanSkip:              func() bool { return false },
			Command:              "prepend",
			IsAbort:              abortFlag,
			IsContinue:           continueFlag,
			IsSkip:               false,
			IsUndo:               undoFlag,
			SkipMessageGenerator: func() string { return "" },
			StepListGenerator: func() steps.StepList {
				config := getPrependConfig(args)
				return getPrependStepList(config)
			},
		})
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 && !abortFlag && !continueFlag && !undoFlag {
			return errors.New("no branch name provided")
		}
		return errortools.FirstError(
			validateMaxArgsFunc(args, 1),
			gittools.ValidateIsRepository,
			validateIsConfigured,
		)
	},
}

func getPrependConfig(args []string) (result prependConfig) {
	result.InitialBranch = gitlib.GetCurrentBranchName()
	result.TargetBranch = args[0]
	if gittools.HasRemote("origin") && !gittools.IsOffline() {
		scriptlib.Fetch()
	}
	workflows.EnsureDoesNotHaveBranch(result.TargetBranch)
	workflows.EnsureIsFeatureBranch(result.InitialBranch, "Only feature branches can have parent branches.")
	gitflows.EnsureKnowsParentBranches([]string{result.InitialBranch})
	result.ParentBranch = gittools.GetParentBranch(result.InitialBranch)
	return
}

func getPrependStepList(config prependConfig) (result steps.StepList) {
	for _, branchName := range gittools.GetAncestorBranches(config.InitialBranch) {
		result.AppendList(steps.GetSyncBranchSteps(branchName))
	}
	result.Append(&steps.CreateBranchStep{BranchName: config.TargetBranch, StartingPoint: config.ParentBranch})
	result.Append(&steps.SetParentBranchStep{BranchName: config.TargetBranch, ParentBranchName: config.ParentBranch})
	result.Append(&steps.SetParentBranchStep{BranchName: config.InitialBranch, ParentBranchName: config.TargetBranch})
	result.Append(&steps.CheckoutBranchStep{BranchName: config.TargetBranch})
	if gittools.HasRemote("origin") && gittools.ShouldHackPush() && !gittools.IsOffline() {
		result.Append(&steps.CreateTrackingBranchStep{BranchName: config.TargetBranch})
	}
	result.Wrap(steps.WrapOptions{RunInGitRoot: true, StashOpenChanges: true})
	return
}

func init() {
	prependCommand.Flags().BoolVar(&abortFlag, "abort", false, abortFlagDescription)
	prependCommand.Flags().BoolVar(&continueFlag, "continue", false, continueFlagDescription)
	prependCommand.Flags().BoolVar(&undoFlag, "undo", false, undoFlagDescription)
	RootCmd.AddCommand(prependCommand)
}
