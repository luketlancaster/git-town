package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// FetchUpstreamStep brings the Git history of the local repository
// up to speed with activities that happened in the upstream remote.
type FetchUpstreamStep struct {
	NoOpStep
}

// Run executes this step.
func (step *FetchUpstreamStep) Run() error {
	return scriptflows.RunCommand("git", "fetch", "upstream")
}
