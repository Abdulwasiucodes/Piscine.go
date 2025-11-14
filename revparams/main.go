package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // skip program name

	for a := len(args) - 1; a >= 0; a-- {
		for _, c := range args[a] {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
