package day14

import (
	"strings"

	"advent-of-code-2022/utils"
)

var Stone = '#'

type Cave struct {
	xOffset int
	grid    [][]rune
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
