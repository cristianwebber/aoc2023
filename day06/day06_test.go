package day06

import (
	"strings"
	"testing"

	"github.com/cristianwebber/aoc2023/util"
	"github.com/stretchr/testify/require"
)

func TestDay06Part1(t *testing.T) {
	input_text := `
Time:      7  15   30
Distance:  9  40  200
`
	expected := 288

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part1(input)
	require.Equal(t, expected, result)
}

func TestDay06Part2(t *testing.T) {
	input_text := `
Time:      7  15   30
Distance:  9  40  200
`
	expected := 71503

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part2(input)
	require.Equal(t, expected, result)
}
