package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 { // handle min int for 64-bit
			printNbrBaseRecursive(-(nbr / len(base)), base)
			z01.PrintRune(rune(base[-(nbr % len(base))]))
			return
		}
		nbr = -nbr
	}

	printNbrBaseRecursive(nbr, base)
}

func printNbrBaseRecursive(nbr int, base string) {
	baseLen := len(base)
	if nbr >= baseLen {
		printNbrBaseRecursive(nbr/baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i, c := range base {
		if c == '+' || c == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if c == rune(base[j]) {
				return false
			}
		}
	}
	return true
}
