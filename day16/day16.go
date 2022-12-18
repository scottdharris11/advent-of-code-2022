package day16

import (
	"log"
	"regexp"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day16", "day-16-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	planner := NewPlanner(lines)
	ans := planner.Route()
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 1 (%dms): Max Released = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

type Valve struct {
	id      string
	flow    int
	tunnels []string
}

func (v Valve) PressureRelief(openMinutes int) int {
	return v.flow * openMinutes
}

var valveParse = regexp.MustCompile(`Valve ([A-Z]+) has flow rate=([\d]+); tunnels* leads* to valves* ([A-Z ,]+)`)

func NewValve(s string) *Valve {
	match := valveParse.FindStringSubmatch(s)
	if match == nil {
		return nil
	}
	return &Valve{
		id:      match[1],
		flow:    utils.Number(match[2]),
		tunnels: strings.Split(match[3], ", "),
	}
}

func NewPlanner(lines []string) Planner {
	p := Planner{}
	p.valves = make(map[string]Valve)
	for _, line := range lines {
		v := NewValve(line)
		if v.flow > 0 {
			p.reliefValves++
		}
		p.valves[v.id] = *v
	}
	return p
}

type Planner struct {
	valves       map[string]Valve
	reliefValves int
}

func (p *Planner) Route() int {
	search := utils.Search{Searcher: p}
	solution := search.Best(utils.SearchMove{
		Cost: 0,
		State: State{
			location:     "AA",
			minRemaining: 30,
		},
	})
	if solution == nil {
		return -1
	}
	return solution.Path[len(solution.Path)-1].(State).relief
}

func (p *Planner) Goal(state interface{}) bool {
	var pState = state.(State)
	return pState.opened == p.reliefValves || pState.minRemaining == 0
}

func (p *Planner) DistanceFromGoal(state interface{}) int {
	var pState = state.(State)
	return p.reliefValves - pState.opened + pState.minRemaining
}

func (p *Planner) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var pState = state.(State)
	v := p.valves[pState.location]

	// Add move to open valve in current tunnel if not opened
	var moves []utils.SearchMove
	minRemaining := pState.minRemaining - 1

	if v.flow > 0 && !strings.Contains(pState.openValves, v.id) {
		nOpen := pState.openValves
		if nOpen != "" {
			nOpen += "-"
		}
		nOpen += v.id
		nState := State{
			location:     v.id,
			opened:       pState.opened + 1,
			openValves:   nOpen,
			minRemaining: minRemaining,
			relief:       pState.relief + v.PressureRelief(minRemaining),
		}
		moves = append(moves, utils.SearchMove{
			Cost:  p.moveCost(nState),
			State: nState,
		})
	}

	// Add moves into any of the connecting tunnels
	for _, t := range v.tunnels {
		nState := State{
			location:     t,
			opened:       pState.opened,
			openValves:   pState.openValves,
			minRemaining: minRemaining,
			relief:       pState.relief,
		}
		moves = append(moves, utils.SearchMove{
			Cost:  p.moveCost(nState),
			State: nState,
		})
	}
	return moves
}

func (p *Planner) moveCost(state State) int {
	cost := 0
	for _, v := range p.valves {
		if v.flow == 0 || strings.Contains(state.openValves, v.id) {
			continue
		}
		cost += v.PressureRelief(state.minRemaining)
	}
	return cost
}

type State struct {
	location     string
	opened       int
	openValves   string
	minRemaining int
	relief       int
}
