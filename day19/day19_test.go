package day19

import (
	"advent-of-code-2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
	"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 33, solvePart1(testInput))
	assert.Equal(t, 1834, solvePart1(utils.ReadLines("day19", "day-19-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3472, solvePart2(testInput))
	assert.Equal(t, 2240, solvePart2(utils.ReadLines("day19", "day-19-input.txt")))
}
