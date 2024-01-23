package main

import (
	"flag"
	"fmt"

	"github.com/cristianwebber/aoc2023/day01"
	"github.com/cristianwebber/aoc2023/day02"
	"github.com/cristianwebber/aoc2023/day03"
	"github.com/cristianwebber/aoc2023/day04"
	"github.com/cristianwebber/aoc2023/day05"
	"github.com/cristianwebber/aoc2023/day06"
	"github.com/cristianwebber/aoc2023/util"
)

var (
	day  int
	part int
)

func init() {
	flag.IntVar(&day, "day", 0, "Specify the day.")
	flag.IntVar(&part, "part", 0, "Specify the part.")
}

var funcMap = map[int]map[int]func([]string) int{
	1: {
		1: day01.Day01Part1,
		2: day01.Day01Part2,
	},
	2: {
		1: day02.Day02Part1,
		2: day02.Day02Part2,
	},
	3: {
		1: day03.Part1,
		2: day03.Part2,
	},
	4: {
		1: day04.Part1,
		2: day04.Part2,
	},
	5: {
		1: day05.Part1,
		2: day05.Part2,
	},
	6: {
		1: day06.Part1,
		2: day06.Part2,
	},
}

func main() {
	flag.Parse()

	if day <= 0 || part <= 0 {
		fmt.Println("Please provide valid values for both 'day' and 'part'.")
		return
	}

	funcForDay, dayExists := funcMap[day]
	if !dayExists {
		fmt.Printf("Day %d not found.\n", day)
		return
	}

	partFunction, partExists := funcForDay[part]
	if !partExists {
		fmt.Printf("Part %d not found for Day %d.\n", part, day)
		return
	}
	fmt.Printf("Running Day %v Part %v.\n", day, part)

	input := util.ReadInput(fmt.Sprintf("day%02d/input.txt", day))

	result := partFunction(input)
	fmt.Printf("Result is:\n%v\n", result)
}
