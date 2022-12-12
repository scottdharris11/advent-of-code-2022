package day12

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day12", "day-12-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(grid []string) int {
	start := time.Now().UnixMilli()
	ans := len(grid)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 1 (%dms): Shortest Route = %d", end-start, ans)
	return ans
}

func solvePart2(grid []string) int {
	start := time.Now().UnixMilli()
	ans := len(grid)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 2 (%dms): Shortest Route = %d", end-start, ans)
	return ans
}
