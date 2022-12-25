package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
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
	// assert.Equal(t, 3472, solvePart2(testInput))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day19", "day-19-input.txt")[0:3]))
}

func TestNewBlueprint(t *testing.T) {
	tests := []struct {
		line      string
		blueprint Blueprint
	}{
		{
			"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
			Blueprint{
				id:       1,
				ore:      RobotCost{ore: 4},
				clay:     RobotCost{ore: 2},
				obsidian: RobotCost{ore: 3, clay: 14},
				geode:    RobotCost{ore: 2, obsidian: 7},
			},
		},
		{
			"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
			Blueprint{
				id:       2,
				ore:      RobotCost{ore: 2},
				clay:     RobotCost{ore: 3},
				obsidian: RobotCost{ore: 3, clay: 8},
				geode:    RobotCost{ore: 3, obsidian: 12},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			assert.Equal(t, tt.blueprint, *NewBlueprint(tt.line))
		})
	}
}
