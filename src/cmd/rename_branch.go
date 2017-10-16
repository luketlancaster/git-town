package cmd

import (
	"errors"
	"fmt"

	"github.com/Originate/git-town/src/flows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/steps"
	"github.com/spf13/cobra"
)

type renameBranchConfig struct {
	OldBranchName string
	NewBranchName string
}

var forceFlag bool

var renameBranchCommand = &cobra.Command{
	Use:   "rename-branch [<old_branch_name>] <new_branch_name>",
	Short: "Renames a branch both locally and remotely",
	Long: `Renames a branch both locally and remotely

Renames the given branch on both the local machine and the remote if one is configured.
Aborts if the new branch name already exists or the tracking branch is out of sync.
This command is intended for feature branches.
Renaming perennial branches has to be confirmed with the "-f" option.

- Creates a branch with the new name
- Deletes the old branch

When there is a remote repository
- Syncs the repository

When there is a tracking branch
- Pushes the new branch to the remote repository
- Deletes the old branch from the remote repository

When run on a perennial branch
- Requires the use of the "-f" option
- Reconfigures git town locally for the perennial branch`,
	Run: func(cmd *cobra.Command, args []string) {
		steps.Run(steps.RunOptions{
			CanSkip:              func() bool { return false },
			Command:              "rename-branch",
			IsAbort:              false,
			IsContinue:           false,
			IsSkip:               false,
			IsUndo:               undoFlag,
			SkipMessageGenerator: func() string { return "" },
			StepListGenerator: func() steps.StepList {
				config := getRenameBranchConfig(args)
				return getRenameBranchStepList(config)
			},
		})
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 && !undoFlag {
			return errors.New("Too few arguments")
		}
		return errortools.FirstError(
			validateMaxArgsFunc(args, 2),
			gittools.ValidateIsRepository,
			validateIsConfigured,
		)
	},
}

func getRenameBranchConfig(args []string) (result renameBranchConfig) {
	if len(args) == 1 {
		result.OldBranchName = gitlib.GetCurrentBranchName()
		result.NewBranchName = args[0]
	} else {
		result.OldBranchName = args[0]
		result.NewBranchName = args[1]
	}
	workflows.EnsureIsNotMainBranch(result.OldBranchName, "The main branch cannot be renamed.")
	if !forceFlag {
		workflows.EnsureIsNotPerennialBranch(result.OldBranchName, fmt.Sprintf("'%s' is a perennial branch. Renaming a perennial branch typically requires other updates. If you are sure you want to do this, use '--force'.", result.OldBranchName))
	}
	if result.OldBranchName == result.NewBranchName {
		flows.ExitWithErrorMessage("Cannot rename branch to current name.")
	}
	if !gittools.IsOffline() {
		scriptlib.Fetch()
	}
	workflows.EnsureHasBranch(result.OldBranchName)
	workflows.EnsureBranchInSync(result.OldBranchName, "Please sync the branches before renaming.")
	workflows.EnsureDoesNotHaveBranch(result.NewBranchName)
	return
}

func getRenameBranchStepList(config renameBranchConfig) (result steps.StepList) {
	result.Append(&steps.CreateBranchStep{BranchName: config.NewBranchName, StartingPoint: config.OldBranchName})
	if gitlib.GetCurrentBranchName() == config.OldBranchName {
		result.Append(&steps.CheckoutBranchStep{BranchName: config.NewBranchName})
	}
	if workflows.IsPerennialBranch(config.OldBranchName) {
		result.Append(&steps.RemoveFromPerennialBranches{BranchName: config.OldBranchName})
		result.Append(&steps.AddToPerennialBranches{BranchName: config.NewBranchName})
	} else {
		result.Append(&steps.DeleteParentBranchStep{BranchName: config.OldBranchName})
		result.Append(&steps.SetParentBranchStep{BranchName: config.NewBranchName, ParentBranchName: gittools.GetParentBranch(config.OldBranchName)})
	}
	for _, child := range gittools.GetChildBranches(config.OldBranchName) {
		result.Append(&steps.SetParentBranchStep{BranchName: child, ParentBranchName: config.NewBranchName})
	}
	result.Append(&steps.DeleteAncestorBranchesStep{})
	if gittools.HasTrackingBranch(config.OldBranchName) && !gittools.IsOffline() {
		result.Append(&steps.CreateTrackingBranchStep{BranchName: config.NewBranchName})
		result.Append(&steps.DeleteRemoteBranchStep{BranchName: config.OldBranchName, IsTracking: true})
	}
	result.Append(&steps.DeleteLocalBranchStep{BranchName: config.OldBranchName})
	result.Wrap(steps.WrapOptions{RunInGitRoot: false, StashOpenChanges: false})
	return
}

func init() {
	renameBranchCommand.Flags().BoolVar(&undoFlag, "undo", false, undoFlagDescription)
	renameBranchCommand.Flags().BoolVar(&forceFlag, "force", false, "Force rename of perennial branch")
	RootCmd.AddCommand(renameBranchCommand)
}
