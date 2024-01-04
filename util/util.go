package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Remove empty string from arrays
func RemoveEmptyStrings(strarr []string) []string {
	var arr []string
	for _, str := range strarr {
		if str != "" {
			arr = append(arr, str)
		}
	}
	return arr
}

// Read a file and return a list of lines.
func ReadInput(path string) []string {
	input, err := os.Open(path)
	if err != nil {
		fmt.Println("File not found.")
		log.Fatal(err)
	}
	defer input.Close()

	var data []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
