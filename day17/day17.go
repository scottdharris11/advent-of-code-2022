package day17

import (
	"fmt"
	"log"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day17", "day-17-input.txt")
	solvePart1(input[0])
	solvePart2(input[0])
}

func solvePart1(wind string) int {
	start := time.Now().UnixMilli()
	cave := NewCave(wind, RockPatterns)
	ans := 0
	for i := 0; i < 2022; i++ {
		ans = cave.dropRock()
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 1 (%dms): Rock Height = %d", end-start, ans)
	return ans
}

func solvePart2(wind string) int {
	start := time.Now().UnixMilli()
	ans := simulateRockFalling(wind, 1000000000000)
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 2 (%dms): Rock Height = %d", end-start, ans)
	return ans
}

func simulateRockFalling(wind string, rocks int) int {
	cave := NewCave(wind, RockPatterns)
	repeat := len(wind) * len(RockPatterns)
	var sections []int
	var ans int
	var remain int
	var prev int
	for i := 0; i < rocks; i++ {
		if i%1000 == 0 {
			cave.prune()
		}
		if i > 0 && i%repeat == 0 {
			top := cave.topRock()
			sections = append(sections, top-prev)
			prev = top
			l := len(sections)
			if l > 1 && l%2 == 1 {
				split := (l - 1) / 2
				if equal(sections[1:split+1], sections[split+1:]) {
					s := sum(sections[1 : split+1])
					ans = (rocks / (repeat * split)) * s
					ans += sections[0]
					remain = (rocks % (repeat * split)) - repeat
					break
				}
			}
		}
		cave.dropRock()
	}
	for i := 0; i < remain; i++ {
		cave.dropRock()
	}
	ans += cave.topRock() - prev
	return ans
}

func sum(integers []int) int {
	s := 0
	for _, i := range integers {
		s += i
	}
	return s
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type WindDirection rune

var Left WindDirection = '<'
var Right WindDirection = '>'

var RockPatterns = []Rock{
	{height: 1, width: 4, pattern: []string{
		"####",
	}},
	{height: 3, width: 3, pattern: []string{
		" # ",
		"###",
		" # ",
	}},
	{height: 3, width: 3, pattern: []string{
		"  #",
		"  #",
		"###",
	}},
	{height: 4, width: 1, pattern: []string{
		"#",
		"#",
		"#",
		"#",
	}},
	{height: 2, width: 2, pattern: []string{
		"##",
		"##",
	}},
}

type Rock struct {
	height  int
	width   int
	pattern []string
}

func (r Rock) CanShiftLeft(xOffset int, rows []string) bool {
	if xOffset == 0 {
		return false
	}
	for i := 0; i < r.height; i++ {
		for j := 0; j < r.width; j++ {
			if r.pattern[i][j] == '#' && rows[i][j+xOffset-1] == '#' {
				return false
			}
		}
	}
	return true
}

func (r Rock) CanShiftRight(xOffset int, rows []string) bool {
	if xOffset+r.width-1 >= len(rows[0])-1 {
		return false
	}
	for i := 0; i < r.height; i++ {
		for j := 0; j < r.width; j++ {
			if r.pattern[i][j] == '#' && rows[i][j+xOffset+1] == '#' {
				return false
			}
		}
	}
	return true
}

func (r Rock) CanShiftDown(xOffset int, rows []string) bool {
	for h := r.height - 1; h >= 0; h-- {
		below := rows[h+1]
		bottom := r.pattern[h]
		for i := 0; i < r.width; i++ {
			if bottom[i] == '#' && below[i+xOffset] == '#' {
				return false
			}
		}
	}
	return true
}

func NewCave(wind string, rocks []Rock) *Cave {
	c := &Cave{width: 7, wind: wind, rocks: rocks}
	c.rows = append(c.rows, "#######")
	return c
}

type Cave struct {
	width   int
	wind    string
	wOffset int
	rocks   []Rock
	rOffset int
	rows    []string
	pruned  int
}

func (c *Cave) windDirection() WindDirection {
	w := c.wind[c.wOffset]
	c.wOffset++
	if c.wOffset >= len(c.wind) {
		c.wOffset = 0
	}
	return WindDirection(w)
}

func (c *Cave) rock() Rock {
	o := c.rOffset
	c.rOffset++
	if c.rOffset >= len(c.rocks) {
		c.rOffset = 0
	}
	return c.rocks[o]
}

func (c *Cave) dropRock() int {
	rock := c.rock()
	xOffset := 2
	yOffset := len(c.rows) - 1 + 4

	for {
		impactRows := c.impactRows(yOffset, rock.height)

		switch c.windDirection() {
		case Left:
			if rock.CanShiftLeft(xOffset, impactRows) {
				xOffset--
			}
		case Right:
			if rock.CanShiftRight(xOffset, impactRows) {
				xOffset++
			}
		}

		if !rock.CanShiftDown(xOffset, impactRows) {
			c.addRock(xOffset, yOffset, rock)
			break
		}
		yOffset--
	}

	return c.topRock()
}

func (c *Cave) impactRows(yOffset int, rockHeight int) []string {
	var rows []string
	caveHeight := len(c.rows)
	for i := rockHeight; i >= 0; i-- {
		caveRow := yOffset + i - 1
		if caveRow >= caveHeight {
			rows = append(rows, "       ")
			continue
		}
		rows = append(rows, c.rows[caveRow])
	}
	return rows
}

func (c *Cave) addRock(xOffset int, yOffset int, rock Rock) {
	caveHeight := len(c.rows)
	for i := 0; i < rock.height; i++ {
		caveRow := yOffset + i
		var row []rune
		switch {
		case caveRow == caveHeight:
			row = []rune("       ")
			c.rows = append(c.rows, string(row))
			caveHeight++
		default:
			row = []rune(c.rows[caveRow])
		}
		for j := 0; j < rock.width; j++ {
			if rock.pattern[rock.height-i-1][j] == ' ' {
				continue
			}
			row[j+xOffset] = '#'
		}
		c.rows[caveRow] = string(row)
	}
}

func (c *Cave) topRock() int {
	return c.pruned + len(c.rows) - 1
}

func (c *Cave) print() {
	for i := len(c.rows) - 1; i >= 0; i-- {
		row := strings.ReplaceAll(c.rows[i], " ", ".")
		b := "|"
		if i == 0 {
			b = "+"
			row = strings.ReplaceAll(row, "#", "-")
		}
		fmt.Printf("%s%s%s\n", b, row, b)
	}
}

func (c *Cave) prune() {
	row := -1
	for i := len(c.rows) - 1; i >= 0; i-- {
		if c.rows[i][0] == '#' {
			row = i
			break
		}
	}

	if row == -1 {
		return
	}
	pruneRow := c.followPath(row)
	if pruneRow == -1 {
		return
	}

	c.pruned += pruneRow
	c.rows = c.rows[pruneRow:]
}

func (c *Cave) followPath(row int) int {
	search := utils.Search{Searcher: c}
	solution := search.Best(utils.SearchMove{
		Cost:  0,
		State: PruneState{row: row, col: 0},
	})
	if solution == nil {
		return -1
	}

	minRow := row
	for _, s := range solution.Path {
		var state = s.(PruneState)
		if state.row < minRow {
			minRow = state.row
		}
	}
	return minRow
}

func (c *Cave) Goal(state interface{}) bool {
	var pState = state.(PruneState)
	return pState.col == c.width-1
}

func (c *Cave) DistanceFromGoal(state interface{}) int {
	var pState = state.(PruneState)
	return c.width - pState.col - 1
}

func (c *Cave) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var pState = state.(PruneState)
	row := pState.row
	col := pState.col

	var moves []utils.SearchMove
	for _, rowCol := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		y := row + rowCol[0]
		x := col + rowCol[1]
		if x >= 0 && x < c.width && y >= 0 && y < len(c.rows)-1 && c.rows[y][x] == '#' {
			moves = append(moves, utils.SearchMove{
				Cost:  1,
				State: PruneState{row: y, col: x},
			})
		}
	}
	return moves
}

type PruneState struct {
	row int
	col int
}
