package day03

import (
	"fmt"
	"strconv"
)

// Transform a lines of text in a matrix
// Consider that all lines have the same length
func createMatrix(lines []string) [][]rune {
	// Create empty matrix
	matrix := make([][]rune, len(lines))
	for i := range matrix {
		matrix[i] = make([]rune, len(lines[0]))
	}

	// Add values
	for y, line := range lines {
		for x, rune := range line {
			matrix[y][x] = rune
		}
	}
	return matrix
}

func checkAdjacent(m [][]rune, row int, col int) bool {
	positions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, pos := range positions {
		newRow, newCol := row+pos[0], col+pos[1]

		// check for out of bounds
		if newRow >= 0 && newRow < len(m) && newCol >= 0 && newCol < len(m[0]) {
			r := m[newRow][newCol]
			// this is wrong
			if is_symbol(r) {
				return true
			}
		}
	}
	return false
}

type Gear struct {
	x int
	y int
}

var nilGear = Gear{}

func checkGearAdjacent(m [][]rune, row int, col int) Gear {
	positions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, pos := range positions {
		newRow, newCol := row+pos[0], col+pos[1]

		// check for out of bounds
		if newRow >= 0 && newRow < len(m) && newCol >= 0 && newCol < len(m[0]) {
			r := m[newRow][newCol]
			// this is wrong
			if r == '*' {
				return Gear{newRow, newCol}
			}
		}
	}
	return nilGear
}

func contains(list []Gear, target Gear) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func is_symbol(r rune) bool {
	symbolList := []rune{'*', '#', '$', '+', '-', '@', '&', '/', '=', '%'}
	for _, v := range symbolList {
		if r == v {
			return true
		}
	}
	return false
}

func Part1(lines []string) int {
	m := createMatrix(lines)
	total := 0
	isValid := false
	runeList := []rune{}

	for row := range m {
		for col := range m[0] {
			r := m[row][col]
			if r >= '0' && r <= '9' {
				runeList = append(runeList, r)
				isAdjacent := checkAdjacent(m, row, col)
				if !isValid && isAdjacent {
					isValid = true
				}
			} else {
				if isValid {
					intNumber, err := strconv.Atoi(string(runeList))
					if err != nil {
						fmt.Println("Error converting string to int:", intNumber)
						panic(err)
					}
					// fmt.Println("intNumber:", intNumber)
					total += intNumber
				}
				runeList = nil
				isValid = false
			}
		}
	}
	return total
}

func Part2(lines []string) int {
	m := createMatrix(lines)
	total := 0
	gearsMap := make(map[Gear][]int)
	runeList := []rune{}
	gearList := []Gear{}

	for row := range m {
		for col := range m[0] {
			r := m[row][col]
			if r >= '0' && r <= '9' {
				runeList = append(runeList, r)
				gear := checkGearAdjacent(m, row, col)
				if gear != nilGear && !contains(gearList, gear) {
					gearList = append(gearList, gear)
				}
			} else {
				// It's not a number, add number to map
				if len(gearList) > 0 {
					intNumber, err := strconv.Atoi(string(runeList))
					if err != nil {
						fmt.Println("Error converting string to int:", intNumber)
						panic(err)
					}
					for _, gear := range gearList {
						gearsMap[gear] = append(gearsMap[gear], intNumber)
					}

				}
				runeList = nil
				gearList = nil
			}
		}
	}
	// we can have same number, so deduplication don't work
	for _, values := range gearsMap {
		if len(values) == 2 {
			total += (values[0] * values[1])
		}

	}
	return total
}
