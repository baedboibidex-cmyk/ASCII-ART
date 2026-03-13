package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Check if exactly one argument is passed
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	// Read banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}

	// Split banner into lines
	lines := strings.Split(string(data), "\n")

	// Split input by newline
	words := strings.Split(input, "\\n")

	// Process each word
	for _, word := range words {

		// If the line is empty, just print a newline
		if word == "" {
			fmt.Println()
			continue
		}

		// Each ASCII character is 8 lines tall
		for row := 0; row < 8; row++ {

			// Loop through each character
			for i := 0; i < len(word); i++ {

				char := word[i]

				// Check ASCII range
				if char < 32 || char > 126 {
					continue
				}

				// Find character position in banner
				index := (int(char) - 32) * 9

				fmt.Print(lines[index+row])
			}

			fmt.Println()
		}
	}
}