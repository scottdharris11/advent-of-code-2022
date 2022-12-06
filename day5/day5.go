package day5

import (
	"log"
	"regexp"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day5", "day-5-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) string {
	start := time.Now().UnixMilli()
	stacks, instructions := parseInput(lines)
	for _, i := range instructions {
		for j := 0; j < i.count; j++ {
			stacks[i.to].Push(stacks[i.from].Pop())
		}
	}
	ans := stackTops(stacks)
	end := time.Now().UnixMilli()
	log.Printf("Day 5, Part 1 (%dms): Stack Tops = %s", end-start, ans)
	return ans
}

func solvePart2(lines []string) string {
	start := time.Now().UnixMilli()
	stacks, instructions := parseInput(lines)
	for _, i := range instructions {
		var hold []interface{}
		for j := 0; j < i.count; j++ {
			hold = append(hold, stacks[i.from].Pop())
		}
		for j := len(hold) - 1; j >= 0; j-- {
			stacks[i.to].Push(hold[j])
		}
	}
	ans := stackTops(stacks)
	end := time.Now().UnixMilli()
	log.Printf("Day 5, Part 2 (%dms): Stack Tops = %s", end-start, ans)
	return ans
}

func stackTops(stacks []utils.Stack) string {
	var output []rune
	for _, s := range stacks {
		c := s.Peek()
		if c == nil {
			c = ' '
		}
		output = append(output, c.(rune))
	}
	return string(output)
}

type Instruction struct {
	from  int
	to    int
	count int
}

var instructionParse = regexp.MustCompile(`move ([\d]+) from ([\d]+) to ([\d]+)`)

func parseInput(lines []string) ([]utils.Stack, []Instruction) {
	eoc := 0
	var crateIndexes []int
	for i, line := range lines {
		if strings.HasPrefix(line, " 1") {
			eoc = i
			for j, c := range line {
				if c != ' ' {
					crateIndexes = append(crateIndexes, j)
				}
			}
			break
		}
	}

	stacks := make([]utils.Stack, len(crateIndexes))
	for i := 0; i < len(crateIndexes); i++ {
		stacks[i] = utils.Stack{}
	}

	for i := eoc - 1; i >= 0; i-- {
		for j, c := range crateIndexes {
			if lines[i][c] != ' ' {
				stacks[j].Push(rune(lines[i][c]))
			}
		}
	}

	var instructions []Instruction
	for i := eoc + 2; i < len(lines); i++ {
		match := instructionParse.FindStringSubmatch(lines[i])
		if match == nil {
			continue
		}
		instructions = append(instructions, Instruction{
			count: utils.Number(match[1]),
			from:  utils.Number(match[2]) - 1,
			to:    utils.Number(match[3]) - 1,
		})
	}

	return stacks, instructions
}
