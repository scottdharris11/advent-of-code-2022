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

func (b BlueprintState) nextMoves(bp Blueprint) []BlueprintState {
	var moves []BlueprintState
	if b.shouldBuyGeodeRobot(bp) {
		moves = append(moves, b.buyGeodeRobot(bp.geode))
	}
	if b.shouldBuyObsidianRobot(bp) {
		moves = append(moves, b.buyObsidianRobot(bp.obsidian))
	}
	if b.shouldBuyClayRobot(bp) {
		moves = append(moves, b.buyClayRobot(bp.clay))
	}
	if b.shouldBuyOreRobot(bp) {
		moves = append(moves, b.buyOreRobot(bp.ore))
	}
	moves = append(moves, b.advanceMinute())
	return moves
}

func (b BlueprintState) canBuy(cost RobotCost) bool {
	return b.ore >= cost.ore &&
		b.clay >= cost.clay &&
		b.obsidian >= cost.obsidian
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

func (b BlueprintState) shouldBuyOreRobot(bp Blueprint) bool {
	return b.oreRobot < bp.maxOre && b.canBuy(bp.ore)
}

func (b BlueprintState) buyOreRobot(cost RobotCost) BlueprintState {
	s := b.advanceMinute()
	s = s.buy(cost)
	s.oreRobot++
	return s
}

func (b BlueprintState) shouldBuyClayRobot(bp Blueprint) bool {
	return b.clayRobot < bp.maxClay && b.canBuy(bp.clay)
}

func (b BlueprintState) buyClayRobot(cost RobotCost) BlueprintState {
	s := b.advanceMinute()
	s = s.buy(cost)
	s.clayRobot++
	return s
}

func (b BlueprintState) shouldBuyObsidianRobot(bp Blueprint) bool {
	return b.obsidianRobot < bp.maxObsidian && b.canBuy(bp.obsidian)
}

func (b BlueprintState) buyObsidianRobot(cost RobotCost) BlueprintState {
	s := b.advanceMinute()
	s = s.buy(cost)
	s.obsidianRobot++
	return s
}

func (b BlueprintState) shouldBuyGeodeRobot(bp Blueprint) bool {
	return b.canBuy(bp.geode)
}

func (b BlueprintState) buyGeodeRobot(cost RobotCost) BlueprintState {
	s := b.advanceMinute()
	s = s.buy(cost)
	s.geodeRobot++
	return s
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

func maxGeodes(remain int, b BlueprintState, bp Blueprint, best int) int {
	if remain == 0 {
		return b.geode
	}
	if best > 0 && b.possibleGeodes(remain) < best {
		return -1
	}
	for _, next := range b.nextMoves(bp) {
		g := maxGeodes(remain-1, next, bp, best)
		if g > 0 && g > best {
			best = g
		}
	}
	return best
}
