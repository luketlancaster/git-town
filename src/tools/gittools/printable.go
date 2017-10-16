package gittools

import "strings"

// GetPrintableMainBranch returns a user printable main branch
func GetPrintableMainBranch() string {
	output := GetMainBranch()
	if output == "" {
		return "[none]"
	}
	return output
}

// GetPrintablePerennialBranches returns a user printable list of perennial branches
func GetPrintablePerennialBranches() string {
	output := strings.Join(GetPerennialBranches(), "\n")
	if output == "" {
		return "[none]"
	}
	return output
}
