package asciiart

import (
	"fmt"
	"os"
)

func FileHandler() []byte {
	// Read banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error")
		return []byte{}
	}


	return data
}