package stringtools

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
