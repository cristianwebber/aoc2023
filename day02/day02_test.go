package day02

import (
	"strings"
	"testing"

	"github.com/cristianwebber/aoc2023/util"
	"github.com/stretchr/testify/require"
)

func TestDay01Part1(t *testing.T) {
	input_text := `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`
	expected := 8

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Day02Part1(input)
	require.Equal(t, expected, result)
}

func TestDay02Part2(t *testing.T) {
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
	result := Day02Part2(input)
	require.Equal(t, expected, result)
}
