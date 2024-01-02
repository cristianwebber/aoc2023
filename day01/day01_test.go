package day01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay01Part1(t *testing.T) {
	input_text := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	expected := 142

	var input []string = strings.Split(input_text, "\n")
	result := Day01Part1(input)
	require.Equal(t, expected, result)
}

func TestDay01Part2(t *testing.T) {
	input_text := `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
	expected := 281

	var input []string = strings.Split(input_text, "\n")
	result := Day01Part2(input)
	require.Equal(t, expected, result)
}
