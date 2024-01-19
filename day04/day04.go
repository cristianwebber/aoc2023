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
	total := 0
	for _, line := range lines {
		lineWinningNumbers := 0
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
	var gameNumber int
	var cardsMap = map[int]int{}
	var err error

	for _, line := range lines {
		lineWinningNumbers := 0
		game := strings.Split(line, ":")
		numStr := strings.Replace(game[0], "Card ", "", -1)
		numStr = strings.ReplaceAll(numStr, " ", "")
		gameNumber, err = strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
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

		// add actual game
		cardsMap[gameNumber] += 1

		// add next games
		for i := gameNumber + 1; i <= gameNumber+lineWinningNumbers; i++ {
			cardsMap[i] += 1 * cardsMap[gameNumber]
		}

	}

	fmt.Println(cardsMap)
	total := 0
	for i := 1; i <= gameNumber; i++ {
		total += cardsMap[i]
	}
	return total
}
