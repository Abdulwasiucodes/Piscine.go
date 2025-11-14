package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var result []string

	for _, t := range s {
		if t == ' ' || t == '\t' || t == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(t)
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return result
}
