package day1

import (
	"log"
	"sort"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day1", "day-1-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	mostCals := 0
	elves := parseElves(lines)
	for _, e := range elves {
		if e.totalCalories > mostCals {
			mostCals = e.totalCalories
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 1, Part 1 (%dms): Most Calories = %d", end-start, mostCals)
	return mostCals
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	elves := parseElves(lines)
	sort.Sort(elfList(elves))
	topCals := 0
	for i := 0; i < 3; i++ {
		topCals += elves[i].totalCalories
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 1, Part 2 (%dms): Top 3 Calories = %d", end-start, topCals)
	return topCals
}

func parseElves(lines []string) []Elf {
	var elves []Elf
	elf := Elf{}
	for _, line := range lines {
		if line == "" {
			elves = append(elves, elf)
			elf = Elf{}
			continue
		}

		cals := utils.Number(line)
		elf.totalCalories += cals
		elf.snacks = append(elf.snacks, cals)
	}
	elves = append(elves, elf)
	return elves
}

type Elf struct {
	totalCalories int
	snacks        []int
}

type elfList []Elf

func (e elfList) Len() int {
	return len(e)
}

func (e elfList) Less(i, j int) bool {
	return e[i].totalCalories > e[j].totalCalories
}

func (e elfList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
