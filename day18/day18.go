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
	droplet := ParseDroplet(lines)
	ans := droplet.ExposedSides()
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 1 (%dms): Exposed Sides = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	droplet := ParseDroplet(lines)
	ans := droplet.ExteriorSides()
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 2 (%dms): Exterior Sides = %d", end-start, ans)
	return ans
}

func ParseDroplet(lines []string) *Droplet {
	d := &Droplet{maxX: -1, minX: -1, maxY: -1, minY: -1, maxZ: -1, minZ: -1}
	d.cubes = make(map[int]*Cube)
	for _, line := range lines {
		d.AddCube(NewCube(line))
	}
	return d
}

type Droplet struct {
	maxX  int
	minX  int
	maxY  int
	minY  int
	maxZ  int
	minZ  int
	cubes map[int]*Cube
}

func (d *Droplet) AddCube(c *Cube) {
	for _, cube := range d.cubes {
		cube.MarkAdjacent(c)
	}
	d.minX, d.maxX = d.minMax(d.minX, d.maxX, c.x)
	d.minY, d.maxY = d.minMax(d.minY, d.maxY, c.y)
	d.minZ, d.maxZ = d.minMax(d.minZ, d.maxZ, c.z)
	d.cubes[d.cubeKey(c.x, c.y, c.z)] = c
}

func (d *Droplet) ExposedSides() int {
	sides := 0
	for _, c := range d.cubes {
		sides += c.ExposedSides()
	}
	return sides
}

func (d *Droplet) ExteriorSides() int {
	sides := 0
	for _, c := range d.cubes {
		for _, a := range [][]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}} {
			if d.isExternalSide(c, a[0], a[1], a[2]) {
				sides++
			}
		}
	}
	return sides
}

func (d *Droplet) isExternalSide(c *Cube, adjX int, adjY int, adjZ int) bool {
	// look for adjacent cube (if found, not external side)
	key := d.cubeKey(c.x+adjX, c.y+adjY, c.z+adjZ)
	if _, ok := d.cubes[key]; ok {
		return false
	}

	// look for a direct path out
	if d.externalDirectPath(c.x, c.y, c.z, adjX, adjY, adjZ) {
		return true
	}

	// search for path to outside
	search := utils.Search{Searcher: d}
	solution := search.Best(utils.SearchMove{
		Cost:  0,
		State: SearchState{c.x + adjX, c.y + adjY, c.z + adjZ},
	})
	return solution != nil
}

func (d *Droplet) externalDirectPath(x int, y int, z int, adjX int, adjY int, adjZ int) bool {
	for {
		x += adjX
		y += adjY
		z += adjZ
		if _, ok := d.cubes[d.cubeKey(x, y, z)]; ok {
			return false
		}
		if x < d.minX || x > d.maxX || y < d.minY || y > d.maxY || z < d.minZ || z > d.maxZ {
			return true
		}
	}
}

func (d *Droplet) minMax(min int, max int, val int) (int, int) {
	rMax := max
	rMin := min
	if val > max {
		rMax = val
	}
	if min == -1 || val < min {
		rMin = val
	}
	return rMin, rMax
}

func (d *Droplet) cubeKey(x int, y int, z int) int {
	return (z * 100000) + (y * 1000) + x
}

func (d *Droplet) Goal(state interface{}) bool {
	var dState = state.(SearchState)
	return dState.x < d.minX || dState.x > d.maxX ||
		dState.y < d.minY || dState.y > d.maxY ||
		dState.z < d.minZ || dState.z > d.maxZ
}

func (d *Droplet) DistanceFromGoal(state interface{}) int {
	var dState = state.(SearchState)
	xDistance := d.distance(d.minX, d.maxX, dState.x)
	yDistance := d.distance(d.minY, d.maxY, dState.y)
	zDistance := d.distance(d.minZ, d.maxZ, dState.z)
	return xDistance + yDistance + zDistance
}

func (d *Droplet) distance(min int, max int, val int) int {
	distance := val - (min - 1)
	if max-val < distance {
		distance = (max + 1) - val
	}
	return distance
}

func (d *Droplet) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var dState = state.(SearchState)
	var moves []utils.SearchMove
	for _, a := range [][]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}} {
		key := d.cubeKey(dState.x+a[0], dState.y+a[1], dState.z+a[2])
		if _, ok := d.cubes[key]; !ok {
			moves = append(moves, utils.SearchMove{
				Cost:  1,
				State: SearchState{dState.x + a[0], dState.y + a[1], dState.z + a[2]},
			})
		}
	}
	return moves
}

type SearchState struct {
	x int
	y int
	z int
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
