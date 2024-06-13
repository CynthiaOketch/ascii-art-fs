package main

import (
	"fmt"
	"os"

	"ascii-art/utils"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}

	// Load ASCII characters from file
	var file string
	if len(os.Args) == 2 {
		file = "standard.txt"
	} else if len(os.Args) == 3 {
		file = utils.DetermineFileName(os.Args[2])
	}
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("invalid text file")
		return
	}
	s := utils.ReplaceEscape(os.Args[1])
	// Check for invalid characters
	for _, char := range s {
		if char > 126 || char < 32 {
			fmt.Printf("Error: Character %q is not accepted\n", char)
			os.Exit(0)
		}
	}

	contentLines := utils.SplitFile(string(content))

	if len(contentLines) != 856 {
		fmt.Println("invalid text file")
		return
	}
	utils.DisplayText(os.Args[1], contentLines)
}
