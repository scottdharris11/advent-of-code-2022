package day25

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day25", "day-25-input.txt")
	solvePart1(input)
}

func solvePart1(lines []string) string {
	start := time.Now().UnixMilli()
	fuel := 0
	for _, line := range lines {
		fuel += stoi(line)
	}
	ans := itos(fuel)
	end := time.Now().UnixMilli()
	log.Printf("Day 25, Part 1 (%dms): SNAFU = %s", end-start, ans)
	return ans
}

var SnafuMultipliers = map[rune]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
var SnafuRepresentation = map[int]rune{2: '2', 1: '1', 0: '0', -1: '-', -2: '='}

func stoi(snafu string) int {
	i := 0
	base := 1
	for p := len(snafu) - 1; p >= 0; p-- {
		i += base * SnafuMultipliers[rune(snafu[p])]
		base *= 5
	}
	return i
}

func itos(i int) string {
	// establish the number of places required to represent
	var places []int
	base := 1
	maxValue := 0
	for {
		places = append(places, 2)
		maxValue += base * 2
		if maxValue >= i {
			break
		}
		base *= 5
	}

	// adjust the place values one by one to get to value
	for idx := 0; idx < len(places); idx++ {
		if maxValue == i {
			break
		}
		for a := 1; a <= 4; a++ {
			adjust := maxValue - base
			if adjust < i {
				break
			}
			maxValue = adjust
			places[idx]--
		}
		base /= 5
	}

	// format the output
	snafu := ""
	for idx := 0; idx < len(places); idx++ {
		snafu += string(SnafuRepresentation[places[idx]])
	}
	return snafu
}
