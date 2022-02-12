package strutils

// Remove remove target from string list.
func Remove(strings []string, target string) []string {
	result := []string{}
	for _, v := range strings {
		if v != target {
			result = append(result, v)
		}
	}
	return result
}
