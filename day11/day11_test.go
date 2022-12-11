package day11

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLines = []string{
	"Monkey 0:",
	"  Starting items: 79, 98",
	"  Operation: new = old * 19",
	"  Test: divisible by 23",
	"    If true: throw to monkey 2",
	"    If false: throw to monkey 3",
	"",
	"Monkey 1:",
	"  Starting items: 54, 65, 75, 74",
	"  Operation: new = old + 6",
	"  Test: divisible by 19",
	"    If true: throw to monkey 2",
	"    If false: throw to monkey 0",
	"",
	"Monkey 2:",
	"  Starting items: 79, 60, 97",
	"  Operation: new = old * old",
	"  Test: divisible by 13",
	"    If true: throw to monkey 1",
	"    If false: throw to monkey 3",
	"",
	"Monkey 3:",
	"  Starting items: 74",
	"  Operation: new = old + 3",
	"  Test: divisible by 17",
	"    If true: throw to monkey 0",
	"    If false: throw to monkey 1",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, int64(10605), solvePart1(testLines))
	assert.Equal(t, int64(95472), solvePart1(utils.ReadLines("day11", "day-11-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, int64(2713310158), solvePart2(testLines))
	assert.Equal(t, int64(0), solvePart2(utils.ReadLines("day11", "day-11-input.txt")))
}

func TestParseMonkeys(t *testing.T) {
	monkeys := parseMonkeys(testLines)
	assert.Equal(t, []*Monkey{
		{
			reliefFactor: 3,
			items:        []int64{79, 98},
			worryChange:  Operation{math: Multiply, constant: 19},
			worryTest:    Operation{math: Modulus, constant: 23},
			throwTrue:    2,
			throwFalse:   3,
		},
		{
			reliefFactor: 3,
			items:        []int64{54, 65, 75, 74},
			worryChange:  Operation{math: Add, constant: 6},
			worryTest:    Operation{math: Modulus, constant: 19},
			throwTrue:    2,
			throwFalse:   0,
		},
		{
			reliefFactor: 3,
			items:        []int64{79, 60, 97},
			worryChange:  Operation{math: Multiply, constant: 0},
			worryTest:    Operation{math: Modulus, constant: 13},
			throwTrue:    1,
			throwFalse:   3,
		},
		{
			reliefFactor: 3,
			items:        []int64{74},
			worryChange:  Operation{math: Add, constant: 3},
			worryTest:    Operation{math: Modulus, constant: 17},
			throwTrue:    0,
			throwFalse:   1,
		},
	}, monkeys)
}

func TestMonkeyBusiness(t *testing.T) {
	tests := []struct {
		monkeys  []*Monkey
		expected int64
	}{
		{[]*Monkey{{inspectedCnt: 101}, {inspectedCnt: 95}, {inspectedCnt: 7}, {inspectedCnt: 105}}, 10605},
		{[]*Monkey{{inspectedCnt: 101}, {inspectedCnt: 105}, {inspectedCnt: 7}, {inspectedCnt: 95}}, 10605},
		{[]*Monkey{{inspectedCnt: 7}, {inspectedCnt: 95}, {inspectedCnt: 101}, {inspectedCnt: 105}}, 10605},
		{[]*Monkey{{inspectedCnt: 101}, {inspectedCnt: 105}, {inspectedCnt: 95}, {inspectedCnt: 7}}, 10605},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, monkeyBusiness(tt.monkeys))
		})
	}
}

func TestExecuteRound(t *testing.T) {
	m0 := &Monkey{
		reliefFactor: 3,
		items:        []int64{79, 98},
		worryChange:  Operation{math: Multiply, constant: 19},
		worryTest:    Operation{math: Modulus, constant: 23},
		throwTrue:    2,
		throwFalse:   3,
	}
	m1 := &Monkey{
		reliefFactor: 3,
		items:        []int64{54, 65, 75, 74},
		worryChange:  Operation{math: Add, constant: 6},
		worryTest:    Operation{math: Modulus, constant: 19},
		throwTrue:    2,
		throwFalse:   0,
	}
	m2 := &Monkey{
		reliefFactor: 3,
		items:        []int64{79, 60, 97},
		worryChange:  Operation{math: Multiply, constant: 0},
		worryTest:    Operation{math: Modulus, constant: 13},
		throwTrue:    1,
		throwFalse:   3,
	}
	m3 := &Monkey{
		reliefFactor: 3,
		items:        []int64{74},
		worryChange:  Operation{math: Add, constant: 3},
		worryTest:    Operation{math: Modulus, constant: 17},
		throwTrue:    0,
		throwFalse:   1,
	}
	monkeys := []*Monkey{m0, m1, m2, m3}

	executeRound(monkeys)

	assert.Equal(t, []int64{20, 23, 27, 26}, m0.items)
	assert.Equal(t, []int64{2080, 25, 167, 207, 401, 1046}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{695, 10, 71, 135, 350}, m0.items)
	assert.Equal(t, []int64{43, 49, 58, 55, 362}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{16, 18, 21, 20, 122}, m0.items)
	assert.Equal(t, []int64{1468, 22, 150, 286, 739}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{491, 9, 52, 97, 248, 34}, m0.items)
	assert.Equal(t, []int64{39, 45, 43, 258}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{15, 17, 16, 88, 1037}, m0.items)
	assert.Equal(t, []int64{20, 110, 205, 524, 72}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{8, 70, 176, 26, 34}, m0.items)
	assert.Equal(t, []int64{481, 32, 36, 186, 2190}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{162, 12, 14, 64, 732, 17}, m0.items)
	assert.Equal(t, []int64{148, 372, 55, 72}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{51, 126, 20, 26, 136}, m0.items)
	assert.Equal(t, []int64{343, 26, 30, 1546, 36}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{116, 10, 12, 517, 14}, m0.items)
	assert.Equal(t, []int64{108, 267, 43, 55, 288}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	executeRound(monkeys)

	assert.Equal(t, []int64{91, 16, 20, 98}, m0.items)
	assert.Equal(t, []int64{481, 245, 22, 26, 1092, 30}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	for i := 0; i < 5; i++ {
		executeRound(monkeys)
	}

	assert.Equal(t, []int64{83, 44, 8, 184, 9, 20, 26, 102}, m0.items)
	assert.Equal(t, []int64{110, 36}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	for i := 0; i < 5; i++ {
		executeRound(monkeys)
	}

	assert.Equal(t, []int64{10, 12, 14, 26, 34}, m0.items)
	assert.Equal(t, []int64{245, 93, 53, 199, 115}, m1.items)
	assert.Equal(t, []int64{}, m2.items)
	assert.Equal(t, []int64{}, m3.items)

	assert.Equal(t, 101, m0.inspectedCnt)
	assert.Equal(t, 95, m1.inspectedCnt)
	assert.Equal(t, 7, m2.inspectedCnt)
	assert.Equal(t, 105, m3.inspectedCnt)
}

func TestExecuteRound_NoRelief(t *testing.T) {
	m0 := &Monkey{
		reliefFactor: 0,
		items:        []int64{79, 98},
		worryChange:  Operation{math: Multiply, constant: 19},
		worryTest:    Operation{math: Modulus, constant: 23},
		throwTrue:    2,
		throwFalse:   3,
	}
	m1 := &Monkey{
		reliefFactor: 0,
		items:        []int64{54, 65, 75, 74},
		worryChange:  Operation{math: Add, constant: 6},
		worryTest:    Operation{math: Modulus, constant: 19},
		throwTrue:    2,
		throwFalse:   0,
	}
	m2 := &Monkey{
		reliefFactor: 0,
		items:        []int64{79, 60, 97},
		worryChange:  Operation{math: Multiply, constant: 0},
		worryTest:    Operation{math: Modulus, constant: 13},
		throwTrue:    1,
		throwFalse:   3,
	}
	m3 := &Monkey{
		reliefFactor: 0,
		items:        []int64{74},
		worryChange:  Operation{math: Add, constant: 3},
		worryTest:    Operation{math: Modulus, constant: 17},
		throwTrue:    0,
		throwFalse:   1,
	}
	monkeys := []*Monkey{m0, m1, m2, m3}

	executeRound(monkeys)

	assert.Equal(t, 2, m0.inspectedCnt)
	assert.Equal(t, 4, m1.inspectedCnt)
	assert.Equal(t, 3, m2.inspectedCnt)
	assert.Equal(t, 6, m3.inspectedCnt)

	for i := 0; i < 19; i++ {
		executeRound(monkeys)
	}

	assert.Equal(t, 99, m0.inspectedCnt)
	assert.Equal(t, 97, m1.inspectedCnt)
	assert.Equal(t, 8, m2.inspectedCnt)
	assert.Equal(t, 103, m3.inspectedCnt)

	for i := 0; i < 980; i++ {
		executeRound(monkeys)
	}

	assert.Equal(t, 5204, m0.inspectedCnt)
	assert.Equal(t, 4792, m1.inspectedCnt)
	assert.Equal(t, 199, m2.inspectedCnt)
	assert.Equal(t, 5192, m3.inspectedCnt)
}

func TestMonkey_InspectItems(t *testing.T) {
	m0 := &Monkey{
		reliefFactor: 3,
		items:        []int64{79, 98},
		worryChange:  Operation{math: Multiply, constant: 19},
		worryTest:    Operation{math: Modulus, constant: 23},
		throwTrue:    2,
		throwFalse:   3,
	}
	m1 := &Monkey{}
	m2 := &Monkey{}
	m3 := &Monkey{}
	monkeys := []*Monkey{m0, m1, m2, m3}

	m0.inspectItems(monkeys)

	assert.Equal(t, 0, len(m0.items))
	assert.Equal(t, 0, len(m1.items))
	assert.Equal(t, 0, len(m2.items))
	assert.Equal(t, []int64{500, 620}, m3.items)
}

func TestNewMonkey(t *testing.T) {
	tests := []struct {
		lines    []string
		expected Monkey
	}{
		{[]string{
			"Monkey 0:",
			"  Starting items: 79, 98",
			"  Operation: new = old * 19",
			"  Test: divisible by 23",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 3",
		}, Monkey{
			reliefFactor: 3,
			items:        []int64{79, 98},
			worryChange:  Operation{math: Multiply, constant: 19},
			worryTest:    Operation{math: Modulus, constant: 23},
			throwTrue:    2,
			throwFalse:   3,
		}},
		{[]string{
			"Monkey 1:",
			"  Starting items: 54, 65, 75, 74",
			"  Operation: new = old + 6",
			"  Test: divisible by 19",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 0",
		}, Monkey{
			reliefFactor: 3,
			items:        []int64{54, 65, 75, 74},
			worryChange:  Operation{math: Add, constant: 6},
			worryTest:    Operation{math: Modulus, constant: 19},
			throwTrue:    2,
			throwFalse:   0,
		}},
		{[]string{
			"Monkey 2:",
			"  Starting items: 79, 60, 97",
			"  Operation: new = old * old",
			"  Test: divisible by 13",
			"    If true: throw to monkey 1",
			"    If false: throw to monkey 3",
		}, Monkey{
			reliefFactor: 3,
			items:        []int64{79, 60, 97},
			worryChange:  Operation{math: Multiply, constant: 0},
			worryTest:    Operation{math: Modulus, constant: 13},
			throwTrue:    1,
			throwFalse:   3,
		}},
		{[]string{
			"Monkey 3:",
			"  Starting items: 74",
			"  Operation: new = old + 3",
			"  Test: divisible by 17",
			"    If true: throw to monkey 0",
			"    If false: throw to monkey 1",
		}, Monkey{
			reliefFactor: 3,
			items:        []int64{74},
			worryChange:  Operation{math: Add, constant: 3},
			worryTest:    Operation{math: Modulus, constant: 17},
			throwTrue:    0,
			throwFalse:   1,
		}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, *NewMonkey(tt.lines))
		})
	}
}

func TestOperation_Execute(t *testing.T) {
	tests := []struct {
		operation Operation
		input     int64
		expected  int64
	}{
		{Operation{math: Add, constant: 4}, 5, 9},
		{Operation{math: Add}, 5, 10},
		{Operation{math: Multiply, constant: 4}, 5, 20},
		{Operation{math: Multiply}, 5, 25},
		{Operation{math: Modulus, constant: 4}, 5, 1},
		{Operation{math: Modulus}, 5, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.operation.execute(tt.input))
		})
	}
}
