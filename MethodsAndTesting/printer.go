package asciiart

import (
	"fmt"
	"strings"
)

func FormatPrinter(input string) string {
	// Split banner into lines
	lines := strings.Split(string(FileHandler()), "\n")

	// Split input by newline
	words := strings.Split(input, "\\n")
	var new strings.Builder
	// Process each word
	for _, word := range words {
		// If the line is empty, just print a newline
		if word == "" {
			fmt.Println()
			continue
		}

		// Each ASCII character is 8 lines tall
		for row := 0; row < 8; row++ {
			// if row < 0 {
			// 	continue
			// }
			var newest strings.Builder

			// Loop through each character
			for i := 0; i < len(word); i++ {

				char := word[i]

				// Check ASCII range
				if char < 32 || char > 126 {
					continue
				}

				// Find character position in banner
				index := (int(char) - 32) * 9 + 1
				newest.WriteString(lines[index+row])
			}
			new.WriteString(newest.String() + "\n")
		}
	}
	return new.String()
}
