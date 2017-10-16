package prompttools

import (
	"fmt"

	"github.com/Originate/git-town/src/exit"
	"github.com/fatih/color"
)

// PrintError prints the given error message to the console.
func PrintError(messages ...string) {
	errHeaderFmt := color.New(color.Bold).Add(color.FgRed)
	errMessageFmt := color.New(color.FgRed)
	fmt.Println()
	_, err := errHeaderFmt.Println("  Error")
	exit.On(err)
	for _, message := range messages {
		_, err = errMessageFmt.Println("  " + message)
		exit.On(err)
	}
	fmt.Println()
}
