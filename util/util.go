package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// Transform a string in a list of numbers
func StringToNumberList(input string) ([]int, error) {
	// Split the string into a slice of number strings
	numberStrings := strings.Fields(input)

	// Convert each number string to an integer
	var numbers []int
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
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
