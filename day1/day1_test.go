package day1

import (
	"advent-of-code-2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 24000, solvePart1(testValues))
	assert.Equal(t, 71924, solvePart1(utils.ReadLines("day1", "day-1-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 45000, solvePart2(testValues))
	assert.Equal(t, 210406, solvePart2(utils.ReadLines("day1", "day-1-input.txt")))
}
