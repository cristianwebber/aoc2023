package day05

import (
	"fmt"
	"github.com/cristianwebber/aoc2023/util"
	"strings"
)

func Part1(lines []string) int {
	fmt.Println("hi")

	strNums := strings.Split(lines[0], ":")[1]
	seeds, _ := util.StringToNumberList(strNums)
	soil := make([]int, len(seeds))
	fertilizer := make([]int, len(seeds))
	water := make([]int, len(seeds))
	light := make([]int, len(seeds))
	temperature := make([]int, len(seeds))
	humidity := make([]int, len(seeds))
	location := make([]int, len(seeds))

	for seedIdx, seed := range seeds {
		var oldMap *[]int
		var selectedMap *[]int
		for _, line := range lines[1:] {
			switch line {
			case "":
				continue
			case "seed-to-soil map:":
				oldMap = &seeds
				selectedMap = &soil
				continue
			case "soil-to-fertilizer map:":
				// fill seed-to-soil
				if soil[seedIdx] == 0 {
					soil[seedIdx] = seed
				}
				oldMap = &soil
				selectedMap = &fertilizer
				continue
			case "fertilizer-to-water map:":
				if fertilizer[seedIdx] == 0 {
					fertilizer[seedIdx] = soil[seedIdx]
				}
				oldMap = &fertilizer
				selectedMap = &water
				continue
			case "water-to-light map:":
				if water[seedIdx] == 0 {
					water[seedIdx] = fertilizer[seedIdx]
				}
				oldMap = &water
				selectedMap = &light
				continue
			case "light-to-temperature map:":
				if light[seedIdx] == 0 {
					light[seedIdx] = water[seedIdx]
				}
				oldMap = &light
				selectedMap = &temperature
				continue
			case "temperature-to-humidity map:":
				if temperature[seedIdx] == 0 {
					temperature[seedIdx] = light[seedIdx]
				}
				oldMap = &temperature
				selectedMap = &humidity
				continue
			case "humidity-to-location map:":
				if humidity[seedIdx] == 0 {
					humidity[seedIdx] = temperature[seedIdx]
				}
				oldMap = &humidity
				selectedMap = &location
				continue
			}
			coords, _ := util.StringToNumberList(line)
			dest := coords[0]
			source := coords[1]
			length := coords[2]

			if (*oldMap)[seedIdx] >= source && (*oldMap)[seedIdx] < source+length {
				final_dest := ((*oldMap)[seedIdx] - source) + dest
				(*selectedMap)[seedIdx] = final_dest
			}
		}
		// add missing location
		if location[seedIdx] == 0 {
			location[seedIdx] = humidity[seedIdx]
		}
	}
	fmt.Println("seed", seeds)
	fmt.Println("soil", soil)
	fmt.Println("fertilizer", fertilizer)
	fmt.Println("water", water)
	fmt.Println("light", light)
	fmt.Println("temperature", temperature)
	fmt.Println("humidity", humidity)
	fmt.Println("location", location)

	minValue := location[0]
	for _, value := range location[1:] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}
func Part2(lines []string) int {
	// Want to do it in wrong way using goroutines. Keeping it for later
	return 1
}
