package day04

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func stringToNumberList(input string) ([]int, error) {
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

func Part1(lines []string) int {
	fmt.Println("hi")
	total := 0
	for _, line := range lines {
		lineWinningNumbers := 0
		fmt.Println(line)
		game := strings.Split(line, ":")
		// gameNumber := strings.Replace(game[0], "Card ", "", -1)
		numbers := strings.Split(game[1], "|")
		winningNumbers, _ := stringToNumberList(numbers[0])
		elfNumbers, _ := stringToNumberList(numbers[1])

		for _, i := range elfNumbers {
			for _, j := range winningNumbers {
				if i == j {
					lineWinningNumbers += 1
				}
			}
		}

		total += int(math.Pow(2, float64(lineWinningNumbers-1)))
	}
	return total
}
func Part2(lines []string) int {
	return 1
}
