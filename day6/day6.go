package day6

import (
	"advent-of-code-2022/utils"
	"log"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day6", "day-6-input.txt")
	solvePart1(input[0])
	solvePart2(input[0])
}

func solvePart1(msg string) int {
	start := time.Now().UnixMilli()
	ans := firstUniqueRange(msg, 4)
	end := time.Now().UnixMilli()
	log.Printf("Day 6, Part 1 (%dms): Packet Marker = %d", end-start, ans)
	return ans
}

func solvePart2(msg string) int {
	start := time.Now().UnixMilli()
	ans := firstUniqueRange(msg, 14)
	end := time.Now().UnixMilli()
	log.Printf("Day 6, Part 2 (%dms): Message Start = %d", end-start, ans)
	return ans
}

func firstUniqueRange(msg string, rangeSize int) int {
	l := len(msg)
	for i := rangeSize; i <= l; i++ {
		if !dupCharacters(msg[i-rangeSize : i]) {
			return i
		}
	}
	return -1
}

func dupCharacters(s string) bool {
	cMap := make(map[rune]int)
	for _, c := range s {
		if _, found := cMap[c]; found {
			return true
		}
		cMap[c]++
	}
	return false
}
