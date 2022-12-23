package day18

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day18", "day-18-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	cubes := ParseDroplet(lines)
	ans := 0
	for _, c := range cubes {
		ans += c.ExposedSides()
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 1 (%dms): Exposed Sides = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

func ParseDroplet(lines []string) []*Cube {
	var cubes []*Cube
	for _, line := range lines {
		cubes = append(cubes, NewCube(line))
	}
	for i := 0; i < len(cubes); i++ {
		for j := 0; j < len(cubes); j++ {
			if i == j {
				continue
			}
			cubes[i].MarkAdjacent(cubes[j])
		}
	}
	return cubes
}

func NewCube(line string) *Cube {
	i := utils.ReadIntegersFromLine(line, ",")
	return &Cube{x: i[0], y: i[1], z: i[2]}
}

type Cube struct {
	x        int
	y        int
	z        int
	adjacent [6]*Cube
}

func (c *Cube) MarkAdjacent(c2 *Cube) {
	switch {
	case c.z == c2.z && c.y == c2.y && c.x+1 == c2.x:
		c.adjacent[1] = c2
		c2.adjacent[0] = c
	case c.z == c2.z && c.y == c2.y && c.x-1 == c2.x:
		c.adjacent[0] = c2
		c2.adjacent[1] = c
	case c.z == c2.z && c.x == c2.x && c.y+1 == c2.y:
		c.adjacent[3] = c2
		c2.adjacent[2] = c
	case c.z == c2.z && c.x == c2.x && c.y-1 == c2.y:
		c.adjacent[2] = c2
		c2.adjacent[3] = c
	case c.x == c2.x && c.y == c2.y && c.z+1 == c2.z:
		c.adjacent[5] = c2
		c2.adjacent[4] = c
	case c.x == c2.x && c.y == c2.y && c.z-1 == c2.z:
		c.adjacent[4] = c2
		c2.adjacent[5] = c
	}
}

func (c *Cube) ExposedSides() int {
	e := 0
	for _, c := range c.adjacent {
		if c == nil {
			e++
		}
	}
	return e
}
