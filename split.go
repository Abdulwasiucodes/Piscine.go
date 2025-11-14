package piscine

func Split(s, sep string) []string {
	var result []string
	word := ""
	sepLen := len(sep)

	for u := 0; u < len(s); {
		if len(s[u:]) >= sepLen && s[u:u+sepLen] == sep {
			result = append(result, word)
			word = ""
			u += sepLen
		} else {
			word += string(s[u])
			u++
		}
	}

	result = append(result, word)
	return result
}
