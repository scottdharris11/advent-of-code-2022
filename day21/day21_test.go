package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"root: pppw + sjmn",
	"dbpl: 5",
	"cczh: sllz + lgvd",
	"zczc: 2",
	"ptdq: humn - dvpt",
	"dvpt: 3",
	"lfqf: 4",
	"humn: 5",
	"ljgn: 2",
	"sjmn: drzm * dbpl",
	"sllz: 4",
	"pppw: cczh / lfqf",
	"lgvd: ljgn * ptdq",
	"drzm: hmdt - zczc",
	"hmdt: 32",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 152, solvePart1(testInput))
	assert.Equal(t, 78342931359552, solvePart1(utils.ReadLines("day21", "day-21-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(testInput))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day21", "day-21-input.txt")))
}
