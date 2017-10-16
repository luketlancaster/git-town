package cmd

import (
	"strings"

	"github.com/Originate/git-town/src/drivers"
	"github.com/Originate/git-town/src/exit"
	"github.com/Originate/git-town/src/flows"
	"github.com/Originate/git-town/src/flows/gitflows"
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/steps"

	"github.com/spf13/cobra"
)

type shipConfig struct {
	BranchToShip  string
	InitialBranch string
}

var commitMessage string

var shipCmd = &cobra.Command{
	Use:   "ship",
	Short: "Deliver a completed feature branch",
	Long: `Deliver a completed feature branch

Squash-merges the current branch, or <branch_name> if given,
into the main branch, resulting in linear history on the main branch.

- syncs the main branch
- pulls remote updates for <branch_name>
- merges the main branch into <branch_name>
- squash-merges <branch_name> into the main branch
  with commit message specified by the user
- pushes the main branch to the remote repository
- deletes <branch_name> from the local and remote repositories

Only shipping of direct children of the main branch is allowed.
To ship a nested child branch, all ancestor branches have to be shipped or killed.

If you are using GitHub, this command can squash merge pull requests via the GitHub API. Setup:
1. Get a GitHub personal access token with the "repo" scope
2. Run 'git config git-town.github-token XXX' (optionally add the '--global' flag)
Now anytime you ship a branch with a pull request on GitHub, it will squash merge via the GitHub API.
It will also update the base branch for any pull requests against that branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		steps.Run(steps.RunOptions{
			CanSkip:              func() bool { return false },
			Command:              "ship",
			IsAbort:              abortFlag,
			IsContinue:           continueFlag,
			IsSkip:               false,
			IsUndo:               undoFlag,
			SkipMessageGenerator: func() string { return "" },
			StepListGenerator: func() steps.StepList {
				config := gitShipConfig(args)
				return getShipStepList(config)
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

func gitShipConfig(args []string) (result shipConfig) {
	result.InitialBranch = gitlib.GetCurrentBranchName()
	if len(args) == 0 {
		result.BranchToShip = result.InitialBranch
	} else {
		result.BranchToShip = args[0]
	}
	if result.BranchToShip == result.InitialBranch {
		workflows.EnsureDoesNotHaveOpenChanges("Did you mean to commit them before shipping?")
	}
	if gittools.HasRemote("origin") && !gittools.IsOffline() {
		scriptflows.Fetch()
	}
	if result.BranchToShip != result.InitialBranch {
		workflows.EnsureHasBranch(result.BranchToShip)
	}
	workflows.EnsureIsFeatureBranch(result.BranchToShip, "Only feature branches can be shipped.")
	gitflows.EnsureKnowsParentBranches([]string{result.BranchToShip})
	ensureParentBranchIsMainOrPerennialBranch(result.BranchToShip)
	return
}

func ensureParentBranchIsMainOrPerennialBranch(branchName string) {
	parentBranch := gittools.GetParentBranch(branchName)
	if !gittools.IsMainBranch(parentBranch) && !workflows.IsPerennialBranch(parentBranch) {
		ancestors := gittools.GetAncestorBranches(branchName)
		ancestorsWithoutMainOrPerennial := ancestors[1:]
		oldestAncestor := ancestorsWithoutMainOrPerennial[0]
		flows.ExitWithErrorMessage(
			"Shipping this branch would ship "+strings.Join(ancestorsWithoutMainOrPerennial, ", ")+" as well.",
			"Please ship \""+oldestAncestor+"\" first.",
		)
	}
}

func getShipStepList(config shipConfig) (result steps.StepList) {
	var isOffline = gittools.IsOffline()
	branchToMergeInto := gittools.GetParentBranch(config.BranchToShip)
	isShippingInitialBranch := config.BranchToShip == config.InitialBranch
	result.AppendList(steps.GetSyncBranchSteps(branchToMergeInto))
	result.Append(&steps.CheckoutBranchStep{BranchName: config.BranchToShip})
	result.Append(&steps.MergeTrackingBranchStep{})
	result.Append(&steps.MergeBranchStep{BranchName: branchToMergeInto})
	result.Append(&steps.EnsureHasShippableChangesStep{BranchName: config.BranchToShip})
	result.Append(&steps.CheckoutBranchStep{BranchName: branchToMergeInto})
	canShipWithDriver, defaultCommitMessage := getCanShipWithDriver(config.BranchToShip, branchToMergeInto)
	if canShipWithDriver {
		result.Append(&steps.DriverMergePullRequestStep{BranchName: config.BranchToShip, CommitMessage: commitMessage, DefaultCommitMessage: defaultCommitMessage})
		result.Append(&steps.PullBranchStep{})
	} else {
		result.Append(&steps.SquashMergeBranchStep{BranchName: config.BranchToShip, CommitMessage: commitMessage})
	}
	if gittools.HasRemote("origin") && !isOffline {
		result.Append(&steps.PushBranchStep{BranchName: branchToMergeInto, Undoable: true})
	}
	childBranches := gittools.GetChildBranches(config.BranchToShip)
	if canShipWithDriver || (gittools.HasTrackingBranch(config.BranchToShip) && len(childBranches) == 0 && !isOffline) {
		result.Append(&steps.DeleteRemoteBranchStep{BranchName: config.BranchToShip, IsTracking: true})
	}
	result.Append(&steps.DeleteLocalBranchStep{BranchName: config.BranchToShip})
	result.Append(&steps.DeleteParentBranchStep{BranchName: config.BranchToShip})
	for _, child := range childBranches {
		result.Append(&steps.SetParentBranchStep{BranchName: child, ParentBranchName: branchToMergeInto})
	}
	result.Append(&steps.DeleteAncestorBranchesStep{})
	if !isShippingInitialBranch {
		result.Append(&steps.CheckoutBranchStep{BranchName: config.InitialBranch})
	}
	result.Wrap(steps.WrapOptions{RunInGitRoot: true, StashOpenChanges: !isShippingInitialBranch})
	return
}

func getCanShipWithDriver(branch, parentBranch string) (bool, string) {
	if !gittools.HasRemote("origin") {
		return false, ""
	}
	if gittools.IsOffline() {
		return false, ""
	}
	driver := drivers.GetActiveDriver()
	if driver == nil {
		return false, ""
	}
	canMerge, defaultCommitMessage, err := driver.CanMergePullRequest(branch, parentBranch)
	exit.On(err)
	return canMerge, defaultCommitMessage
}

func init() {
	shipCmd.Flags().BoolVar(&abortFlag, "abort", false, abortFlagDescription)
	shipCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Specify the commit message for the squash commit")
	shipCmd.Flags().BoolVar(&continueFlag, "continue", false, continueFlagDescription)
	shipCmd.Flags().BoolVar(&undoFlag, "undo", false, undoFlagDescription)
	RootCmd.AddCommand(shipCmd)
}
