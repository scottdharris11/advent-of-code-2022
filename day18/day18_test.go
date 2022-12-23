package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"2,2,2",
	"1,2,2",
	"3,2,2",
	"2,1,2",
	"2,3,2",
	"2,2,1",
	"2,2,3",
	"2,2,4",
	"2,2,6",
	"1,2,5",
	"3,2,5",
	"2,1,5",
	"2,3,5",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 64, solvePart1(testInput))
	assert.Equal(t, 0, solvePart1(utils.ReadLines("day18", "day-18-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	// assert.Equal(t, 0, solvePart2(testInput))
	// assert.Equal(t, 0, solvePart2(utils.ReadLines("day18", "day-18-input.txt")))
}
