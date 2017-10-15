package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Originate/git-town/src/cfmt"
	"github.com/Originate/git-town/src/exit"
	"github.com/fatih/color"
)

// ExitWithErrorMessage prints the given error message and terminates the application.
func ExitWithErrorMessage(messages ...string) {
	PrintError(messages...)
	os.Exit(1)
}

var inputReader = bufio.NewReader(os.Stdin)

// GetUserInput reads input from the user and returns it.
func GetUserInput() string {
	text, err := inputReader.ReadString('\n')
	exit.OnWrap(err, "Error getting user input")
	return strings.TrimSpace(text)
}

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

// PrintLabelAndValue prints the label bolded and underlined
// the value indented on the next line
// followed by an empty line
func PrintLabelAndValue(label, value string) {
	labelFmt := color.New(color.Bold).Add(color.Underline)
	_, err := labelFmt.Println(label + ":")
	exit.On(err)
	cfmt.Println(Indent(value, 1))
	fmt.Println()
}
