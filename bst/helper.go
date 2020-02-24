package bst

func greatest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// pInt is a helper for getting a pointer to an int from an int.
func pInt(val int) *int {
	return &val
}

func withSpaces(n int) string {
	str := ""
	for idx := 0; idx < n; idx++ {
		str += " "
	}
	return str
}
