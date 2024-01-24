package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cristianwebber/aoc2023/util"
)

func Part1(lines []string) int {
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

// Binary Search
func recursiveBinarySearch(lowHold int, highHold int, timeInt int, distanceInt int) int {
	holdTime := lowHold + ((highHold - lowHold) / 2)
	totalDistance := holdTime * (timeInt - holdTime)
	fmt.Println("hold: ", holdTime, "Distance: ", totalDistance)

	if totalDistance >= distanceInt {
		if distanceInt > (holdTime-1)*(timeInt-(holdTime-1)) {
			return holdTime
		}
	}

	if totalDistance > distanceInt {
		// winning
		return recursiveBinarySearch(lowHold, holdTime, timeInt, distanceInt)
	} else {
		// losing
		return recursiveBinarySearch(holdTime+1, highHold, timeInt, distanceInt)
	}
}

func Part2(lines []string) int {
	strTime := strings.Split(lines[0], ":")[1]
	time := strings.ReplaceAll(strTime, " ", "")
	timeInt, _ := strconv.Atoi(time)
	strDistance := strings.Split(lines[1], ":")[1]
	distance := strings.ReplaceAll(strDistance, " ", "")
	distanceInt, _ := strconv.Atoi(distance)

	lowerBoundary := recursiveBinarySearch(1, timeInt-1, timeInt, distanceInt)

	// We need only to find lower, because it's symetrical
	total := timeInt - (lowerBoundary * 2) + 1

	return total
}
