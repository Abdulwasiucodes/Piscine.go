package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0] // start with the first element
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i] // add separator before each next element
	}
	return result
}
