package piscine

// isPrime is an optimized helper function to check for primality.
func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	// Optimized check: only check odd divisors up to the square root of nb.
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	// If nb is 0 or negative, the first prime found is 2.
	if nb <= 2 {
		return 2
	}

	// Start checking from nb and iterate upwards.
	for i := nb; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}
