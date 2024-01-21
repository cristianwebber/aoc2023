package day05

import (
	"fmt"
	"github.com/cristianwebber/aoc2023/util"
	"strings"
)

func Part1(lines []string) int {
	fmt.Println("hi")
	var maps = map[string]interface{}{
		"seedToSoil":            make(map[int]int),
		"soilToFertilizer":      make(map[int]int),
		"fertilizerToWater":     make(map[int]int),
		"waterToLight":          make(map[int]int),
		"lightToTemperature":    make(map[int]int),
		"temperatureToHumidity": make(map[int]int),
		"humidityToLocation":    make(map[int]int),
	}

	strNums := strings.Split(lines[0], ":")[1]
	seeds, _ := util.StringToNumberList(strNums)

	var selectedMap string
	for _, line := range lines[1:] {
		switch line {
		case "":
			continue
		case "seed-to-soil map:":
			selectedMap = "seedToSoil"
			continue
		case "soil-to-fertilizer map:":
			selectedMap = "soilToFertilizer"
			continue
		case "fertilizer-to-water map:":
			selectedMap = "fertilizerToWater"
			continue
		case "water-to-light map:":
			selectedMap = "waterToLight"
			continue
		case "light-to-temperature map:":
			selectedMap = "lightToTemperature"
			continue
		case "temperature-to-humidity map:":
			selectedMap = "temperatureToHumidity"
			continue
		case "humidity-to-location map:":
			selectedMap = "humidityToLocation"
			continue
		}

		coords, _ := util.StringToNumberList(line)
		dest := coords[0]
		source := coords[1]
		range_length := coords[2]
		for i := 0; i < range_length; i++ {
			mapToAdd, _ := maps[selectedMap].(map[int]int)
			mapToAdd[source+i] = dest + i
		}

	}

	// get the final location
	var finalLocations []int
	for _, seed := range seeds {
		soil, ok := maps["seedToSoil"].(map[int]int)[seed]
		if !ok {
			soil = seed
		}
		fertilizer, ok := maps["soilToFertilizer"].(map[int]int)[soil]
		if !ok {
			fertilizer = soil
		}
		water, ok := maps["fertilizerToWater"].(map[int]int)[fertilizer]
		if !ok {
			water = fertilizer
		}
		light, ok := maps["waterToLight"].(map[int]int)[water]
		if !ok {
			light = water
		}
		temperature, ok := maps["lightToTemperature"].(map[int]int)[light]
		if !ok {
			temperature = light
		}
		humidity, ok := maps["temperatureToHumidity"].(map[int]int)[temperature]
		if !ok {
			humidity = temperature
		}
		location, ok := maps["humidityToLocation"].(map[int]int)[humidity]
		if !ok {
			location = humidity
		}

		finalLocations = append(finalLocations, location)
	}

	minLocation := finalLocations[0]
	for _, value := range finalLocations[1:] {
		if value < minLocation {
			minLocation = value
		}
	}

	// for key, value := range maps {
	// 	fmt.Println(key, value)
	// }

	return minLocation
}
func Part2(lines []string) int {
	return 1
}
