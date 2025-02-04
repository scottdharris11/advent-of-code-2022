package day24

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day24", "day-24-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	rp := parseInput(lines)
	ans := rp.Route()
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 1 (%dms): Answer = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 2 (%dms): Round = %d", end-start, ans)
	return ans
}

type Direction rune

var RIGHT = Direction('>')
var LEFT = Direction('<')
var DOWN = Direction('v')
var UP = Direction('^')

var MOVES = map[Direction]Tuple{
	RIGHT: {x: 1, y: 0},
	LEFT:  {x: -1, y: 0},
	UP:    {x: 0, y: -1},
	DOWN:  {x: 0, y: 1},
}

type Tuple struct {
	x int
	y int
}

type RouteState struct {
	minutes int
	loc     Tuple
}

type BlizzardState struct {
	winds map[Tuple][]Direction
}

func (b BlizzardState) nextMinute(rp RoutePlanner) BlizzardState {
	winds := make(map[Tuple][]Direction)
	for loc, dirs := range b.winds {
		for _, d := range dirs {
			move := MOVES[d]
			l := Tuple{loc.x + move.x, loc.y + move.y}
			if _, wall := rp.walls[l]; wall {
				switch {
				case move.x == 1:
					l.x = 1
				case move.x == -1:
					l.x = rp.max.x - 1
				case move.y == 1:
					l.y = 1
				case move.y == -1:
					l.y = rp.max.y - 1
				}
			}
			nd := winds[l]
			nd = append(nd, d)
			winds[l] = nd
		}
	}
	return BlizzardState{winds: winds}
}

type RoutePlanner struct {
	start     RouteState
	goal      Tuple
	max       Tuple
	walls     map[Tuple]bool
	blizzards map[int]BlizzardState
}

func (r *RoutePlanner) Route() int {
	search := utils.Search{Searcher: r}
	solution := search.Best(utils.SearchMove{
		Cost:  0,
		State: r.start,
	})
	if solution == nil {
		return -1
	}
	return solution.Cost
}

func (r *RoutePlanner) Goal(state interface{}) bool {
	var routeState = state.(RouteState)
	return routeState.loc.x == r.goal.x && routeState.loc.y == r.goal.y
}

func (r *RoutePlanner) DistanceFromGoal(state interface{}) int {
	var routeState = state.(RouteState)
	xDistance := routeState.loc.x - r.goal.x
	if xDistance < 0 {
		xDistance *= -1
	}
	yDistance := routeState.loc.y - r.goal.y
	if yDistance < 0 {
		yDistance *= -1
	}
	return xDistance + yDistance
}

func (r *RoutePlanner) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var routeState = state.(RouteState)
	minute := routeState.minutes + 1
	bs := r.blizzardState(minute)
	var moves []utils.SearchMove
	for _, m := range []Tuple{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		nl := Tuple{routeState.loc.x + m.x, routeState.loc.y + m.y}
		if nl.x < 0 || nl.y < 0 || nl.x > r.max.x || nl.y > r.max.y {
			continue
		}
		if _, wall := r.walls[nl]; wall {
			continue
		}
		if _, ok := bs.winds[nl]; !ok {
			moves = append(moves, utils.SearchMove{Cost: 1, State: RouteState{minutes: minute, loc: nl}})
		}
	}
	if _, ok := bs.winds[routeState.loc]; !ok {
		moves = append(moves, utils.SearchMove{Cost: 1, State: RouteState{minutes: minute, loc: routeState.loc}})
	}
	return moves
}

func (r *RoutePlanner) blizzardState(minute int) BlizzardState {
	bs, ok := r.blizzards[minute]
	if !ok {
		prev := r.blizzards[minute-1]
		bs = prev.nextMinute(*r)
		r.blizzards[minute] = bs
	}
	return bs
}

func parseInput(lines []string) RoutePlanner {
	walls := make(map[Tuple]bool)
	locations := make(map[Tuple][]Direction)
	var start *Tuple
	var goal Tuple
	var max Tuple
	for y, line := range lines {
		for x, col := range line {
			loc := Tuple{x, y}
			switch col {
			case '#':
				walls[loc] = true
			case '>', '<', 'v', '^':
				locations[loc] = []Direction{Direction(col)}
			case '.':
				if start == nil {
					start = &Tuple{loc.x, loc.y}
				}
				goal = loc
			}
			max = loc
		}
	}
	return RoutePlanner{
		start:     RouteState{minutes: 0, loc: *start},
		goal:      goal,
		max:       max,
		walls:     walls,
		blizzards: map[int]BlizzardState{0: {winds: locations}},
	}
}
