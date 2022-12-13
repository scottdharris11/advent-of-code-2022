package day12

import (
	"fmt"
	"log"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day12", "day-12-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(grid []string) int {
	start := time.Now().UnixMilli()
	rp := NewRoutePlanner(grid)
	ans := rp.Route()
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 1 (%dms): Shortest Route = %d", end-start, ans)
	return ans
}

func solvePart2(grid []string) int {
	start := time.Now().UnixMilli()
	ans := len(grid)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 2 (%dms): Shortest Route = %d", end-start, ans)
	return ans
}

func NewRoutePlanner(grid []string) *RoutePlanner {
	r := &RoutePlanner{}
	for y, row := range grid {
		nRow := row
		if idx := strings.IndexRune(row, 'S'); idx >= 0 {
			r.startLocation = Position{idx, y}
			nRow = strings.Join([]string{nRow[:idx], "a", nRow[idx+1:]}, "")
		}
		if idx := strings.IndexRune(row, 'E'); idx >= 0 {
			r.goalLocation = Position{idx, y}
			nRow = strings.Join([]string{nRow[:idx], "z", nRow[idx+1:]}, "")
		}
		r.grid = append(r.grid, nRow)
	}
	return r
}

type RoutePlanner struct {
	grid          []string
	startLocation Position
	goalLocation  Position
}

func (r *RoutePlanner) Route() int {
	search := utils.Search{Searcher: r}
	solution := search.Best(utils.SearchMove{
		Cost: 0,
		State: RouteState{
			currentX: r.startLocation.x,
			currentY: r.startLocation.y,
		},
	})
	return solution.Cost
}

func (r *RoutePlanner) Goal(state interface{}) bool {
	var routeState = state.(RouteState)
	return routeState.currentX == r.goalLocation.x && routeState.currentY == r.goalLocation.y
}

func (r *RoutePlanner) DistanceFromGoal(state interface{}) int {
	var routeState = state.(RouteState)
	xDistance := routeState.currentX - r.goalLocation.x
	if xDistance < 0 {
		xDistance *= -1
	}
	yDistance := routeState.currentY - r.goalLocation.y
	if yDistance < 0 {
		yDistance *= -1
	}
	return xDistance + yDistance
}

func (r *RoutePlanner) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var routeState = state.(RouteState)
	current := Position{x: routeState.currentX, y: routeState.currentY}

	e := rune(r.grid[current.y][current.x])
	up := Position{current.x, current.y - 1}
	down := Position{current.x, current.y + 1}
	left := Position{current.x - 1, current.y}
	right := Position{current.x + 1, current.y}
	var moves []utils.SearchMove
	for _, pos := range []Position{up, down, left, right} {
		if !r.available(e, pos) {
			continue
		}
		move := utils.SearchMove{
			Cost: 1,
			State: RouteState{
				currentX: pos.x,
				currentY: pos.y,
			},
		}
		moves = append(moves, move)
	}
	return moves
}

func (r *RoutePlanner) available(currElevation rune, pos Position) bool {
	ok, e := r.elevation(pos)
	if !ok {
		return false
	}
	return int(e)-int(currElevation) <= 1
}

func (r *RoutePlanner) elevation(pos Position) (bool, rune) {
	if pos.y < 0 || pos.y >= len(r.grid) {
		return false, ' '
	}
	if pos.x < 0 || pos.x >= len(r.grid[0]) {
		return false, ' '
	}
	return true, rune(r.grid[pos.y][pos.x])
}

type RouteState struct {
	currentX int
	currentY int
}

type Position struct {
	x int
	y int
}

func (p *Position) encode() string {
	return fmt.Sprintf(`"%d::%d"`, p.x, p.y)
}
