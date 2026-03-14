package main

import (
	asciiart "ascii-art/MethodsAndTesting"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: enter 2 arguments")
		return
	}

	input := os.Args[1]

	if input == "" {
		return
	}
	fmt.Print(asciiart.FormatPrinter(input))
}