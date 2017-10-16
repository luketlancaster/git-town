package promptflows

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Originate/git-town/src/exit"
	"github.com/Originate/git-town/src/lib/promptlib"
	"github.com/Originate/git-town/src/tools/cfmt"
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/Originate/git-town/src/tools/prompttools"
	"github.com/fatih/color"
)

// BranchPromptConfig is
type BranchPromptConfig struct {
	branchNames       []string
	defaultBranchName string
	prompt            string
	validate          func(branchName string) error
}

// AskForBranch ...
func AskForBranch(config BranchPromptConfig) string {
	for {
		cfmt.Print(config.prompt)
		branchName, err := parseBranch(config, promptlib.GetUserInput())
		if err == nil {
			err = config.validate(branchName)
			if err == nil {
				return branchName
			}
		}
		prompttools.PrintError(err.Error())
	}
}

func parseBranch(config BranchPromptConfig, userInput string) (string, error) {
	numericRegex, err := regexp.Compile("^[0-9]+$")
	exit.OnWrap(err, "Error compiling numeric regular expression")

	if numericRegex.MatchString(userInput) {
		return parseBranchNumber(config.branchNames, userInput)
	}
	if userInput == "" {
		return config.defaultBranchName, nil
	}
	if gittools.HasBranch(userInput) {
		return userInput, nil
	}

	return "", fmt.Errorf("Branch '%s' doesn't exist", userInput)
}

// ConfigurePerennialBranches has the user to confgure the perennial branches
func ConfigurePerennialBranches() {
	printConfigurationHeader()
	var newPerennialBranches []string
	for {
		newPerennialBranch := AskForBranch(BranchPromptConfig{
			branchNames:       gittools.GetLocalBranches(),
			defaultBranchName: "",
			prompt:            getPerennialBranchesPrompt(),
			validate: func(branchName string) error {
				if branchName == gittools.GetMainBranch() {
					return fmt.Errorf("'%s' is already set as the main branch", branchName)
				}
				return nil
			},
		})
		if newPerennialBranch == "" {
			break
		}
		newPerennialBranches = append(newPerennialBranches, newPerennialBranch)
	}
	gittools.SetPerennialBranches(newPerennialBranches)
}

// ConfigureMainBranch has the user to confgure the main branch
func ConfigureMainBranch() {
	printConfigurationHeader()
	newMainBranch := AskForBranch(BranchPromptConfig{
		branchNames:       gittools.GetLocalBranches(),
		defaultBranchName: "",
		prompt:            getMainBranchPrompt(),
		validate: func(branchName string) error {
			if branchName == "" {
				return errors.New("A main development branch is required to enable the features provided by Git Town")
			}
			return nil
		},
	})
	gittools.SetMainBranch(newMainBranch)
}

// Helpers

var configurationHeaderShown bool

func getMainBranchPrompt() (result string) {
	result += "Please specify the main development branch by name or number"
	currentMainBranch := gittools.GetMainBranch()
	if currentMainBranch != "" {
		coloredBranchName := color.New(color.Bold).Add(color.FgCyan).Sprintf(currentMainBranch)
		result += fmt.Sprintf(" (current value: %s)", coloredBranchName)
	}
	result += ": "
	return
}

func getPerennialBranchesPrompt() (result string) {
	result += "Please specify a perennial branch by name or number. Leave it blank to finish"
	currentPerennialBranches := gittools.GetPerennialBranches()
	if len(currentPerennialBranches) > 0 {
		coloredBranchNames := color.New(color.Bold).Add(color.FgCyan).Sprintf(strings.Join(currentPerennialBranches, ", "))
		result += fmt.Sprintf(" (current value: %s)", coloredBranchNames)
	}
	result += ": "
	return
}

func printConfigurationHeader() {
	if !configurationHeaderShown {
		configurationHeaderShown = true
		fmt.Println("Git Town needs to be configured")
		fmt.Println()
		promptlib.PrintNumberedBranches(gittools.GetLocalBranches())
		fmt.Println()
	}
}

// HELPERS

func parseBranchNumber(branchNames []string, userInput string) (string, error) {
	index, err := strconv.Atoi(userInput)
	exit.OnWrap(err, "Error parsing string to integer")
	if index >= 1 && index <= len(branchNames) {
		return branchNames[index-1], nil
	}

	return "", errors.New("Invalid branch number")
}
