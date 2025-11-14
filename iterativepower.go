package piscine

func IterativePower(nb int, power int) int {
	// Constraint: Negative powers must return 0.
	if power < 0 {
		return 0
	}

	// Base case: Any number raised to the power of 0 is 1.
	if power == 0 {
		return 1
	}

	// Iterative calculation
	result := nb

	for i := 1; i < power; i++ {
		result *= nb
	}

	return result
}
