package day14

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day14", "day-14-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	cave := parseCave(lines)
	ans := 0
	for {
		if cave.DropSand() {
			break
		}
		ans++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 1 (%dms): Sand Drops = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

func parseCave(lines []string) Cave {
	var rocks []Rock
	for _, line := range lines {
		rocks = append(rocks, NewRock(line))
	}
	return NewCave(rocks)
}

var Stone = '#'
var Sand = 'o'
var Empty = ' '

type Cave struct {
	xOffset int
	maxX    int
	maxY    int
	grid    [][]rune
}

func (c *Cave) DropSand() bool {
	sandX := 500 - c.xOffset
	sandY := 0
	for {
		if c.offGrid(sandX, sandY+1) {
			return true
		}
		switch c.grid[sandY+1][sandX] {
		case Stone, Sand:
			switch {
			case c.offGrid(sandX-1, sandY+1):
				return true
			case c.canMoveTo(sandX-1, sandY+1):
				sandX--
				sandY++
			case c.offGrid(sandX+1, sandY+1):
				return true
			case c.canMoveTo(sandX+1, sandY+1):
				sandX++
				sandY++
			default:
				c.grid[sandY][sandX] = Sand
				return false
			}
		default:
			sandY++
		}
	}
}

func (c *Cave) offGrid(x int, y int) bool {
	return x < 0 || x > c.maxX || y > c.maxY
}

func (c *Cave) canMoveTo(x int, y int) bool {
	return c.grid[y][x] == Empty
}

func NewCave(rocks []Rock) Cave {
	minX := -1
	maxX := 0
	maxY := 0
	for _, r := range rocks {
		for _, p := range r.points {
			if minX == -1 || p.x < minX {
				minX = p.x
			}
			if p.x > maxX {
				maxX = p.x
			}
			if p.y > maxY {
				maxY = p.y
			}
		}
	}

	var grid [][]rune
	width := maxX - minX + 1
	for i := 0; i <= maxY; i++ {
		grid = append(grid, make([]rune, width))
		for j := 0; j < width; j++ {
			grid[i][j] = ' '
		}
	}

	for _, r := range rocks {
		start := r.points[0]
		for i := 1; i < len(r.points); i++ {
			end := r.points[i]
			xDiff := start.x - end.x
			yDiff := start.y - end.y
			switch {
			case xDiff > 0:
				for x := 0; x <= xDiff; x++ {
					grid[start.y][start.x-minX-x] = Stone
				}
			case xDiff < 0:
				for x := 0; x >= xDiff; x-- {
					grid[start.y][start.x-minX-x] = Stone
				}
			case yDiff > 0:
				for y := 0; y <= yDiff; y++ {
					grid[start.y-y][start.x-minX] = Stone
				}
			case yDiff < 0:
				for y := 0; y >= yDiff; y-- {
					grid[start.y-y][start.x-minX] = Stone
				}
			}
			start = end
		}
	}

	return Cave{
		xOffset: minX,
		maxX:    maxX - minX,
		maxY:    maxY,
		grid:    grid,
	}
}

type Rock struct {
	points []Point
}

func NewRock(s string) Rock {
	points := strings.Split(s, " -> ")
	r := Rock{}
	for _, point := range points {
		r.points = append(r.points, NewPoint(point))
	}
	return r
}

type Point struct {
	x int
	y int
}

func NewPoint(s string) Point {
	split := strings.Split(s, ",")
	return Point{
		x: utils.Number(split[0]),
		y: utils.Number(split[1]),
	}
}
