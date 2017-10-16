package scriptflows

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Originate/git-town/src/browsers"
	"github.com/Originate/git-town/src/exit"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/tools/dryrun"
	"github.com/Originate/git-town/src/tools/gittools"

	"github.com/fatih/color"
)

var dryRunMessage = `
In dry run mode. No commands will be run. When run in normal mode, the command
output will appear beneath the command. Some commands will only be run if
necessary. For example: 'git push' will run if and only if there are local
commits not on the remote.
`

// OpenBrowser opens the default browser with the given URL.
func OpenBrowser(url string) {
	command := browsers.GetOpenBrowserCommand()
	err := RunCommand(command, url)
	exit.On(err)
}

// PrintCommand prints the given command-line operation on the console.
func PrintCommand(cmd ...string) {
	header := ""
	for index, part := range cmd {
		if strings.Contains(part, " ") {
			part = "\"" + strings.Replace(part, "\"", "\\\"", -1) + "\""
		}
		if index != 0 {
			header = header + " "
		}
		header = header + part
	}
	if strings.HasPrefix(header, "git") && gittools.IsRepository() {
		header = fmt.Sprintf("[%s] %s", gitlib.GetCurrentBranchName(), header)
	}
	fmt.Println()
	_, err := color.New(color.Bold).Println(header)
	exit.On(err)
}

// RunCommand executes the given command-line operation.
func RunCommand(cmd ...string) error {
	PrintCommand(cmd...)
	if dryrun.IsActive() {
		if len(cmd) == 3 && cmd[0] == "git" && cmd[1] == "checkout" {
			dryrun.SetCurrentBranchName(cmd[2])
		}
		return nil
	}
	// Windows commands run inside CMD
	// because opening browsers is done via "start"
	if runtime.GOOS == "windows" {
		cmd = append([]string{"cmd", "/C"}, cmd...)
	}
	subProcess := exec.Command(cmd[0], cmd[1:]...) // #nosec
	subProcess.Stderr = os.Stderr
	subProcess.Stdin = os.Stdin
	subProcess.Stdout = os.Stdout
	return subProcess.Run()
}

// RunCommandSafe executes the given command-line operation, exiting if the command errors
func RunCommandSafe(cmd ...string) {
	err := RunCommand(cmd...)
	exit.On(err)
}
