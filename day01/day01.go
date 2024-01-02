package day01

import (
	"strings"
	"unicode"
)

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func containsKey(input string, keys map[string]int) (int, bool) {
	for key, value := range keys {
		if strings.Contains(input, key) {
			return value, true
		}
	}
	return 0, false
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Day01Part1(lines []string) int {
	var total int = 0
	for _, line := range lines {

		var runes []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				runes = append(runes, char)
			}
		}

		if len(runes) > 0 {
			line_value := int(runes[0]-'0')*10 + int(runes[len(runes)-1]-'0')
			total += line_value
		}
	}
	return total
}

func Day01Part2(lines []string) int {
	var total int = 0
	for _, line := range lines {
		var number_str string

		// find first number
		for _, char := range line {

			if unicode.IsDigit(char) {
				total += 10 * int(char-'0')
				break
			}

			if unicode.IsLetter(char) {
				number_str += string(char)
				value, found := containsKey(number_str, wordToDigit)
				if found {
					total += 10 * value
					break
				}
			}

		}

		// find second number
		var number_str2 string
		for _, char := range reverseString(line) {

			if unicode.IsDigit(char) {
				total += int(char - '0')
				break
			}

			if unicode.IsLetter(char) {
				number_str2 = string(char) + number_str2
				value, found := containsKey(number_str2, wordToDigit)
				if found {
					total += value
					break
				}
			}

		}
	}
	return total
}
