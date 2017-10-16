package promptlib

import (
	"bufio"
	"os"
	"strings"

	"github.com/Originate/git-town/src/exit"
)

var inputReader = bufio.NewReader(os.Stdin)

// GetUserInput reads input from the user and returns it.
func GetUserInput() string {
	text, err := inputReader.ReadString('\n')
	exit.OnWrap(err, "Error getting user input")
	return strings.TrimSpace(text)
}
