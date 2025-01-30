package day21

import (
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day20", "day-20-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	values, formulas := parseInput(lines)
	for {
		updated := false
		for k, f := range formulas {
			computed, result := f.compute(values)
			if computed {
				values[k] = result
				delete(formulas, k)
				updated = true
			}
		}
		if !updated || len(formulas) == 0 {
			break
		}
	}
	ans := values["root"]
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 1 (%dms): Root = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 2 (%dms): Root = %d", end-start, ans)
	return ans
}

type Formula struct {
	key1      string
	key2      string
	operation string
}

func (f Formula) compute(values map[string]int) (bool, int) {
	val1, ok := values[f.key1]
	if !ok {
		return false, -1
	}
	val2, ok := values[f.key2]
	if !ok {
		return false, -1
	}
	switch f.operation {
	case "+":
		return true, val1 + val2
	case "-":
		return true, val1 - val2
	case "*":
		return true, val1 * val2
	case "/":
		return true, val1 / val2
	}
	return false, -1
}

func parseInput(lines []string) (map[string]int, map[string]Formula) {
	formulas := make(map[string]Formula)
	values := make(map[string]int)
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		key := fields[0]
		vf := strings.Split(fields[1], " ")
		switch len(vf) == 1 {
		case true:
			if v, err := strconv.Atoi(vf[0]); err == nil {
				values[key] = v
			}
		default:
			formulas[key] = Formula{key1: vf[0], key2: vf[2], operation: vf[1]}
		}
	}

	return values, formulas
}
