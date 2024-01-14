package day03

import (
	"strings"
	"testing"

	"github.com/cristianwebber/aoc2023/util"
	"github.com/stretchr/testify/require"
)

func TestDay03Part1(t *testing.T) {
	input_text := `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`
	expected := 4361

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part1(input)
	require.Equal(t, expected, result)
}

func TestDay03Part2(t *testing.T) {
	input_text := `
  xxx
`
	expected := 100

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part2(input)
	require.Equal(t, expected, result)
}
