package asciiart

import (
	"fmt"
	"os"
)

func FileHandler() ([]byte, bool) {
	// Read banner file
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error")
		return []byte{}, false
	}


	return data, true
}