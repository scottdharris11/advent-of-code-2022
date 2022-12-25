package day19

import (
	"fmt"
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
		ans += NewBlueprint(line).QualityLevel(24)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 19, Part 1 (%dms): Quality Level = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 1
	for _, line := range lines {
		ans *= NewBlueprint(line).MaxGeode(32)
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

func NewBlueprint(line string) *Blueprint {
	match := blueprintParser.FindStringSubmatch(line)
	if match == nil {
		return nil
	}
	return &Blueprint{
		id:       utils.Number(match[1]),
		ore:      NewRobotCost(match[2]),
		clay:     NewRobotCost(match[3]),
		obsidian: NewRobotCost(match[4]),
		geode:    NewRobotCost(match[5]),
	}
}

type Blueprint struct {
	id       int
	ore      RobotCost
	clay     RobotCost
	obsidian RobotCost
	geode    RobotCost
	minutes  int
}

func (b *Blueprint) QualityLevel(minutes int) int {
	return b.MaxGeode(minutes) * b.id
}

func (b *Blueprint) MaxGeode(minutes int) int {
	b.minutes = minutes
	m := utils.MaxFinder{Maximizer: b}
	s, max := m.Max(BlueprintState{oreRobot: 1}, nil, 0)
	fmt.Printf("ID: %d, Max geode: %d, State: %v\n", b.id, max, s)
	return max
}

func (b *Blueprint) Goal(state interface{}) (bool, int) {
	var bState = state.(BlueprintState)
	return bState.minute == b.minutes, bState.geode
}

func (b *Blueprint) PossibleNextStates(state interface{}, currentMax int) []interface{} {
	var bState = state.(BlueprintState)

	var moves []interface{}
	base := bState.advanceMinute()
	if goal, _ := b.Goal(base); goal {
		moves = append(moves, base)
		return moves
	}
	if !base.canAchieve(currentMax, b.minutes) {
		return nil
	}

	buyCnt := 0
	if bState.canBuy(b.geode) {
		buyCnt++
		moves = append(moves, base.buyGeodeRobot(b.geode, &bState))
	}
	if base.minute < 23 && bState.canBuy(b.obsidian) {
		buyCnt++
		moves = append(moves, base.buyObsidianRobot(b.obsidian, &bState))
	}
	if base.minute < 22 && bState.canBuy(b.clay) {
		buyCnt++
		moves = append(moves, base.buyClayRobot(b.clay, &bState))
	}
	if base.minute < 23 && bState.canBuy(b.ore) {
		buyCnt++
		moves = append(moves, base.buyOreRobot(b.ore, &bState))
	}
	if buyCnt < 4 {
		moves = append(moves, base)
	}
	return moves
}

type BlueprintState struct {
	parent        *BlueprintState
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

func (b BlueprintState) advanceMinute() BlueprintState {
	return BlueprintState{
		parent:        &b,
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

func (b BlueprintState) canBuy(cost RobotCost) bool {
	return b.ore >= cost.ore &&
		b.clay >= cost.clay &&
		b.obsidian >= cost.obsidian
}

func (b BlueprintState) buyOreRobot(cost RobotCost, parent *BlueprintState) BlueprintState {
	s := b.buy(cost, parent)
	s.oreRobot++
	return s
}

func (b BlueprintState) buyClayRobot(cost RobotCost, parent *BlueprintState) BlueprintState {
	s := b.buy(cost, parent)
	s.clayRobot++
	return s
}

func (b BlueprintState) buyObsidianRobot(cost RobotCost, parent *BlueprintState) BlueprintState {
	s := b.buy(cost, parent)
	s.obsidianRobot++
	return s
}

func (b BlueprintState) buyGeodeRobot(cost RobotCost, parent *BlueprintState) BlueprintState {
	s := b.buy(cost, parent)
	s.geodeRobot++
	return s
}

func (b BlueprintState) buy(cost RobotCost, parent *BlueprintState) BlueprintState {
	return BlueprintState{
		parent:        parent,
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

func (b BlueprintState) canAchieve(val int, goalMinutes int) bool {
	minutes := goalMinutes - b.minute + 1
	future := b.geodeRobot * minutes
	future += (minutes * (minutes - 1)) / 2
	return b.geode+future > val
}
