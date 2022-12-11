package day11

import (
	"advent-of-code-2022/utils"
	"log"
	"sort"
	"strings"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day11", "day-11-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int64 {
	start := time.Now().UnixMilli()
	monkeys := parseMonkeys(lines)
	for i := 0; i < 20; i++ {
		executeRound(monkeys)
	}
	ans := monkeyBusiness(monkeys)
	end := time.Now().UnixMilli()
	log.Printf("Day 11, Part 1 (%dms): Monkey Business = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int64 {
	start := time.Now().UnixMilli()
	monkeys := parseMonkeys(lines)
	adjustReliefFactor(monkeys)
	for i := 0; i < 10000; i++ {
		executeRound(monkeys)
	}
	ans := monkeyBusiness(monkeys)
	end := time.Now().UnixMilli()
	log.Printf("Day 11, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

func parseMonkeys(lines []string) []*Monkey {
	var monkeys []*Monkey
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, NewMonkey([]string{
			lines[i], lines[i+1], lines[i+2], lines[i+3], lines[i+4], lines[i+5],
		}))
	}
	return monkeys
}

func adjustReliefFactor(monkeys []*Monkey) {
	// compute the relief factor by using the product of all the worry test values
	// to enable modulus operations that will keep the integrity of the values
	// while maintaining a small number.
	r := 1
	for _, m := range monkeys {
		r *= m.worryTest.constant
	}

	// apply the updated factor to all the monkeys
	for _, m := range monkeys {
		m.reliefByMod = true
		m.reliefFactor = r
	}
}

func monkeyBusiness(monkeys []*Monkey) int64 {
	var i []int
	for _, m := range monkeys {
		i = append(i, m.inspectedCnt)
	}
	sort.Ints(i)
	return int64(i[len(i)-1]) * int64(i[len(i)-2])
}

func executeRound(monkeys []*Monkey) {
	for _, m := range monkeys {
		m.inspectItems(monkeys)
	}
}

type Monkey struct {
	reliefByMod  bool
	reliefFactor int
	items        []int64
	worryChange  Operation
	worryTest    Operation
	throwTrue    int
	throwFalse   int
	inspectedCnt int
}

func (m *Monkey) addItem(item int64) {
	m.items = append(m.items, item)
}

func (m *Monkey) inspectItems(monkeys []*Monkey) {
	for _, i := range m.items {
		m.inspectedCnt++
		worry := m.worryChange.execute(i)
		worry = m.applyRelief(worry)
		switch m.worryTest.execute(worry) == 0 {
		case true:
			monkeys[m.throwTrue].addItem(worry)
		default:
			monkeys[m.throwFalse].addItem(worry)
		}
	}
	m.items = m.items[:0]
}

func (m *Monkey) applyRelief(worry int64) int64 {
	w := worry
	switch m.reliefByMod {
	case true:
		w %= int64(m.reliefFactor)
	case false:
		w /= int64(m.reliefFactor)
	}
	return w
}

func NewMonkey(lines []string) *Monkey {
	m := &Monkey{}
	m.reliefFactor = 3
	integers := utils.ReadIntegersFromLine(strings.Split(lines[1], ": ")[1])
	for _, i := range integers {
		m.items = append(m.items, int64(i))
	}

	s := strings.Split(lines[2], " ")
	m.worryChange = Operation{
		math:     MathType(s[len(s)-2][0]),
		constant: utils.Number(s[len(s)-1]),
	}

	s = strings.Split(lines[3], " ")
	m.worryTest = Operation{
		math:     Modulus,
		constant: utils.Number(s[len(s)-1]),
	}

	s = strings.Split(lines[4], " ")
	m.throwTrue = utils.Number(s[len(s)-1])

	s = strings.Split(lines[5], " ")
	m.throwFalse = utils.Number(s[len(s)-1])

	return m
}

type MathType rune

var (
	Add      MathType = '+'
	Multiply MathType = '*'
	Modulus  MathType = '%'
)

type Operation struct {
	math     MathType
	constant int
}

func (o Operation) execute(v int64) int64 {
	by := v
	if o.constant > 0 {
		by = int64(o.constant)
	}
	switch o.math {
	case Add:
		return v + by
	case Multiply:
		return v * by
	default:
		return v % by
	}
}
