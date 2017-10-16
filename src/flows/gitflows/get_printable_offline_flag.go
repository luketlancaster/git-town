package gitflows

import (
	"strconv"

	"github.com/Originate/git-town/src/lib/gitlib"
)

// GetPrintableOfflineFlag returns a user printable offline flag
func GetPrintableOfflineFlag() string {
	return strconv.FormatBool(gitlib.IsOffline())
}
