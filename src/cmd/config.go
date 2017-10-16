package cmd

import (
	"fmt"

	"github.com/Originate/git-town/src/flows/promptflows"
	"github.com/Originate/git-town/src/lib/promptlib"
	"github.com/spf13/cobra"
)

var resetFlag bool
var setupFlag bool

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Displays or resets your Git Town configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if resetFlag {
			resetConfig()
		} else if setupFlag {
			setupConfig()
		} else {
			printConfig()
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errortools.FirstError(
			validateMaxArgsFunc(args, 0),
			gittools.ValidateIsRepository,
		)
	},
}

func printConfig() {
	fmt.Println()
	promptlib.PrintLabelAndValue("Main branch", gittools.GetPrintableMainBranch())
	promptlib.PrintLabelAndValue("Perennial branches", gittools.GetPrintablePerennialBranches())

	mainBranch := gittools.GetMainBranch()
	if mainBranch != "" {
		promptlib.PrintLabelAndValue("Branch Ancestry", gittools.GetPrintableBranchTree(mainBranch))
	}

	promptlib.PrintLabelAndValue("Pull branch strategy", gittools.GetPullBranchStrategy())
	promptlib.PrintLabelAndValue("git-hack push flag", gittools.GetPrintableHackPushFlag())
}

func resetConfig() {
	gittools.RemoveAllConfiguration()
}

func setupConfig() {
	promptflows.ConfigureMainBranch()
	promptflows.ConfigurePerennialBranches()
}

func init() {
	configCommand.Flags().BoolVar(&resetFlag, "reset", false, "Remove all Git Town configuration from the current repository")
	configCommand.Flags().BoolVar(&setupFlag, "setup", false, "Run the Git Town configuration wizard")
	RootCmd.AddCommand(configCommand)
}
