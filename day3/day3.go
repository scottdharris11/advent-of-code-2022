package day3

import (
	"advent-of-code-2022/utils"
	"log"
	"strings"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day3", "day-3-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	totalPriority := 0
	for _, l := range lines {
		r := newRuckSack(l)
		for _, d := range r.misplaced() {
			totalPriority += priority(d)
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 3, Part 1 (%dms): Priority Sum = %d", end-start, totalPriority)
	return totalPriority
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	totalPriority := 0
	groupCnt := len(lines) / 3
	for i := 0; i < groupCnt; i++ {
		g := Group{
			ruckSacks: []RuckSack{
				newRuckSack(lines[i*3]),
				newRuckSack(lines[(i*3)+1]),
				newRuckSack(lines[(i*3)+2]),
			},
		}
		totalPriority += priority(g.badge())
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 3, Part 2 (%dms): Badge Sum = %d", end-start, totalPriority)
	return totalPriority
}

func newRuckSack(line string) RuckSack {
	r := RuckSack{}
	breakpoint := len(line) / 2
	r.container1 = append(r.container1, []rune(line[:breakpoint])...)
	r.container2 = append(r.container2, []rune(line[breakpoint:])...)
	return r
}

type RuckSack struct {
	container1 []rune
	container2 []rune
}

func (r RuckSack) misplaced() []rune {
	var misplaced []rune
	checked := make(map[rune]bool)
	for _, c1 := range r.container1 {
		if checked[c1] {
			continue
		}
		checked[c1] = true
		if strings.ContainsRune(string(r.container2), c1) {
			misplaced = append(misplaced, c1)
		}
	}
	return misplaced
}

func (r RuckSack) items() []rune {
	return []rune(string(r.container1) + string(r.container2))
}

func (r RuckSack) containsItem(item rune) bool {
	return strings.ContainsRune(string(r.container1), item) ||
		strings.ContainsRune(string(r.container2), item)
}

type Group struct {
	ruckSacks []RuckSack
}

func (g Group) badge() rune {
	var badge rune
	for _, item := range g.ruckSacks[0].items() {
		found := true
		for i := 1; i < len(g.ruckSacks); i++ {
			if !g.ruckSacks[i].containsItem(item) {
				found = false
				break
			}
		}
		if found {
			return item
		}
	}
	return badge
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	return int(r) - int('A') + 27
}
