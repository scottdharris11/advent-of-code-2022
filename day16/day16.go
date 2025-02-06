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
	valves := parseValves(lines)
	vRoutes := valveRoutes(valves)
	rValves := reliefValves(valves)
	startState := RouteState{loc: valves["AA"], remain: 30}
	ans, route := maxReliefRoute(startState, vRoutes, rValves)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 1 (%dms): Max Released via Route(%s) = %d", end-start, route, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	valves := parseValves(lines)
	vRoutes := valveRoutes(valves)
	rValves := reliefValves(valves)
	meStartState := RouteState{loc: valves["AA"], remain: 26}
	elStartState := RouteState{loc: valves["AA"], remain: 26}
	ans, routes := maxReliefDualRoute(meStartState, elStartState, vRoutes, rValves)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 2 (%dms): Max Released via Routes(%s, %s)= %d", end-start, routes[0], routes[1], ans)
	return ans
}

type Valve struct {
	id      string
	flow    int
	tunnels []*Valve
}

var valveParse = regexp.MustCompile(`Valve ([A-Z]+) has flow rate=([\d]+); tunnels* leads* to valves* ([A-Z ,]+)`)

func parseValves(lines []string) map[string]*Valve {
	// build valves
	valves := make(map[string]*Valve)
	for _, line := range lines {
		match := valveParse.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		v := &Valve{
			id:   match[1],
			flow: utils.Number(match[2]),
		}
		valves[v.id] = v
	}

	// link valves
	for _, line := range lines {
		match := valveParse.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		v := valves[match[1]]
		for _, connection := range strings.Split(match[3], ", ") {
			v.tunnels = append(v.tunnels, valves[connection])
		}
	}
	return valves
}

func valveRoutes(valves map[string]*Valve) map[string]int {
	l := len(valves)
	vs := make([]*Valve, 0, l)
	for _, valve := range valves {
		vs = append(vs, valve)
	}
	routes := make(map[string]int)
	for a := 0; a < l; a++ {
		for b := a + 1; b < l; b++ {
			m := minutesToValve(vs[a], vs[b].id, 0, -1, "")
			routes[vs[a].id+"-"+vs[b].id] = m
			routes[vs[b].id+"-"+vs[a].id] = m
		}
	}
	return routes
}

func minutesToValve(v *Valve, goal string, minutes int, best int, route string) int {
	if v.id == goal {
		return minutes
	}
	if minutes >= 28 {
		return -1
	}
	if best > 0 && minutes >= best-1 {
		return -1
	}
	if strings.Contains(route, v.id) {
		return -1
	}
	route += v.id
	for _, next := range v.tunnels {
		m := minutesToValve(next, goal, minutes+1, best, route)
		if m > 0 && (best == -1 || m < best) {
			best = m
		}
	}
	return best
}

func reliefValves(valves map[string]*Valve) []*Valve {
	var rValves []*Valve
	for _, valve := range valves {
		if valve.flow > 0 {
			rValves = append(rValves, valve)
		}
	}
	return rValves
}

type RouteState struct {
	loc    *Valve
	remain int
	route  string
	relief int
}

func maxReliefRoute(current RouteState, routes map[string]int, valves []*Valve) (int, string) {
	next := possibleNextValves(current.loc, current.remain, routes, valves, current.route)
	if len(next) == 0 {
		return current.relief, current.route
	}
	bestNext := 0
	bestRoute := ""
	for _, v := range next {
		br, vr := maxReliefRoute(nextState(current, v, routes), routes, valves)
		if br > bestNext {
			bestNext = br
			bestRoute = vr
		}
	}
	return bestNext, bestRoute
}

func maxReliefDualRoute(me RouteState, el RouteState, routes map[string]int, valves []*Valve) (int, []string) {
	// compute next possible valves for both me and elephant
	mNext := possibleNextValves(me.loc, me.remain, routes, valves, me.route+el.route)
	eNext := possibleNextValves(el.loc, el.remain, routes, valves, me.route+el.route)

	// if no more possible, return
	if len(mNext) == 0 && len(eNext) == 0 {
		return me.relief + el.relief, []string{me.route, el.route}
	}

	// adjust (add nil) so that values are equal
	for diff := len(mNext) - len(eNext); diff > 0; diff-- {
		eNext = append(eNext, nil)
	}
	for diff := len(eNext) - len(mNext); diff > 0; diff-- {
		mNext = append(mNext, nil)
	}

	// cycle through all the possible next state move to find best combo
	bestNext := 0
	var bestRoute []string
	for _, m := range mNext {
		for _, e := range eNext {
			if m == e {
				continue
			}
			mState := nextState(me, m, routes)
			eState := nextState(el, e, routes)
			br, rr := maxReliefDualRoute(mState, eState, routes, valves)
			if br > bestNext {
				bestNext = br
				bestRoute = rr
			}
		}
	}
	return bestNext, bestRoute
}

func possibleNextValves(current *Valve, remain int, routes map[string]int, valves []*Valve, route string) []*Valve {
	var next []*Valve
	if current == nil {
		return next
	}
	for _, v := range valves {
		if strings.Contains(route, v.id) {
			continue
		}
		mRemain := remain - (routes[current.id+"-"+v.id] + 1)
		if mRemain <= 0 {
			continue
		}
		next = append(next, v)
	}
	return next
}

func nextState(current RouteState, next *Valve, routes map[string]int) RouteState {
	if next == nil {
		return RouteState{loc: nil, remain: 0, relief: current.relief, route: current.route}
	}
	remain := current.remain - (routes[current.loc.id+"-"+next.id] + 1)
	return RouteState{
		loc:    next,
		remain: remain,
		relief: current.relief + (remain * next.flow),
		route:  current.route + next.id,
	}
}
