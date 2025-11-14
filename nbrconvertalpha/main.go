package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n, ok := toInt(arg)
		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			if upper {
				z01.PrintRune(rune(64 + n)) // 'A' + n - 1
			} else {
				z01.PrintRune(rune(96 + n)) // 'a' + n - 1
			}
		}
	}
	z01.PrintRune('\n')
}

func toInt(s string) (int, bool) {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		n = n*10 + int(r-'0')
	}
	return n, true
}
