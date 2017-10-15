package util

import (
	"strconv"
	"strings"

	"github.com/Originate/git-town/src/exit"
)

// DoesStringArrayContain returns whether the given string slice
// contains the given string.
func DoesStringArrayContain(list []string, value string) bool {
	for _, element := range list {
		if element == value {
			return true
		}
	}
	return false
}

// Indent indents the given string with the given level of indentation
// on each line. Each level of indentation is two spaces.
func Indent(message string, level int) string {
	prefix := strings.Repeat("  ", level)
	return prefix + strings.Replace(message, "\n", "\n"+prefix, -1)
}

// Pluralize outputs the count and the word. The word is made plural
// if the count isn't one
func Pluralize(count, word string) string {
	result := count + " " + word
	if count != "1" {
		result = result + "s"
	}
	return result
}

// RemoveStringFromSlice returns a new string slice which is the given string slice
// with the given string removed
func RemoveStringFromSlice(list []string, value string) (result []string) {
	for _, element := range list {
		if element != value {
			result = append(result, element)
		}
	}
	return
}

// StringToBool parses the given string into a bool
func StringToBool(arg string) bool {
	value, err := strconv.ParseBool(arg)
	exit.On(err)
	return value
}
