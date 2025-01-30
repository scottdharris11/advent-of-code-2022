package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{"1", "2", "-3", "3", "-2", "0", "4"}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 3, solvePart1(testInput))
	assert.Equal(t, 7228, solvePart1(utils.ReadLines("day20", "day-20-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	// assert.Equal(t, 0, solvePart2(testInput))
	// assert.Equal(t, 0, solvePart2(utils.ReadLines("day20", "day-20-input.txt")))
}
