package piscine

func NRune(s string, n int) rune {
	if n < 1 {
		return 0
	}
	runes := []rune(s)
	if n > len(runes) {
		return 0
	}
	return runes[n-1]
}
