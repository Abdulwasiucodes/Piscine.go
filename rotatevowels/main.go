package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Collect all vowels (in order)
	vowels := []rune{}
	for _, word := range args {
		for _, r := range word {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}

	// Reverse the vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Replace vowels in each argument
	vowelIndex := 0
	for i, word := range args {
		for _, r := range word {
			if isVowel(r) {
				printRune(vowels[vowelIndex])
				vowelIndex++
			} else {
				printRune(r)
			}
		}
		// Print space between arguments (but not after the last)
		if i < len(args)-1 {
			printRune(' ')
		}
	}
	z01.PrintRune('\n')
}

// Checks if a rune is a vowel (case-insensitive, y excluded)
func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

// Convenience for z01.PrintRune
func printRune(r rune) {
	z01.PrintRune(r)
}
