package piscine

// ConvertBase converts a number from baseFrom to baseTo
func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Helper: find index of a rune in a string
	indexRune := func(s string, r rune) int {
		for i, c := range s {
			if c == r {
				return i
			}
		}
		return -1 // should never happen with valid input
	}

	// Step 1: Convert nbr from baseFrom to decimal
	decimalValue := 0
	baseLenFrom := len(baseFrom)
	for _, c := range nbr {
		decimalValue = decimalValue*baseLenFrom + indexRune(baseFrom, c)
	}

	// Step 2: Convert decimal to baseTo
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	baseLenTo := len(baseTo)
	result := ""
	for decimalValue > 0 {
		remainder := decimalValue % baseLenTo
		result = string(baseTo[remainder]) + result
		decimalValue /= baseLenTo
	}

	return result
}
