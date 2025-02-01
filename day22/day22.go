package day22

import (
	"log"
	"regexp"
	"strconv"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day22", "day-22-input.txt")
	solvePart1(input, cubeSides())
	solvePart2(input, cubeSides())
}

type DIRECTION int32

var RIGHT = DIRECTION('>')
var LEFT = DIRECTION('<')
var DOWN = DIRECTION('v')
var UP = DIRECTION('^')
var DIRECTIONS = []DIRECTION{RIGHT, DOWN, LEFT, UP}
var MOVES = map[DIRECTION]Tuple{
	RIGHT: {x: 1, y: 0},
	LEFT:  {x: -1, y: 0},
	UP:    {x: 0, y: -1},
	DOWN:  {x: 0, y: 1},
}

func solvePart1(lines []string, sides map[int]CubeSide) int {
	start := time.Now().UnixMilli()
	board, instructions := parseInput(lines, sides, false)
	location, direction := followInstructions(board, instructions)
	ans := password(location, direction)
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 1 (%dms): Password = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string, sides map[int]CubeSide) int {
	start := time.Now().UnixMilli()
	board, instructions := parseInput(lines, sides, true)
	location, direction := followInstructions(board, instructions)
	ans := password(location, direction)
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

type Tuple struct {
	x int
	y int
}

type CubeSide struct {
	id          int
	topLeft     Tuple
	bottomRight Tuple
	links       map[DIRECTION]CubeSideLinkage
}

func (c CubeSide) pointOn(point Tuple) bool {
	if point.x < c.topLeft.x || point.x > c.bottomRight.x {
		return false
	}
	if point.y < c.topLeft.y || point.y > c.bottomRight.y {
		return false
	}
	return true
}

func (c CubeSide) onEdge(point Tuple, dir DIRECTION) bool {
	if point.x == c.topLeft.x && dir == LEFT {
		return true
	}
	if point.x == c.bottomRight.x && dir == RIGHT {
		return true
	}
	if point.y == c.topLeft.y && dir == UP {
		return true
	}
	if point.y == c.bottomRight.y && dir == DOWN {
		return true
	}
	return false
}

func (c CubeSide) nextPoint(point Tuple, dir DIRECTION) (Tuple, CubeSide, DIRECTION) {
	if !c.onEdge(point, dir) {
		// moving within cube
		move := MOVES[dir]
		return Tuple{point.x + move.x, point.y + move.y}, c, dir
	}

	link := c.links[dir]
	if dir == link.toDir {
		// wrapping, but moving in same direction
		move := MOVES[dir]
		return Tuple{point.x + move.x, point.y + move.y}, link.to, dir
	}

	topOffset := point.y - c.topLeft.y
	rightOffset := point.x - c.topLeft.x
	var next Tuple
	switch {
	case dir == UP && link.toDir == DOWN:
		next = Tuple{link.to.bottomRight.x - rightOffset, link.to.topLeft.y}
	case dir == UP && link.toDir == LEFT:
		next = Tuple{link.to.bottomRight.x, link.to.bottomRight.y - rightOffset}
	case dir == UP && link.toDir == RIGHT:
		next = Tuple{link.to.topLeft.x, link.to.topLeft.y + rightOffset}
	case dir == DOWN && link.toDir == UP:
		next = Tuple{link.to.bottomRight.x - rightOffset, link.to.bottomRight.y}
	case dir == DOWN && link.toDir == LEFT:
		next = Tuple{link.to.bottomRight.x, link.to.topLeft.y + rightOffset}
	case dir == DOWN && link.toDir == RIGHT:
		next = Tuple{link.to.topLeft.x, link.to.bottomRight.y - rightOffset}
	case dir == LEFT && link.toDir == DOWN:
		next = Tuple{link.to.topLeft.x + topOffset, link.to.topLeft.y}
	case dir == LEFT && link.toDir == UP:
		next = Tuple{link.to.bottomRight.x - topOffset, link.to.bottomRight.y}
	case dir == LEFT && link.toDir == RIGHT:
		next = Tuple{link.to.topLeft.x, link.to.bottomRight.y - topOffset}
	case dir == RIGHT && link.toDir == DOWN:
		next = Tuple{link.to.bottomRight.x - topOffset, link.to.topLeft.y}
	case dir == RIGHT && link.toDir == UP:
		next = Tuple{link.to.bottomRight.x - topOffset, link.to.bottomRight.y}
	case dir == RIGHT && link.toDir == LEFT:
		next = Tuple{link.to.bottomRight.x, link.to.bottomRight.y - topOffset}
	}
	return next, link.to, link.toDir
}

type CubeSideLinkage struct {
	to    CubeSide
	toDir DIRECTION
}

type Board struct {
	walls      map[Tuple]bool
	spaces     map[Tuple]bool
	rowEdges   map[int]Tuple
	colEdges   map[int]Tuple
	cubeSides  map[int]CubeSide
	pointSides map[Tuple]int
	cube       bool
}

func (b Board) nextPoint(point Tuple, dir DIRECTION, spaces int) (Tuple, DIRECTION) {
	if b.cube {
		return b.nextCubePoint(point, dir, spaces)
	}
	move := MOVES[dir]
	current := point
	for i := 0; i < spaces; i++ {
		np := Tuple{current.x + move.x, current.y + move.y}
		if _, ok := b.walls[np]; ok {
			break
		}
		if _, ok := b.spaces[np]; !ok {
			switch {
			case move.x == 1:
				edge := b.rowEdges[np.y]
				np.x = edge.x
			case move.x == -1:
				edge := b.rowEdges[np.y]
				np.x = edge.y
			case move.y == 1:
				edge := b.colEdges[np.x]
				np.y = edge.x
			case move.y == -1:
				edge := b.colEdges[np.x]
				np.y = edge.y
			}
			if _, ok = b.walls[np]; ok {
				break
			}
		}
		current = np
	}
	return current, dir
}

func (b Board) nextCubePoint(point Tuple, dir DIRECTION, spaces int) (Tuple, DIRECTION) {
	d := dir
	c := point
	cs := b.cubeSides[b.pointSides[point]]
	for i := 0; i < spaces; i++ {
		np, ns, nd := cs.nextPoint(c, d)
		if _, ok := b.walls[np]; ok {
			break
		}
		c = np
		d = nd
		cs = ns
	}
	return c, d
}

func parseInput(lines []string, sides map[int]CubeSide, cube bool) (Board, []string) {
	b := Board{
		cube:       cube,
		walls:      make(map[Tuple]bool),
		spaces:     make(map[Tuple]bool),
		rowEdges:   make(map[int]Tuple),
		colEdges:   make(map[int]Tuple),
		cubeSides:  sides,
		pointSides: make(map[Tuple]int),
	}
	var idx int
	for idx = 0; idx < len(lines); idx++ {
		if lines[idx] == "" {
			idx++
			break
		}
		for x, c := range lines[idx] {
			point := Tuple{x + 1, idx + 1}
			switch c {
			case '#':
				b.walls[point] = true
			case '.':
				b.spaces[point] = true
			default:
				continue
			}

			for id, cs := range sides {
				if cs.pointOn(point) {
					b.pointSides[point] = id
				}
			}
			if _, ok := b.pointSides[point]; !ok {
				log.Printf("Didn't match cube side: %+v", point)
			}

			edges, ok := b.rowEdges[point.y]
			if !ok {
				edges = Tuple{x: point.x, y: point.x}
			}
			if point.x < edges.x {
				edges.x = point.x
			}
			if point.x > edges.y {
				edges.y = point.x
			}
			b.rowEdges[point.y] = edges

			edges, ok = b.colEdges[point.x]
			if !ok {
				edges = Tuple{x: point.y, y: point.y}
			}
			if point.y < edges.x {
				edges.x = point.y
			}
			if point.y > edges.y {
				edges.y = point.y
			}
			b.colEdges[point.x] = edges
		}
	}
	return b, parseInstructions(lines[idx])
}

func parseInstructions(line string) []string {
	re := regexp.MustCompile(`([\d]+)([RL])*`)
	matches := re.FindAllStringSubmatch(line, -1)
	var instructions []string
	for _, match := range matches {
		for i := 1; i < len(match); i++ {
			if match[i] == "" {
				continue
			}
			instructions = append(instructions, match[i])
		}
	}
	return instructions
}

func followInstructions(board Board, instructions []string) (Tuple, DIRECTION) {
	edge := board.rowEdges[1]
	location := Tuple{x: edge.x, y: 1}
	direction := RIGHT
	for _, i := range instructions {
		switch i {
		case "R":
			direction = turn(direction, 1)
		case "L":
			direction = turn(direction, -1)
		default:
			spaces, err := strconv.Atoi(i)
			if err != nil {
				continue
			}
			location, direction = board.nextPoint(location, direction, spaces)
		}
	}
	return location, direction
}

func turn(current DIRECTION, offset int) DIRECTION {
	idx := directionOffset(current) + offset
	if idx < 0 {
		idx = len(DIRECTIONS) - 1
	}
	return DIRECTIONS[idx%len(DIRECTIONS)]
}

func directionOffset(dir DIRECTION) int {
	for i, d := range DIRECTIONS {
		if d == dir {
			return i
		}
	}
	return -1
}

func password(location Tuple, dir DIRECTION) int {
	return (1000 * location.y) + (4 * location.x) + directionOffset(dir)
}

// probably could make this work for any solution, but not doing that now and just preconfiguring
// the sides based on my input.  might come back and make this more generic.
func cubeSides() map[int]CubeSide {
	sides := make(map[int]CubeSide)
	side1 := CubeSide{id: 1, topLeft: Tuple{x: 51, y: 1}, bottomRight: Tuple{x: 100, y: 50}, links: make(map[DIRECTION]CubeSideLinkage)}
	side2 := CubeSide{id: 2, topLeft: Tuple{x: 101, y: 1}, bottomRight: Tuple{x: 150, y: 50}, links: make(map[DIRECTION]CubeSideLinkage)}
	side3 := CubeSide{id: 3, topLeft: Tuple{x: 51, y: 51}, bottomRight: Tuple{x: 100, y: 100}, links: make(map[DIRECTION]CubeSideLinkage)}
	side4 := CubeSide{id: 4, topLeft: Tuple{x: 1, y: 101}, bottomRight: Tuple{x: 50, y: 150}, links: make(map[DIRECTION]CubeSideLinkage)}
	side5 := CubeSide{id: 5, topLeft: Tuple{x: 51, y: 101}, bottomRight: Tuple{x: 100, y: 150}, links: make(map[DIRECTION]CubeSideLinkage)}
	side6 := CubeSide{id: 6, topLeft: Tuple{x: 1, y: 151}, bottomRight: Tuple{x: 50, y: 200}, links: make(map[DIRECTION]CubeSideLinkage)}

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
