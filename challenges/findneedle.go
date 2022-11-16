package challenges

// -1, means index of the needle not found
func FindNeedle(strs []string, s string) int {
	for i, v := range strs {
		if v == s {
			return i
		}
	}

	return -1
}
