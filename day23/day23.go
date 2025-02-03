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

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	elves := parseInput(lines)
	_, elves = placeElves(elves, 10)
	ans := countEmpty(elves)
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 1 (%dms): Empty Space = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	elves := parseInput(lines)
	round, _ := placeElves(elves, -1)
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 2 (%dms): Round = %d", end-start, round)
	return round
}

type Elf struct {
	x int
	y int
}

func (e Elf) clear(elves map[Elf]bool, locs [][]int) bool {
	for _, adjust := range locs {
		ce := Elf{e.x + adjust[0], e.y + adjust[1]}
		if _, ok := elves[ce]; ok {
			return false
		}
	}
	return true
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

var AllAround = [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}
var LOOKS = []Look{
	{checks: [][]int{{0, -1}, {1, -1}, {-1, -1}}, move: []int{0, -1}},
	{checks: [][]int{{0, 1}, {1, 1}, {-1, 1}}, move: []int{0, 1}},
	{checks: [][]int{{-1, 0}, {-1, -1}, {-1, 1}}, move: []int{-1, 0}},
	{checks: [][]int{{1, 0}, {1, -1}, {1, 1}}, move: []int{1, 0}},
}

func placeElves(elves map[Elf]bool, rounds int) (int, map[Elf]bool) {
	li := 0
	round := 1
	for {
		var looks []Look
		looks = append(looks, LOOKS[li:]...)
		looks = append(looks, LOOKS[:li]...)
		li = (li + 1) % len(LOOKS)
		updated, ne := doRound(elves, looks)
		elves = ne
		if !updated || round == rounds {
			break
		}
		round++
	}
	return round, elves
}

func doRound(elves map[Elf]bool, looks []Look) (bool, map[Elf]bool) {
	done := make(map[Elf]bool, len(elves))
	proposed := make(map[Elf]Elf, len(elves))
	updated := 0
	for elf := range elves {
		if elf.clear(elves, AllAround) {
			done[elf] = true
			continue
		}

		couldPropose := false
		for _, look := range looks {
			if elf.clear(elves, look.checks) {
				ce := Elf{elf.x + look.move[0], elf.y + look.move[1]}
				if pe, ok := proposed[ce]; ok {
					delete(done, ce)
					done[pe] = true
					updated--
					break
				}
				done[ce] = true
				proposed[ce] = elf
				couldPropose = true
				updated++
				break
			}
		}

		if !couldPropose {
			done[elf] = true
		}
	}
	return updated > 0, done
}

func countEmpty(elves map[Elf]bool) int {
	minX, minY, maxX, maxY := 0, 0, 0, 0
	first := true
	for elf := range elves {
		if first {
			minX, maxX, minY, maxY = elf.x, elf.x, elf.y, elf.y
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
