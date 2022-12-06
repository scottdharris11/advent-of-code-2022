package day5

import (
	"advent-of-code-2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	" ",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, "CMZ", solvePart1(testInput))
	assert.Equal(t, "VJSFHWGFT", solvePart1(utils.ReadLines("day5", "day-5-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, "MCD", solvePart2(testInput))
	assert.Equal(t, "LCTQFBVZV", solvePart2(utils.ReadLines("day5", "day-5-input.txt")))
}

func TestParseInput(t *testing.T) {
	stacks, instructions := parseInput(testInput)

	assert.Equal(t, 3, len(stacks))
	assert.Equal(t, 'N', stacks[0].Pop())
	assert.Equal(t, 'Z', stacks[0].Pop())
	assert.Equal(t, 'D', stacks[1].Pop())
	assert.Equal(t, 'C', stacks[1].Pop())
	assert.Equal(t, 'M', stacks[1].Pop())
	assert.Equal(t, 'P', stacks[2].Pop())

	assert.Equal(t, []Instruction{
		{count: 1, from: 1, to: 0},
		{count: 3, from: 0, to: 2},
		{count: 2, from: 1, to: 0},
		{count: 1, from: 0, to: 1},
	}, instructions)
}
