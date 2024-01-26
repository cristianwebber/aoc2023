package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func decideHandType(cards string) string {
	hand := make(map[rune]int)
	for _, c := range cards {
		hand[c] += 1
	}

	counts := []int{}
	for _, count := range hand {
		counts = append(counts, count)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if counts[0] == 5 {
		return "Five of a kind"
	}
	if counts[0] == 4 {
		return "Four of a kind"
	}
	if counts[0] == 3 && counts[1] == 2 {
		return "Full house"
	}
	if counts[0] == 3 {
		return "Three of a kind"
	}
	if counts[0] == 2 && counts[1] == 2 {
		return "Two pair"
	}
	if counts[0] == 2 {
		return "One pair"
	}
	return "High Card"
}

type ByCards []Hand

// Implements sort.Interface
func (s ByCards) Len() int      { return len(s) }
func (s ByCards) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByCards) Less(i, j int) bool {
	cardsMap := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	// transform to list of runes
	iRunes := []rune(s[i].cards)
	jRunes := []rune(s[j].cards)

	for c := range s[i].cards {
		iInt := cardsMap[iRunes[c]]
		jInt := cardsMap[jRunes[c]]
		if iInt < jInt {
			return true
		} else if iInt > jInt {
			return false
		} else {
			continue
		}
	}
	panic("Same game")
}

func Part1(lines []string) int {
	points := 0
	gameTypes := make(map[string][]Hand)
	fmt.Println("hi")
	lowToHigh := []string{"High Card", "One pair", "Two pair", "Three of a kind", "Full house", "Four of a kind", "Five of a kind"}

	for _, line := range lines {
		pairs := strings.Split(line, " ")
		cards := pairs[0]
		bid := pairs[1]
		bidInt, _ := strconv.Atoi(bid)

		handType := decideHandType(cards)
		gameTypes[handType] = append(gameTypes[handType], Hand{cards, bidInt})

	}
	for _, gameType := range lowToHigh {
		games := gameTypes[gameType]
		sort.Sort(ByCards(games))
	}

	idx := 1
	for _, gameType := range lowToHigh {
		for _, game := range gameTypes[gameType] {
			fmt.Println(idx, game)

			points += idx * game.bid
			idx++
		}
	}
	return points
}
func Part2(lines []string) int {
	return 1
}
