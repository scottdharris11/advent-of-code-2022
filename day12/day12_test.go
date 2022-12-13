package day12

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testGrid = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 31, solvePart1(testGrid))
	assert.Equal(t, 447, solvePart1(utils.ReadLines("day12", "day-12-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 29, solvePart2(testGrid))
	assert.Equal(t, 446, solvePart2(utils.ReadLines("day12", "day-12-input.txt")))
}

func TestNewRoutePlanner(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	assert.Equal(t, []string{
		"aabqponm",
		"abcryxxl",
		"accszzxk",
		"acctuvwj",
		"abdefghi",
	}, rp.grid)
	assert.Equal(t, Position{0, 0}, rp.startLocation)
	assert.Equal(t, Position{5, 2}, rp.goalLocation)
}

func TestRoute(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	assert.Equal(t, 31, rp.Route())
}

func TestGoal(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	tests := []struct {
		state    RouteState
		expected bool
	}{
		{RouteState{currentX: 0, currentY: 0}, false},
		{RouteState{currentX: 1, currentY: 1}, false},
		{RouteState{currentX: 5, currentY: 2}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.state), func(t *testing.T) {
			assert.Equal(t, tt.expected, rp.Goal(tt.state))
		})
	}
}

func TestDistanceFromGoal(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	tests := []struct {
		state    RouteState
		expected int
	}{
		{RouteState{currentX: 0, currentY: 0}, 7},
		{RouteState{currentX: 1, currentY: 1}, 5},
		{RouteState{currentX: 5, currentY: 2}, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.state), func(t *testing.T) {
			assert.Equal(t, tt.expected, rp.DistanceFromGoal(tt.state))
		})
	}
}

func TestAvailablePositions(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	tests := []struct {
		state    RouteState
		expected []utils.SearchMove
	}{
		{RouteState{currentX: 1, currentY: 1}, []utils.SearchMove{
			{Cost: 1, State: RouteState{currentX: 1, currentY: 0}},
			{Cost: 1, State: RouteState{currentX: 1, currentY: 2}},
			{Cost: 1, State: RouteState{currentX: 0, currentY: 1}},
			{Cost: 1, State: RouteState{currentX: 2, currentY: 1}},
		}},
		{RouteState{currentX: 3, currentY: 2}, []utils.SearchMove{
			{Cost: 1, State: RouteState{currentX: 3, currentY: 1}},
			{Cost: 1, State: RouteState{currentX: 3, currentY: 3}},
			{Cost: 1, State: RouteState{currentX: 2, currentY: 2}},
		}},
		{RouteState{currentX: 0, currentY: 2}, []utils.SearchMove{
			{Cost: 1, State: RouteState{currentX: 0, currentY: 1}},
			{Cost: 1, State: RouteState{currentX: 0, currentY: 3}},
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.state), func(t *testing.T) {
			assert.Equal(t, tt.expected, rp.PossibleNextMoves(tt.state))
		})
	}
}

func TestAvailable(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	tests := []struct {
		elevation rune
		pos       Position
		expected  bool
	}{
		{'a', Position{1, 1}, true},
		{'c', Position{1, 1}, true},
		{'a', Position{2, 1}, false},
		{'b', Position{2, 1}, true},
		{'c', Position{2, 1}, true},
		{'d', Position{2, 1}, true},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, rp.available(tt.elevation, tt.pos))
		})
	}
}

func TestElevation(t *testing.T) {
	rp := NewRoutePlanner(testGrid)
	tests := []struct {
		pos      Position
		valid    bool
		expected rune
	}{
		{Position{-1, 0}, false, ' '},
		{Position{0, -1}, false, ' '},
		{Position{8, 0}, false, ' '},
		{Position{0, 5}, false, ' '},
		{Position{0, 0}, true, 'a'},
		{Position{1, 1}, true, 'b'},
		{Position{7, 1}, true, 'l'},
		{Position{7, 4}, true, 'i'},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.pos), func(t *testing.T) {
			v, e := rp.elevation(tt.pos)
			assert.Equal(t, tt.valid, v)
			assert.Equal(t, tt.expected, e)
		})
	}
}
