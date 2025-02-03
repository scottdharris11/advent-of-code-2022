package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	".....",
	"..##.",
	"..#..",
	".....",
	"..##.",
	".....",
}

var testInput2 = []string{
	"..............",
	"..............",
	".......#......",
	".....###.#....",
	"...#...#.#....",
	"....#...##....",
	"...#.###......",
	"...##.#.##....",
	"....#..#......",
	"..............",
	"..............",
	"..............",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 25, solvePart1(testInput))
	assert.Equal(t, 110, solvePart1(testInput2))
	assert.Equal(t, 4241, solvePart1(utils.ReadLines("day23", "day-23-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(testInput))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day23", "day-23-input.txt")))
}
