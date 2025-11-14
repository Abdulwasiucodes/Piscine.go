package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if !inWord {
				// start of word
				if r >= 'a' && r <= 'z' {
					runes[i] -= 32 // capitalize
				}
				inWord = true
			} else {
				// inside word
				if r >= 'A' && r <= 'Z' {
					runes[i] += 32 // lowercase
				}
			}
		} else {
			inWord = false // non-alphanumeric ends a word
		}
	}
	return string(runes)
}
