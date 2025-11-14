package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	name := ""

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			name = path[i+1:]
			break
		}
	}
	if name == "" {
		name = path
	}

	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
