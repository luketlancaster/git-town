package gittools

import (
	"errors"

	"github.com/Originate/git-town/src/tools/command"
)

// ValidateIsRepository asserts that the current directory is in a repository
func ValidateIsRepository() error {
	if IsRepository() {
		return nil
	}
	return errors.New("This is not a Git repository")
}

// IsRepository returns whether or not the current directory is in a repository
func IsRepository() bool {
	return command.New("git", "rev-parse").Err() == nil
}
