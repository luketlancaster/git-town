package gitflows

import (
	"strconv"

	"github.com/Originate/git-town/src/lib/gitlib"
)

// GetPrintableHackPushFlag returns a user printable hack push flag
func GetPrintableHackPushFlag() string {
	return strconv.FormatBool(gitlib.ShouldHackPush())
}
