package day19

import (
	"log"
	"regexp"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day19", "day-19-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	for _, line := range lines {
		bp := NewBlueprint(line)
		mg := maxGeodes(24, BlueprintState{oreRobot: 1}, bp, -1)
		ans += bp.qualityLevel(mg)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 19, Part 1 (%dms): Quality Level = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	trimmed := lines
	if len(lines) > 3 {
		trimmed = lines[:3]
	}
	ans := 1
	for _, line := range trimmed {
		bp := NewBlueprint(line)
		mg := maxGeodes(32, BlueprintState{oreRobot: 1}, bp, -1)
		ans *= mg
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 19, Part 2 (%dms): Max Geodes = %d", end-start, ans)
	return ans
}

var robotCostParser = regexp.MustCompile(`([\d]+) (ore|clay|obsidian)`)

func NewRobotCost(s string) RobotCost {
	r := RobotCost{}
	matches := robotCostParser.FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		switch m[2] {
		case "ore":
			r.ore = utils.Number(m[1])
		case "clay":
			r.clay = utils.Number(m[1])
		case "obsidian":
			r.obsidian = utils.Number(m[1])
		}
	}
	return r
}

type RobotCost struct {
	ore      int
	clay     int
	obsidian int
}

var blueprintParser = regexp.MustCompile(`Blueprint ([\d]+): ([^.]+). ([^.]+). ([^.]+). ([^.]+).`)

func NewBlueprint(line string) Blueprint {
	match := blueprintParser.FindStringSubmatch(line)
	oreCost := NewRobotCost(match[2])
	maxOre, maxClay, maxObsidian := maxCosts(oreCost, 0, 0, 0)
	clayCost := NewRobotCost(match[3])
	maxOre, maxClay, maxObsidian = maxCosts(clayCost, maxOre, maxClay, maxObsidian)
	obsidianCost := NewRobotCost(match[4])
	maxOre, maxClay, maxObsidian = maxCosts(obsidianCost, maxOre, maxClay, maxObsidian)
	geodeCost := NewRobotCost(match[5])
	maxOre, maxClay, maxObsidian = maxCosts(geodeCost, maxOre, maxClay, maxObsidian)
	return Blueprint{
		id:          utils.Number(match[1]),
		ore:         oreCost,
		maxOre:      maxOre,
		clay:        clayCost,
		maxClay:     maxClay,
		obsidian:    obsidianCost,
		maxObsidian: maxObsidian,
		geode:       geodeCost,
	}
}

func maxCosts(rc RobotCost, maxOre int, maxClay int, maxObsidian int) (int, int, int) {
	or := maxOre
	if rc.ore > or {
		or = rc.ore
	}
	cl := maxClay
	if rc.clay > cl {
		cl = rc.clay
	}
	ob := maxObsidian
	if rc.obsidian > ob {
		ob = rc.obsidian
	}
	return or, cl, ob
}

type Blueprint struct {
	id          int
	ore         RobotCost
	maxOre      int
	clay        RobotCost
	maxClay     int
	obsidian    RobotCost
	maxObsidian int
	geode       RobotCost
}

func (b Blueprint) qualityLevel(geodeCnt int) int {
	if geodeCnt < 0 {
		return 0
	}
	return geodeCnt * b.id
}

type BlueprintState struct {
	minute        int
	oreRobot      int
	ore           int
	clayRobot     int
	clay          int
	obsidianRobot int
	obsidian      int
	geodeRobot    int
	geode         int
}

func (b BlueprintState) nextBuyStates(bp Blueprint, remain int) []BlueprintState {
	var moves []BlueprintState
	state := b.buyGeodeNextState(bp, remain)
	if state != nil {
		moves = append(moves, *state)
	}
	state = b.buyObsidianNextState(bp, remain)
	if state != nil {
		moves = append(moves, *state)
	}
	state = b.buyClayNextState(bp, remain)
	if state != nil {
		moves = append(moves, *state)
	}
	state = b.buyOreNextState(bp, remain)
	if state != nil {
		moves = append(moves, *state)
	}
	return moves
}

func (b BlueprintState) advanceMinute() BlueprintState {
	return BlueprintState{
		minute:        b.minute + 1,
		oreRobot:      b.oreRobot,
		ore:           b.ore + b.oreRobot,
		clayRobot:     b.clayRobot,
		clay:          b.clay + b.clayRobot,
		obsidianRobot: b.obsidianRobot,
		obsidian:      b.obsidian + b.obsidianRobot,
		geodeRobot:    b.geodeRobot,
		geode:         b.geode + b.geodeRobot,
	}
}

func (b BlueprintState) buyOreNextState(bp Blueprint, remain int) *BlueprintState {
	if b.oreRobot == bp.maxOre {
		return nil
	}
	bs := b.buyNextState(bp.ore, remain)
	if bs == nil {
		return nil
	}
	bs.oreRobot++
	return bs
}

func (b BlueprintState) buyClayNextState(bp Blueprint, remain int) *BlueprintState {
	if b.clayRobot == bp.maxClay {
		return nil
	}
	bs := b.buyNextState(bp.clay, remain)
	if bs == nil {
		return nil
	}
	bs.clayRobot++
	return bs
}

func (b BlueprintState) buyObsidianNextState(bp Blueprint, remain int) *BlueprintState {
	if b.obsidianRobot == bp.maxObsidian {
		return nil
	}
	bs := b.buyNextState(bp.obsidian, remain)
	if bs == nil {
		return nil
	}
	bs.obsidianRobot++
	return bs
}

func (b BlueprintState) buyGeodeNextState(bp Blueprint, remain int) *BlueprintState {
	bs := b.buyNextState(bp.geode, remain)
	if bs == nil {
		return nil
	}
	bs.geodeRobot++
	return bs
}

func (b BlueprintState) buyNextState(rc RobotCost, remain int) *BlueprintState {
	if rc.ore > 0 && b.oreRobot == 0 {
		return nil
	}
	if rc.clay > 0 && b.clayRobot == 0 {
		return nil
	}
	if rc.obsidian > 0 && b.obsidianRobot == 0 {
		return nil
	}

	steps := 1
	if rc.ore > 0 && b.ore < rc.ore {
		oreDiff := rc.ore - b.ore
		oreSteps := (oreDiff / b.oreRobot) + 1
		if oreDiff%b.oreRobot > 0 {
			oreSteps++
		}
		if oreSteps > steps {
			steps = oreSteps
		}
	}

	if rc.clay > 0 && b.clay < rc.clay {
		clayDiff := rc.clay - b.clay
		claySteps := (clayDiff / b.clayRobot) + 1
		if clayDiff%b.clayRobot > 0 {
			claySteps++
		}
		if claySteps > steps {
			steps = claySteps
		}
	}

	if rc.obsidian > 0 && b.obsidian < rc.obsidian {
		obDiff := rc.obsidian - b.obsidian
		obSteps := (obDiff / b.obsidianRobot) + 1
		if obDiff%b.obsidianRobot > 0 {
			obSteps++
		}
		if obSteps > steps {
			steps = obSteps
		}
	}

	if steps >= remain {
		return nil
	}

	s := b.advanceMinute()
	for i := 1; i < steps; i++ {
		s = s.advanceMinute()
	}
	s = s.buy(rc)
	return &s
}

func (b BlueprintState) buy(cost RobotCost) BlueprintState {
	return BlueprintState{
		minute:        b.minute,
		oreRobot:      b.oreRobot,
		ore:           b.ore - cost.ore,
		clayRobot:     b.clayRobot,
		clay:          b.clay - cost.clay,
		obsidianRobot: b.obsidianRobot,
		obsidian:      b.obsidian - cost.obsidian,
		geodeRobot:    b.geodeRobot,
		geode:         b.geode,
	}
}

func (b BlueprintState) possibleGeodes(remain int) int {
	geodes := b.geode + (b.geodeRobot * remain)
	for g := remain - 1; g > 0; g-- {
		geodes += g
	}
	return geodes
}

func maxGeodes(minutes int, b BlueprintState, bp Blueprint, best int) int {
	remain := minutes - b.minute
	if remain <= 0 {
		return b.geode
	}
	if best > 0 && b.possibleGeodes(remain) < best {
		return -1
	}
	buys := b.nextBuyStates(bp, remain)
	if len(buys) == 0 {
		// no more buys, advance to the end with current robots
		next := b.advanceMinute()
		for r := 1; r < remain; r++ {
			next = next.advanceMinute()
		}
		return next.geode
	}
	for _, next := range buys {
		g := maxGeodes(minutes, next, bp, best)
		if g > 0 && g > best {
			best = g
		}
	}
	return best
}
