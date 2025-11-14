package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'A' && c <= 'Z' {
			runes[i] = c + 32
		}
	}
	return string(runes)
}
