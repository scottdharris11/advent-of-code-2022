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
	solvePart1(input)
	solvePart2(input)
}

var DIRECTIONS = []int32{'>', 'v', '<', '^'}
var MOVES = map[int32]Tuple{
	'>': {x: 1, y: 0},
	'<': {x: -1, y: 0},
	'^': {x: 0, y: -1},
	'v': {x: 0, y: 1},
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	board, instructions := parseInput(lines)
	edge := board.rowEdges[1]
	location := Tuple{x: edge.x, y: 1}
	direction := 0
	for _, i := range instructions {
		switch i {
		case "R":
			direction = (direction + 1) % len(DIRECTIONS)
		case "L":
			direction--
			if direction < 0 {
				direction = len(DIRECTIONS) - 1
			}
		default:
			spaces, err := strconv.Atoi(i)
			if err != nil {
				continue
			}
			location = board.nextPoint(location, direction, spaces)
		}
	}
	ans := (1000 * location.y) + (4 * location.x) + direction
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 1 (%dms): Answer = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

type Tuple struct {
	x int
	y int
}

type Board struct {
	walls    map[Tuple]bool
	spaces   map[Tuple]bool
	rowEdges map[int]Tuple
	colEdges map[int]Tuple
}

func (b Board) nextPoint(point Tuple, dir int, spaces int) Tuple {
	move := MOVES[DIRECTIONS[dir]]
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
	return current
}

func parseInput(lines []string) (Board, []string) {
	b := Board{
		walls:    make(map[Tuple]bool),
		spaces:   make(map[Tuple]bool),
		rowEdges: make(map[int]Tuple),
		colEdges: make(map[int]Tuple),
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
