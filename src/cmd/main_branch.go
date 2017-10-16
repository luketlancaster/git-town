package cmd

import (
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/spf13/cobra"
)

var mainBranchCommand = &cobra.Command{
	Use:   "main-branch [<branch>]",
	Short: "Displays or sets your main development branch",
	Long: `Displays or sets your main development branch

The main branch is the Git branch from which new feature branches are cut.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printMainBranch()
		} else {
			setMainBranch(args[0])
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errortools.FirstError(
			validateMaxArgsFunc(args, 1),
			gittools.ValidateIsRepository,
		)
	},
}

func printMainBranch() {
	cfmt.Println(gittools.GetPrintableMainBranch())
}

func setMainBranch(branchName string) {
	workflows.EnsureHasBranch(branchName)
	gittools.SetMainBranch(branchName)
}

func init() {
	RootCmd.AddCommand(mainBranchCommand)
}
