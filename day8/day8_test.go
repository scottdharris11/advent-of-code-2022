package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testLines = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 21, solvePart1(testLines))
	assert.Equal(t, 1546, solvePart1(utils.ReadLines("day8", "day-8-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 8, solvePart2(testLines))
	assert.Equal(t, 519064, solvePart2(utils.ReadLines("day8", "day-8-input.txt")))
}

func TestVisible(t *testing.T) {
	grid := utils.ReadIntegerGrid(testLines)
	tests := []struct {
		x       int
		y       int
		visible bool
	}{
		{1, 1, true},
		{2, 1, true},
		{3, 1, false},
		{1, 2, true},
		{2, 2, false},
		{3, 2, true},
		{1, 3, false},
		{2, 3, true},
		{3, 3, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.visible, visible(grid, tt.x, tt.y))
		})
	}
}

func TestScenicScore(t *testing.T) {
	grid := utils.ReadIntegerGrid(testLines)
	tests := []struct {
		x     int
		y     int
		score int
	}{
		{2, 1, 4},
		{2, 3, 8},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.score, scenicScore(grid, tt.x, tt.y))
		})
	}
}

func TestVisibleUp(t *testing.T) {
	grid := utils.ReadIntegerGrid(testLines)
	tests := []struct {
		x       int
		y       int
		visible bool
		count   int
	}{
		{1, 1, true, 1},
		{2, 1, true, 1},
		{3, 1, false, 1},
		{1, 2, false, 1},
		{2, 2, false, 1},
		{3, 2, false, 2},
		{1, 3, false, 1},
		{2, 3, false, 2},
		{3, 3, false, 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			v, c := visibleUp(grid, tt.x, tt.y)
			assert.Equal(t, tt.visible, v)
			assert.Equal(t, tt.count, c)
		})
	}
}

func TestVisibleDown(t *testing.T) {
	grid := utils.ReadIntegerGrid(testLines)
	tests := []struct {
		x       int
		y       int
		visible bool
		count   int
	}{
		{1, 1, false, 1},
		{2, 1, false, 2},
		{3, 1, false, 1},
		{1, 2, false, 2},
		{2, 2, false, 1},
		{3, 2, false, 1},
		{1, 3, false, 1},
		{2, 3, true, 1},
		{3, 3, false, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			v, c := visibleDown(grid, tt.x, tt.y)
			assert.Equal(t, tt.visible, v)
			assert.Equal(t, tt.count, c)
		})
	}
}

func TestVisibleLeft(t *testing.T) {
	grid := utils.ReadIntegerGrid(testLines)
	tests := []struct {
		x       int
		y       int
		visible bool
		count   int
	}{
		{1, 1, true, 1},
		{2, 1, false, 1},
		{3, 1, false, 1},
		{1, 2, false, 1},
		{2, 2, false, 1},
		{3, 2, false, 1},
		{1, 3, false, 1},
		{2, 3, true, 2},
		{3, 3, false, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			v, c := visibleLeft(grid, tt.x, tt.y)
			assert.Equal(t, tt.visible, v)
			assert.Equal(t, tt.count, c)
		})
	}
}

func TestVisibleRight(t *testing.T) {
	grid := utils.ReadIntegerGrid(testLines)
	tests := []struct {
		x       int
		y       int
		visible bool
		count   int
	}{
		{1, 1, false, 1},
		{2, 1, true, 2},
		{3, 1, false, 1},
		{1, 2, true, 3},
		{2, 2, false, 1},
		{3, 2, true, 1},
		{1, 3, false, 1},
		{2, 3, false, 2},
		{3, 3, false, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			v, c := visibleRight(grid, tt.x, tt.y)
			assert.Equal(t, tt.visible, v)
			assert.Equal(t, tt.count, c)
		})
	}
}
