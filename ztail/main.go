package main

import (
	"fmt"
	"os"
)

func atoi(s string) (int, bool) {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 || args[0] != "-c" {
		fmt.Println("Usage: go run . -c <num> <file1> [file2 ...]")
		os.Exit(1)
	}
	numBytes, ok := atoi(args[1])
	if !ok || numBytes < 0 {
		fmt.Println("Invalid number of bytes")
		os.Exit(1)
	}
	files := args[2:]
	exitCode := 0
	for i, filename := range files {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", filename)
			exitCode = 1
			continue
		}
		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", filename)
		}
		start := 0
		if len(data) > numBytes {
			start = len(data) - numBytes
		}
		fmt.Print(string(data[start:]))
	}
	os.Exit(exitCode)
}
