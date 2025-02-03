package day23

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day23", "day-23-input.txt")
	solvePart1(input)
	solvePart2(input)
}

var LOOKS = []Look{
	{checks: [][]int{{0, -1}, {1, -1}, {-1, -1}}, move: []int{0, -1}},
	{checks: [][]int{{0, 1}, {1, 1}, {-1, 1}}, move: []int{0, 1}},
	{checks: [][]int{{-1, 0}, {-1, -1}, {-1, 1}}, move: []int{-1, 0}},
	{checks: [][]int{{1, 0}, {1, -1}, {1, 1}}, move: []int{1, 0}},
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	elves := parseInput(lines)
	li := 0
	for i := 0; i < 10; i++ {
		var looks []Look
		looks = append(looks, LOOKS[li:]...)
		looks = append(looks, LOOKS[:li]...)
		li = (li + 1) % len(LOOKS)
		elves = doRound(elves, looks)
	}
	ans := countEmpty(elves)
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 1 (%dms): Answer = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

type Elf struct {
	x int
	y int
}

type Look struct {
	checks [][]int
	move   []int
}

func parseInput(lines []string) map[Elf]bool {
	elves := make(map[Elf]bool)
	for y, row := range lines {
		for x, col := range row {
			if col == '#' {
				elves[Elf{x, y}] = true
			}
		}
	}
	return elves
}

func doRound(elves map[Elf]bool, looks []Look) map[Elf]bool {
	done := make(map[Elf]bool)
	proposed := make(map[Elf]Elf)
	for elf := range elves {
		inVicinity := false
		for _, adjust := range [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}} {
			ce := Elf{elf.x + adjust[0], elf.y + adjust[1]}
			if _, ok := elves[ce]; ok {
				inVicinity = true
				break
			}
		}
		if !inVicinity {
			done[elf] = true
			continue
		}

		couldPropose := false
		for _, look := range looks {
			found := false
			for _, check := range look.checks {
				ce := Elf{elf.x + check[0], elf.y + check[1]}
				if _, ok := elves[ce]; ok {
					found = true
					break
				}
			}
			if !found {
				ce := Elf{elf.x + look.move[0], elf.y + look.move[1]}
				if pe, ok := proposed[ce]; ok {
					delete(done, ce)
					done[pe] = true
					break
				}
				done[ce] = true
				proposed[ce] = elf
				couldPropose = true
				break
			}
		}

		if !couldPropose {
			done[elf] = true
		}
	}
	return done
}

func countEmpty(elves map[Elf]bool) int {
	minX, minY, maxX, maxY := 0, 0, 0, 0
	first := true
	for elf := range elves {
		if first {
			minX = elf.x
			maxX = elf.x
			minY = elf.y
			maxY = elf.y
			first = false
			continue
		}
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}

	empty := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if _, ok := elves[Elf{x, y}]; !ok {
				empty++
			}
		}
	}
	return empty
}
