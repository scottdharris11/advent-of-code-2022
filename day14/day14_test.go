package day14

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 24, solvePart1(testInput))
	assert.Equal(t, 592, solvePart1(utils.ReadLines("day14", "day-14-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 93, solvePart2(testInput))
	assert.Equal(t, 30367, solvePart2(utils.ReadLines("day14", "day-14-input.txt")))
}

func TestCave_DropSand(t *testing.T) {
	cave := Cave{
		xOffset: 494,
		maxX:    9,
		maxY:    9,
		grid: [][]rune{
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("    #   ##"),
			[]rune("    #   # "),
			[]rune("  ###   # "),
			[]rune("        # "),
			[]rune("        # "),
			[]rune("######### "),
		},
	}

	sandPoints := []Point{
		{6, 8}, {5, 8}, {7, 8}, {6, 7}, {4, 8},
		{5, 7}, {7, 7}, {6, 6}, {3, 8}, {4, 7},
		{5, 6}, {7, 6}, {6, 5}, {5, 5}, {7, 5},
		{6, 4}, {5, 4}, {7, 4}, {6, 3}, {5, 3},
		{7, 3}, {6, 2}, {3, 5}, {1, 8},
	}
	for _, s := range sandPoints {
		abyss := cave.DropSand()
		assert.False(t, abyss)
		assert.Equal(t, Sand, cave.grid[s.y][s.x])
	}

	abyss := cave.DropSand()
	assert.True(t, abyss)
}

func TestNewCave(t *testing.T) {
	cave := Cave{
		xOffset: 494,
		maxX:    9,
		maxY:    9,
		floor:   false,
		grid: [][]rune{
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("    #   ##"),
			[]rune("    #   # "),
			[]rune("  ###   # "),
			[]rune("        # "),
			[]rune("        # "),
			[]rune("######### "),
		},
		lSand: make([]int, 9),
		rSand: make([]int, 9),
	}

	caveWithFloor := Cave{
		xOffset: 494,
		maxX:    9,
		maxY:    11,
		floor:   true,
		grid: [][]rune{
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("    #   ##"),
			[]rune("    #   # "),
			[]rune("  ###   # "),
			[]rune("        # "),
			[]rune("        # "),
			[]rune("######### "),
			[]rune("          "),
			[]rune("          "),
		},
		lSand: make([]int, 11),
		rSand: make([]int, 11),
	}

	tests := []struct {
		rocks    []Rock
		floor    bool
		expected Cave
	}{
		{[]Rock{
			{[]Point{{498, 4}, {498, 6}, {496, 6}}},
			{[]Point{{503, 4}, {502, 4}, {502, 9}, {494, 9}}},
		}, false, cave},
		{[]Rock{
			{[]Point{{496, 6}, {498, 6}, {498, 4}}},
			{[]Point{{494, 9}, {502, 9}, {502, 4}, {503, 4}}},
		}, false, cave},
		{[]Rock{
			{[]Point{{498, 4}, {498, 6}, {496, 6}}},
			{[]Point{{503, 4}, {502, 4}, {502, 9}, {494, 9}}},
		}, true, caveWithFloor},
		{[]Rock{
			{[]Point{{496, 6}, {498, 6}, {498, 4}}},
			{[]Point{{494, 9}, {502, 9}, {502, 4}, {503, 4}}},
		}, true, caveWithFloor},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, NewCave(tt.rocks, tt.floor))
		})
	}
}

func TestNewRock(t *testing.T) {
	tests := []struct {
		input string
		rock  Rock
	}{
		{"498,4 -> 498,6 -> 496,6", Rock{[]Point{
			{498, 4},
			{498, 6},
			{496, 6},
		}}},
		{"503,4 -> 502,4 -> 502,9 -> 494,9", Rock{[]Point{
			{503, 4},
			{502, 4},
			{502, 9},
			{494, 9},
		}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.rock, NewRock(tt.input))
		})
	}
}

func TestNewPoint(t *testing.T) {
	tests := []struct {
		input string
		point Point
	}{
		{"498,4", Point{498, 4}},
		{"498,6", Point{498, 6}},
		{"496,6", Point{496, 6}},
		{"503,4", Point{503, 4}},
		{"502,4", Point{502, 4}},
		{"502,9", Point{502, 9}},
		{"494,9", Point{494, 9}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.point, NewPoint(tt.input))
		})
	}
}
