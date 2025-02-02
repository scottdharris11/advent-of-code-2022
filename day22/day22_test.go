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
	assert.Equal(t, 6032, solvePart1(testInput, sampleCubeSides()))
	assert.Equal(t, 133174, solvePart1(utils.ReadLines("day22", "day-22-input.txt"), cubeSides()))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 5031, solvePart2(testInput, sampleCubeSides()))
	assert.Equal(t, 15410, solvePart2(utils.ReadLines("day22", "day-22-input.txt"), cubeSides()))
}

func sampleCubeSides() map[int]CubeSide {
	sides := make(map[int]CubeSide)
	side1 := CubeSide{id: 1, topLeft: Tuple{x: 9, y: 1}, bottomRight: Tuple{x: 12, y: 4}, links: make(map[DIRECTION]CubeSideLinkage)}
	side2 := CubeSide{id: 2, topLeft: Tuple{x: 1, y: 5}, bottomRight: Tuple{x: 4, y: 8}, links: make(map[DIRECTION]CubeSideLinkage)}
	side3 := CubeSide{id: 3, topLeft: Tuple{x: 5, y: 5}, bottomRight: Tuple{x: 8, y: 8}, links: make(map[DIRECTION]CubeSideLinkage)}
	side4 := CubeSide{id: 4, topLeft: Tuple{x: 9, y: 5}, bottomRight: Tuple{x: 12, y: 8}, links: make(map[DIRECTION]CubeSideLinkage)}
	side5 := CubeSide{id: 5, topLeft: Tuple{x: 9, y: 9}, bottomRight: Tuple{x: 12, y: 12}, links: make(map[DIRECTION]CubeSideLinkage)}
	side6 := CubeSide{id: 6, topLeft: Tuple{x: 13, y: 9}, bottomRight: Tuple{x: 16, y: 12}, links: make(map[DIRECTION]CubeSideLinkage)}

	side1.links[DOWN] = CubeSideLinkage{to: side4, toDir: DOWN}
	side1.links[UP] = CubeSideLinkage{to: side2, toDir: DOWN}
	side1.links[LEFT] = CubeSideLinkage{to: side3, toDir: DOWN}
	side1.links[RIGHT] = CubeSideLinkage{to: side6, toDir: LEFT}

	side2.links[DOWN] = CubeSideLinkage{to: side5, toDir: UP}
	side2.links[UP] = CubeSideLinkage{to: side1, toDir: DOWN}
	side2.links[LEFT] = CubeSideLinkage{to: side6, toDir: UP}
	side2.links[RIGHT] = CubeSideLinkage{to: side3, toDir: RIGHT}

	side3.links[DOWN] = CubeSideLinkage{to: side5, toDir: RIGHT}
	side3.links[UP] = CubeSideLinkage{to: side1, toDir: RIGHT}
	side3.links[LEFT] = CubeSideLinkage{to: side2, toDir: LEFT}
	side3.links[RIGHT] = CubeSideLinkage{to: side4, toDir: RIGHT}

	side4.links[DOWN] = CubeSideLinkage{to: side5, toDir: DOWN}
	side4.links[UP] = CubeSideLinkage{to: side1, toDir: UP}
	side4.links[LEFT] = CubeSideLinkage{to: side3, toDir: LEFT}
	side4.links[RIGHT] = CubeSideLinkage{to: side6, toDir: DOWN}

	side5.links[DOWN] = CubeSideLinkage{to: side2, toDir: UP}
	side5.links[UP] = CubeSideLinkage{to: side4, toDir: UP}
	side5.links[LEFT] = CubeSideLinkage{to: side3, toDir: UP}
	side5.links[RIGHT] = CubeSideLinkage{to: side6, toDir: RIGHT}

	side6.links[DOWN] = CubeSideLinkage{to: side2, toDir: RIGHT}
	side6.links[UP] = CubeSideLinkage{to: side4, toDir: LEFT}
	side6.links[LEFT] = CubeSideLinkage{to: side5, toDir: LEFT}
	side6.links[RIGHT] = CubeSideLinkage{to: side1, toDir: LEFT}

	sides[1] = side1
	sides[2] = side2
	sides[3] = side3
	sides[4] = side4
	sides[5] = side5
	sides[6] = side6
	return sides
}
