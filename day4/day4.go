package day4

import (
	"advent-of-code-2022/utils"
	"log"
	"regexp"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day4", "day-4-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	pairs := parseAssignments(lines)
	overlaps := 0
	for _, p := range pairs {
		if p.fullOverlap() {
			overlaps++
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 4, Part 1 (%dms): Overlaps = %d", end-start, overlaps)
	return overlaps
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	pairs := parseAssignments(lines)
	overlaps := 0
	for _, p := range pairs {
		if p.partialOverlap() {
			overlaps++
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 4, Part 2 (%dms): Overlaps = %d", end-start, overlaps)
	return overlaps
}

type SectionAssignment struct {
	min int
	max int
}

type PairAssignment struct {
	assign1 SectionAssignment
	assign2 SectionAssignment
}

func (p PairAssignment) fullOverlap() bool {
	return contains(p.assign1.min, p.assign1.max, p.assign2.min, p.assign2.max) ||
		contains(p.assign2.min, p.assign2.max, p.assign1.min, p.assign1.max)
}

func (p PairAssignment) partialOverlap() bool {
	return overlap(p.assign1.min, p.assign1.max, p.assign2.min, p.assign2.max) ||
		overlap(p.assign2.min, p.assign2.max, p.assign1.min, p.assign1.max)
}

func contains(b1 int, e1 int, b2 int, e2 int) bool {
	return b1 <= b2 && e1 >= e2
}

func overlap(b1 int, e1 int, b2 int, e2 int) bool {
	return b1 <= e2 && b2 <= e1
}

var re = regexp.MustCompile(`([\d]+)-([\d]+),([\d]+)-([\d]+)`)

func parseAssignments(lines []string) []PairAssignment {
	var pairs []PairAssignment
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		pairs = append(pairs, PairAssignment{
			assign1: SectionAssignment{min: utils.Number(match[1]), max: utils.Number(match[2])},
			assign2: SectionAssignment{min: utils.Number(match[3]), max: utils.Number(match[4])},
		})
	}
	return pairs
}
