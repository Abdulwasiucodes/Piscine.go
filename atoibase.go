package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, c := range s {
		index := indexInBase(c, base)
		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}
	return result
}

func indexInBase(r rune, base string) int {
	for i, c := range base {
		if c == r {
			return i
		}
	}
	return -1
}
