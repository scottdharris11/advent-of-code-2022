package day9

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testLines = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

var testLines2 = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 13, solvePart1(testLines))
	assert.Equal(t, 6337, solvePart1(utils.ReadLines("day9", "day-9-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 1, solvePart2(testLines))
	assert.Equal(t, 36, solvePart2(testLines2))
	assert.Equal(t, 2455, solvePart2(utils.ReadLines("day9", "day-9-input.txt")))
}

func TestRopeKnot_Move(t *testing.T) {
	// base location
	k := NewKnot(nil)
	assert.Equal(t, 0, k.y)
	assert.Equal(t, 0, k.x)
	assert.Equal(t, 1, k.PositionsVisited())

	// up one, right two
	k.Move(1, 2)
	assert.Equal(t, 1, k.y)
	assert.Equal(t, 2, k.x)
	assert.Equal(t, 2, k.PositionsVisited())

	// up three
	k.Move(3, 0)
	assert.Equal(t, 4, k.y)
	assert.Equal(t, 2, k.x)
	assert.Equal(t, 3, k.PositionsVisited())

	// down two, left one
	k.Move(-2, -1)
	assert.Equal(t, 2, k.y)
	assert.Equal(t, 1, k.x)
	assert.Equal(t, 4, k.PositionsVisited())

	// down one, right one (repeat visit)
	k.Move(-1, 1)
	assert.Equal(t, 1, k.y)
	assert.Equal(t, 2, k.x)
	assert.Equal(t, 4, k.PositionsVisited())
}

func TestRopeKnot_Align(t *testing.T) {
	tests := []struct {
		head      Knot
		tail      Knot
		tailAfter Knot
	}{
		{BuildKnot(0, 0), BuildKnot(0, 0), BuildKnot(0, 0)},
		{BuildKnot(3, 1), BuildKnot(2, 1), BuildKnot(2, 1)},
		{BuildKnot(4, 2), BuildKnot(3, 1), BuildKnot(3, 1)},
		{BuildKnot(5, 3), BuildKnot(5, 2), BuildKnot(5, 2)},
		{BuildKnot(5, 3), BuildKnot(3, 3), BuildKnot(4, 3)},
		{BuildKnot(5, 3), BuildKnot(5, 1), BuildKnot(5, 2)},
		{BuildKnot(5, 3), BuildKnot(3, 1), BuildKnot(4, 2)},
		{BuildKnot(5, 3), BuildKnot(6, 1), BuildKnot(5, 2)},

		{BuildKnot(3, 1), BuildKnot(4, 1), BuildKnot(4, 1)},
		{BuildKnot(4, 2), BuildKnot(5, 3), BuildKnot(5, 3)},
		{BuildKnot(5, 3), BuildKnot(5, 4), BuildKnot(5, 4)},
		{BuildKnot(5, 3), BuildKnot(7, 3), BuildKnot(6, 3)},
		{BuildKnot(5, 3), BuildKnot(5, 5), BuildKnot(5, 4)},
		{BuildKnot(5, 3), BuildKnot(7, 5), BuildKnot(6, 4)},
		{BuildKnot(5, 3), BuildKnot(7, 1), BuildKnot(6, 2)},

		{BuildKnot(-3, -1), BuildKnot(-2, -1), BuildKnot(-2, -1)},
		{BuildKnot(-4, -2), BuildKnot(-3, -1), BuildKnot(-3, -1)},
		{BuildKnot(-5, -3), BuildKnot(-5, -2), BuildKnot(-5, -2)},
		{BuildKnot(-5, -3), BuildKnot(-3, -3), BuildKnot(-4, -3)},
		{BuildKnot(-5, -3), BuildKnot(-5, -1), BuildKnot(-5, -2)},
		{BuildKnot(-5, -3), BuildKnot(-3, -1), BuildKnot(-4, -2)},

		{BuildKnot(-3, -1), BuildKnot(-4, -1), BuildKnot(-4, -1)},
		{BuildKnot(-4, -2), BuildKnot(-5, -3), BuildKnot(-5, -3)},
		{BuildKnot(-5, -3), BuildKnot(-5, -4), BuildKnot(-5, -4)},
		{BuildKnot(-5, -3), BuildKnot(-7, -3), BuildKnot(-6, -3)},
		{BuildKnot(-5, -3), BuildKnot(-5, -5), BuildKnot(-5, -4)},
		{BuildKnot(-5, -3), BuildKnot(-7, -5), BuildKnot(-6, -4)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			tt.tail.Align(tt.head)
			assert.Equal(t, tt.tail.x, tt.tailAfter.x)
			assert.Equal(t, tt.tail.y, tt.tailAfter.y)
		})
	}
}

func TestRopeKnot_Adjacent(t *testing.T) {
	tests := []struct {
		head     Knot
		tail     Knot
		adjacent bool
		xDiff    int
		yDiff    int
	}{
		{BuildKnot(0, 0), BuildKnot(0, 0), true, 0, 0},

		{BuildKnot(3, 1), BuildKnot(2, 1), true, 1, 0},
		{BuildKnot(4, 2), BuildKnot(3, 1), true, 1, 1},
		{BuildKnot(5, 3), BuildKnot(5, 2), true, 0, 1},
		{BuildKnot(5, 3), BuildKnot(3, 2), false, 2, 1},
		{BuildKnot(5, 3), BuildKnot(5, 1), false, 0, 2},
		{BuildKnot(5, 3), BuildKnot(3, 1), false, 2, 2},

		{BuildKnot(3, 1), BuildKnot(4, 1), true, -1, 0},
		{BuildKnot(4, 2), BuildKnot(5, 3), true, -1, -1},
		{BuildKnot(5, 3), BuildKnot(5, 4), true, 0, -1},
		{BuildKnot(5, 3), BuildKnot(7, 2), false, -2, 1},
		{BuildKnot(5, 3), BuildKnot(5, 5), false, 0, -2},
		{BuildKnot(5, 3), BuildKnot(7, 5), false, -2, -2},

		{BuildKnot(-3, -1), BuildKnot(-2, -1), true, -1, 0},
		{BuildKnot(-4, -2), BuildKnot(-3, -1), true, -1, -1},
		{BuildKnot(-5, -3), BuildKnot(-5, -2), true, 0, -1},
		{BuildKnot(-5, -3), BuildKnot(-3, -2), false, -2, -1},
		{BuildKnot(-5, -3), BuildKnot(-5, -1), false, 0, -2},
		{BuildKnot(-5, -3), BuildKnot(-3, -1), false, -2, -2},

		{BuildKnot(-3, -1), BuildKnot(-4, -1), true, 1, 0},
		{BuildKnot(-4, -2), BuildKnot(-5, -3), true, 1, 1},
		{BuildKnot(-5, -3), BuildKnot(-5, -4), true, 0, 1},
		{BuildKnot(-5, -3), BuildKnot(-7, -2), false, 2, -1},
		{BuildKnot(-5, -3), BuildKnot(-5, -5), false, 0, 2},
		{BuildKnot(-5, -3), BuildKnot(-7, -5), false, 2, 2},

		{BuildKnot(2, 3), BuildKnot(-2, -3), false, 4, 6},
		{BuildKnot(-2, -3), BuildKnot(2, 3), false, -4, -6},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			adj, xDiff, yDiff := tt.tail.adjacent(tt.head)
			assert.Equal(t, tt.adjacent, adj)
			assert.Equal(t, tt.xDiff, xDiff)
			assert.Equal(t, tt.yDiff, yDiff)

			adj, xDiff, yDiff = tt.head.adjacent(tt.tail)
			assert.Equal(t, tt.adjacent, adj)
			assert.Equal(t, tt.xDiff*-1, xDiff)
			assert.Equal(t, tt.yDiff*-1, yDiff)
		})
	}
}

func BuildKnot(x int, y int) Knot {
	k := NewKnot(nil)
	k.x = x
	k.y = y
	return *k
}
