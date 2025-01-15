package utils

import (
	"fmt"
	"os"
)

func ReadFile(file string) (string, error) {
	jsCode, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
	}
	return string(jsCode), err
}
