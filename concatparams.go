package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := ""
	for x, str := range args {
		result += str
		if x < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
