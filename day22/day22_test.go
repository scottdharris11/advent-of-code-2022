package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"        ...#",
	"        .#..",
	"        #...",
	"        ....",
	"...#.......#",
	"........#...",
	"..#....#....",
	"..........#.",
	"        ...#....",
	"        .....#..",
	"        .#......",
	"        ......#.",
	"",
	"10R5L5R10L4R5L5",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 6032, solvePart1(testInput))
	assert.Equal(t, 133174, solvePart1(utils.ReadLines("day22", "day-22-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(testInput))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day22", "day-22-input.txt")))
}
