package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || hasFlag(args, "--help", "-h") {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	mainStr := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]

		// handle --insert=value or -i=value
		if startsWith(arg, "--insert=") {
			insertStr = arg[len("--insert="):]
			continue
		}
		if startsWith(arg, "-i=") {
			insertStr = arg[len("-i="):]
			continue
		}

		// handle --insert value or -i value
		if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
			continue
		}

		// handle --order or -o
		if arg == "--order" || arg == "-o" {
			order = true
			continue
		}

		// main argument string
		mainStr += arg
	}

	result := mainStr + insertStr
	if order {
		result = orderString(result)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Check if flag is present
func hasFlag(args []string, long, short string) bool {
	for _, a := range args {
		if a == long || a == short {
			return true
		}
	}
	return false
}

// Check prefix manually
func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := range prefix {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

// Manual bubble sort in ASCII order
func orderString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes)-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

// Print help exactly as specified
func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
