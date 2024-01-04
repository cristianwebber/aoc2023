package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02Part1(lines []string) int {
	total := 0
ThisIsDefinitelyABadPractice:
	for _, line := range lines {
		game := strings.Split(line, ": ")
		game_number := strings.Replace(game[0], "Game ", "", -1)
		rounds := strings.Split(game[1], "; ")

		for _, cubes := range rounds {
			draws := strings.Split(cubes, ", ")
			for _, color := range draws {
				count_and_color := strings.Split(color, " ")
				count := count_and_color[0]
				color := count_and_color[1]

				int_count, err := strconv.Atoi(count)
				if err != nil {
					panic(err)
				}

				//check condition
				if int_count > 12 && color == "red" {
					continue ThisIsDefinitelyABadPractice
				}
				if int_count > 13 && color == "green" {
					continue ThisIsDefinitelyABadPractice
				}
				if int_count > 14 && color == "blue" {
					continue ThisIsDefinitelyABadPractice
				}
			}
		}

		int_game_number, err := strconv.Atoi(game_number)
		if err != nil {
			fmt.Println("Something wrong during game parsing")
			panic(err)
		}
		total += int_game_number
	}
	return total
}
func Day02Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")
		max_game := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, cubes := range rounds {
			draws := strings.Split(cubes, ", ")
			for _, color := range draws {
				count_and_color := strings.Split(color, " ")
				count := count_and_color[0]
				color := count_and_color[1]

				int_count, err := strconv.Atoi(count)
				if err != nil {
					panic(err)
				}

				if int_count > max_game[color] {
					max_game[color] = int_count
				}

			}
		}
		total += max_game["red"] * max_game["blue"] * max_game["green"]

	}
	return total
}
