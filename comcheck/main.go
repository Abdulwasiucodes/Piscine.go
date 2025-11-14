package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	targetStrings := []string{"01", "galaxy", "galaxy 01"}

	foundMatch := false

	for _, arg := range args {
		for _, target := range targetStrings {
			if arg == target {
				foundMatch = true
				break
			}
		}
		if foundMatch {
			break
		}
	}

	if foundMatch {
		fmt.Println("Alert!!!")
	}
}
