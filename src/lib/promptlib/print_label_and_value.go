package promptlib

import (
	"fmt"

	"github.com/Originate/git-town/src/exit"
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/Originate/git-town/src/tools/stringtools"
	"github.com/fatih/color"
)

// PrintLabelAndValue prints the label bolded and underlined
// the value indented on the next line
// followed by an empty line
func PrintLabelAndValue(label, value string) {
	labelFmt := color.New(color.Bold).Add(color.Underline)
	_, err := labelFmt.Println(label + ":")
	exit.On(err)
	cfmt.Println(stringtools.Indent(value, 1))
	fmt.Println()
}
