package asciiart

import (
	"fmt"
	"os"
)

func FileHandler() ([]byte, string) {
	// Read banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error")
		return []byte{}, "couldn't read the file "
	}


	return data, "file read successfully"
}