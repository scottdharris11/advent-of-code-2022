package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"#.#####",
	"#.....#",
	"#>....#",
	"#.....#",
	"#...v.#",
	"#.....#",
	"#####.#",
}

var testInput2 = []string{
	"#.######",
	"#>>.<^<#",
	"#.<..<<#",
	"#>v.><>#",
	"#<^v^^>#",
	"######.#",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 10, solvePart1(testInput))
	assert.Equal(t, 18, solvePart1(testInput2))
	assert.Equal(t, 242, solvePart1(utils.ReadLines("day24", "day-24-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(testInput))
	assert.Equal(t, 0, solvePart2(testInput2))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day24", "day-24-input.txt")))
}
