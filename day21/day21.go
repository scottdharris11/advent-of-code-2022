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
	ans, _ := computeResult("root", formulas, values)
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 1 (%dms): Root = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	var formulas map[string]Formula
	var values map[string]int
	values, formulas = parseInput(lines)
	root := formulas["root"]
	found, path := pathToKey("humn", root.key1, formulas)
	valueKey := root.key2
	if !found {
		_, path = pathToKey("humn", root.key2, formulas)
		valueKey = root.key1
	}
	_, output := computeResult("root", formulas, values)
	needed := output[valueKey]
	_, formulas = parseInput(lines)
	ans := findNecessaryValue(needed, path, formulas, values)

	values, formulas = parseInput(lines)
	values["humn"] = ans
	check, _ := computeResult(path[0], formulas, values)
	if check != needed {
		log.Printf("Something is wrong with ans %d, return value not matching expected %d != %d", ans, needed, check)
		ans = -1
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 2 (%dms): humn entry = %d", end-start, ans)
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

func (f Formula) computeKey(key string, result int, values map[string]int) int {
	if key == f.key1 {
		return f.computeKey1(result, values)
	}
	return f.computeKey2(result, values)
}

func (f Formula) computeKey1(result int, values map[string]int) int {
	switch f.operation {
	case "+":
		return result - values[f.key2]
	case "-":
		return result + values[f.key2]
	case "*":
		return result / values[f.key2]
	case "/":
		return result * values[f.key2]
	}
	return -1
}

func (f Formula) computeKey2(result int, values map[string]int) int {
	switch f.operation {
	case "+":
		return result - values[f.key1]
	case "-":
		return values[f.key1] - result
	case "*", "/":
		return result / values[f.key1]
	}
	return -1
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

func computeResult(key string, formulas map[string]Formula, values map[string]int) (int, map[string]int) {
	for {
		updated := false
		for k, f := range formulas {
			computed, result := f.compute(values)
			if computed {
				values[k] = result
				delete(formulas, k)
				updated = true
				if k == key {
					return result, values
				}
			}
		}
		if !updated || len(formulas) == 0 {
			break
		}
	}
	return -1, nil
}

func pathToKey(from string, to string, formulas map[string]Formula) (bool, []string) {
	if from == to {
		return true, []string{to}
	}
	for k, f := range formulas {
		if f.key1 == from || f.key2 == from {
			rp, path := pathToKey(k, to, formulas)
			if rp {
				if from == f.key2 {
					return true, append(path, f.key2)
				}
				return true, append(path, f.key1)
			}
		}
	}
	return false, nil
}

func findNecessaryValue(expected int, path []string, formulas map[string]Formula, values map[string]int) int {
	if len(path) == 1 {
		return expected
	}
	f := formulas[path[0]]
	ne := f.computeKey(path[1], expected, values)
	return findNecessaryValue(ne, path[1:], formulas, values)
}
