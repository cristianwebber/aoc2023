package util

import (
	"bufio"
	"log"
	"os"
)

// Read a file and return a list of lines.
func ReadInput(path string) []string {
	input, err := os.Open(path)
	if err != nil {
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
