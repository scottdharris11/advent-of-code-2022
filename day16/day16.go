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
	ans, route := maxReliefRoute(valves["AA"], 30, vRoutes, rValves, "")
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 1 (%dms): Max Released vi Route(%s) = %d", end-start, route, ans)
	return ans
}

func removeValve(valves []*Valve, id string) []*Valve {
	idx := -1
	for i, v := range valves {
		if v.id == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return valves
	}
	return append(valves[:idx], valves[idx+1:]...)
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	valves := parseValves(lines)
	vRoutes := valveRoutes(valves)
	rValves := reliefValves(valves)

	ans, route := maxReliefRoute(valves["AA"], 26, vRoutes, rValves, "")
	log.Printf("Max Released vi Route(%s) = %d", route, ans)
	for i := 0; i < len(route); i += 2 {
		rValves = removeValve(rValves, route[i:i+2])
	}
	ans, route = maxReliefRoute(valves["AA"], 26, vRoutes, rValves, "")
	log.Printf("Max Released vi Route(%s) = %d", route, ans)

	//ans := bestElephantRelief(0, valves["AA"], valves["AA"], 0, 0, 26, vRoutes, rValves, "")
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 2 (%dms): Max Released = %d", end-start, ans)
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

func reliefValves(valves map[string]*Valve) []*Valve {
	var rValves []*Valve
	for _, valve := range valves {
		if valve.flow > 0 {
			rValves = append(rValves, valve)
		}
	}
	return rValves
}

func maxReliefRoute(c *Valve, remain int, routes map[string]int, valves []*Valve, route string) (int, string) {
	relief := remain * c.flow
	bestNext := 0
	bestRoute := route
	for _, v := range valves {
		if strings.Contains(route, v.id) {
			continue
		}
		vm := remain - (routes[c.id+"-"+v.id] + 1)
		if vm <= 0 {
			continue
		}
		br, vr := maxReliefRoute(v, vm, routes, valves, route+v.id)
		if br > bestNext {
			bestNext = br
			bestRoute = vr
		}
	}
	return relief + bestNext, bestRoute
}

func bestElephantRelief(best int, mc *Valve, ec *Valve, mSteps int, eSteps int, remain int, routes map[string]int, valves []*Valve, route string) int {
	relief := 0
	assign := 0
	potential := 0
	if mSteps == 0 && mc != nil {
		relief += remain * mc.flow
		//log.Printf("route %s, valve %s opened at %d, total relief = %d", route, mc.id, remain, remain*mc.flow)
		assign++
	} else if mc != nil {
		potential += (remain - mSteps) * mc.flow
	}
	if eSteps == 0 && ec != nil {
		relief += remain * ec.flow
		//log.Printf("route %s, valve %s opened at %d, total relief = %d", route, ec.id, remain, remain*ec.flow)
		assign++
	} else if ec != nil {
		potential += (remain - eSteps) * ec.flow
	}

	if remain == 0 {
		return relief
	}

	var left []*Valve
	for _, v := range valves {
		if strings.Contains(route, v.id) {
			continue
		}
		left = append(left, v)
		potential += (remain - 2) * v.flow
	}

	if assign == 0 || len(left) == 0 {
		if mSteps > 0 || eSteps > 0 {
			return relief + bestElephantRelief(best, mc, ec, mSteps-1, eSteps-1, remain-1, routes, valves, route)
		}
		return relief
	}

	if relief+potential < best {
		//log.Printf("abandon route: %s, best: %d, potential: %d", route, best, relief+potential)
		return 0
	}

	if assign == 1 {
		bestNext := 0
		for _, v := range left {
			c := mc
			if eSteps == 0 && ec != nil {
				c = ec
			}
			steps := routes[c.id+"-"+v.id] + 1
			if remain-steps <= 0 {
				continue
			}
			br := 0
			nr := route + v.id
			if mSteps == 0 && mc != nil {
				br = bestElephantRelief(bestNext, v, ec, steps-1, eSteps-1, remain-1, routes, valves, nr)
			} else {
				br = bestElephantRelief(bestNext, mc, v, mSteps-1, steps-1, remain-1, routes, valves, nr)
			}
			if br > bestNext {
				//log.Printf("selecting valve %s is best so far: %d", v.id, br)
				bestNext = br
			}
		}
		return relief + bestNext
	}
	bestNext := 0
	var combos [][]*Valve
	if mc == ec {
		combos = uniquecombos(left)
	} else {
		combos = allcombos(left)
	}

	for _, combo := range combos {
		if remain == 26 {
			log.Printf("starting combo of %s,%s", combo[0].id, combo[1].id)
		}
		m := routes[mc.id+"-"+combo[0].id]
		nr := route + combo[0].id
		e := 0
		if combo[1] != nil {
			e = routes[ec.id+"-"+combo[1].id]
			nr += combo[1].id
		}
		br := bestElephantRelief(bestNext, combo[0], combo[1], m, e, remain-1, routes, valves, nr)
		if remain == 26 {
			log.Printf("results of combo of %s,%s: %d", combo[0].id, combo[1].id, br)
		}
		if br > bestNext {
			//id0 := combo[0].id
			//id1 := ""
			//if combo[1] != nil {
			//	id1 = combo[1].id
			//}
			//log.Printf("results for combo of %s,%s is best so far: %d", id0, id1, br)
			bestNext = br
		}
	}
	return relief + bestNext
}

func uniquecombos(valves []*Valve) [][]*Valve {
	if len(valves) == 1 {
		return [][]*Valve{{valves[0], nil}}
	}
	var c [][]*Valve
	for a := 0; a < len(valves)-1; a++ {
		for b := a + 1; b < len(valves); b++ {
			c = append(c, []*Valve{valves[a], valves[b]})
		}
	}
	return c
}

func allcombos(valves []*Valve) [][]*Valve {
	if len(valves) == 1 {
		return [][]*Valve{{valves[0], nil}}
	}
	var c [][]*Valve
	for a := 0; a < len(valves); a++ {
		for b := 0; b < len(valves); b++ {
			if a == b {
				continue
			}
			c = append(c, []*Valve{valves[a], valves[b]})
		}
	}
	return c
}
