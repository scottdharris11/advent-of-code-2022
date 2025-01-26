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
	var reliefValves []*Valve
	for _, valve := range valves {
		if valve.flow > 0 {
			reliefValves = append(reliefValves, valve)
		}
	}
	ans := bestRelief(valves["AA"], 30, vRoutes, reliefValves, "")
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
	tunnels []*Valve
}

func (v Valve) PressureRelief(openMinutes int) int {
	return v.flow * openMinutes
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

func bestRelief(c *Valve, remain int, routes map[string]int, valves []*Valve, route string) int {
	relief := remain * c.flow
	bestNext := 0
	for _, v := range valves {
		if strings.Contains(route, v.id) {
			continue
		}
		vm := remain - (routes[c.id+"-"+v.id] + 1)
		if vm <= 0 {
			continue
		}
		br := bestRelief(v, vm, routes, valves, route+v.id)
		if br > bestNext {
			bestNext = br
		}
	}
	return relief + bestNext
}
