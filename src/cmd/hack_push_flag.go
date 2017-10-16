package cmd

import (
	"github.com/Originate/git-town/src/flows"
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/spf13/cobra"
)

var hackPushFlagCommand = &cobra.Command{
	Use:   "hack-push-flag [(true | false)]",
	Short: "Displays or sets your hack push flag",
	Long: `Displays or sets your hack push flag

Newly hacked branches will be pushed upon creation
if and only if "hack-push-flag" is true.
The default value is false.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printHackPushFlag()
		} else {
			setHackPushFlag(workflows.StringToBool(args[0]))
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			err := validateBooleanArgument(args[0])
			if err != nil {
				return err
			}
		}
		return errortools.FirstError(
			validateMaxArgsFunc(args, 1),
			gittools.ValidateIsRepository,
		)
	},
}

func printHackPushFlag() {
	cfmt.Println(gittools.GetPrintableHackPushFlag())
}

func setHackPushFlag(value bool) {
	gittools.UpdateShouldHackPush(value)
}

func init() {
	RootCmd.AddCommand(hackPushFlagCommand)
}
