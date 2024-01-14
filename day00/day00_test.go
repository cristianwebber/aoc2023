package day00

import (
	"strings"
	"testing"

	"github.com/cristianwebber/aoc2023/util"
	"github.com/stretchr/testify/require"
)

func TestDay00Part1(t *testing.T) {
	input_text := `
  xxx
`
	expected := 100

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part1(input)
	require.Equal(t, expected, result)
}

func TestDay00Part2(t *testing.T) {
	input_text := `
  xxx
`
	expected := 100

	var input []string = strings.Split(input_text, "\n")
	input = util.RemoveEmptyStrings(input)
	result := Part2(input)
	require.Equal(t, expected, result)
}
