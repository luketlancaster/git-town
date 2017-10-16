package gitlib

import (
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/Originate/git-town/src/tools/stringtools"
)

// IsOffline returns whether Git Town is currently in offline mode
func IsOffline() bool {
	return stringtools.StringToBool(gittools.GetConfigurationValueWithDefault("git-town.offline", "false"))
}
