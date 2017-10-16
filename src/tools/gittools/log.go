package gittools

import "github.com/Originate/git-town/src/tools/command"

// GetLastCommitMessage returns the commit message for the last commit
func GetLastCommitMessage() string {
	return command.New("git", "log", "-1", "--format=%B").Output()
}
