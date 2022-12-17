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
	cave := parseCave(lines, false)
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
	cave := parseCave(lines, true)
	ans := 0
	for {
		if cave.DropSand() {
			break
		}
		ans++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 2 (%dms): Sand Drops = %d", end-start, ans)
	return ans
}

func parseCave(lines []string, floor bool) Cave {
	var rocks []Rock
	for _, line := range lines {
		rocks = append(rocks, NewRock(line))
	}
	return NewCave(rocks, floor)
}

var Stone = '#'
var Sand = 'o'
var Empty = ' '

type Cave struct {
	xOffset int
	maxX    int
	maxY    int
	grid    [][]rune
	floor   bool
}

func (c *Cave) DropSand() bool {
	sandX := 500 - c.xOffset
	sandY := 0
	for {
		if c.grid[sandY][sandX] == Sand {
			return true
		}
		if off, _ := c.offGrid(sandX, sandY+1); off {
			return true
		}
		switch c.at(sandX, sandY+1) {
		case Stone, Sand:
			off, adjusted := c.offGrid(sandX-1, sandY-1)
			if off {
				return true
			}
			if adjusted {
				sandX++
			}
			if c.canMoveTo(sandX-1, sandY+1) {
				sandX--
				sandY++
				break
			}

			if off, _ = c.offGrid(sandX+1, sandY+1); off {
				return true
			}
			if c.canMoveTo(sandX+1, sandY+1) {
				sandX++
				sandY++
				break
			}

			c.grid[sandY][sandX] = Sand
			return false
		default:
			sandY++
		}
	}
}

func (c *Cave) offGrid(x int, y int) (bool, bool) {
	if c.floor {
		switch {
		case x < 0:
			for i := 0; i < len(c.grid); i++ {
				c.grid[i] = append([]rune{Empty}, c.grid[i]...)
			}
			c.xOffset--
			return false, true
		case x > c.maxX:
			for i := 0; i < len(c.grid); i++ {
				c.grid[i] = append(c.grid[i], Empty)
			}
			return false, false
		}
	}
	return x < 0 || x > c.maxX || y > c.maxY, false
}

func (c *Cave) canMoveTo(x int, y int) bool {
	return c.at(x, y) == Empty
}

func (c *Cave) at(x int, y int) rune {
	if c.floor && y == c.maxY {
		return Stone
	}
	return c.grid[y][x]
}

func NewCave(rocks []Rock, floor bool) Cave {
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

	if floor {
		maxY += 2
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
		floor:   floor,
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
