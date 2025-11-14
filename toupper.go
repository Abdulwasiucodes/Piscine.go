package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'a' && c <= 'z' {
			runes[i] = c - 32
		}
	}
	return string(runes)
}
