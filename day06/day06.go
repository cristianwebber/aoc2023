package day06

import (
	"fmt"
	"github.com/cristianwebber/aoc2023/util"
	"strings"
)

func Part1(lines []string) int {
	fmt.Println("hi")
	strTime := strings.Split(lines[0], ":")[1]
	time, _ := util.StringToNumberList(strTime)
	strDistance := strings.Split(lines[1], ":")[1]
	distance, _ := util.StringToNumberList(strDistance)
	total := 1
	for i := 0; i < len(time); i++ {
		run := 0
		fmt.Println("Time: ", time[i], "; Distance: ", distance[i])
		for holdTime := 1; holdTime < time[i]; holdTime++ {
			x := holdTime * (time[i] - holdTime)
			if x > distance[i] {
				run++
			}
		}
		fmt.Println("Total from the run: ", run)
		total = total * run
	}
	return total
}
func Part2(lines []string) int {
	return 1
}
