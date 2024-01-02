package main

import (
	"flag"
	"fmt"

	"github.com/cristianwebber/aoc2023/day01"
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
