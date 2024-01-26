package day07

import (
	"strings"
	"testing"

	"github.com/cristianwebber/aoc2023/util"
	"github.com/stretchr/testify/require"
)

func TestDay07Part1(t *testing.T) {
	input_text := `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`
	expected := 6440

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part1(input)
	require.Equal(t, expected, result)
}

func TestDay07Part2(t *testing.T) {
	input_text := `
  xxx
`
	expected := 100

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part2(input)
	require.Equal(t, expected, result)
}
