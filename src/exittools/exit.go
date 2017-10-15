package exittools

import (
	"os"

	"github.com/Originate/git-town/src/prompt"
)

// ExitWithErrorMessage prints the given error message and terminates the application.
func ExitWithErrorMessage(messages ...string) {
	prompt.PrintError(messages...)
	os.Exit(1)
}
