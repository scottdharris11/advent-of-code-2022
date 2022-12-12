package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testGrid = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 0, solvePart1(testGrid))
	assert.Equal(t, 0, solvePart1(utils.ReadLines("day12", "day-12-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(testGrid))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day12", "day-12-input.txt")))
}
