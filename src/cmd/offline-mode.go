package cmd

import (
	"github.com/Originate/git-town/src/flows"
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/spf13/cobra"
)

var offlineCommand = &cobra.Command{
	Use:   "offline [(true | false)]",
	Short: "Displays or sets offline mode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printOfflineFlag()
		} else {
			setOfflineFlag(workflows.StringToBool(args[0]))
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			err := validateBooleanArgument(args[0])
			if err != nil {
				return err
			}
		}
		return validateMaxArgs(args, 1)
	},
}

func printOfflineFlag() {
	cfmt.Println(gittools.GetPrintableOfflineFlag())
}

func setOfflineFlag(value bool) {
	gittools.UpdateOffline(value)
}

func init() {
	RootCmd.AddCommand(offlineCommand)
}
