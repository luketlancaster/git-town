package flows

import (
	"os"

	"github.com/Originate/git-town/src/tools/prompttools"
)

// ExitWithErrorMessage prints the given error message and terminates the application.
func ExitWithErrorMessage(messages ...string) {
	prompttools.PrintError(messages...)
	os.Exit(1)
}
