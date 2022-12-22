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
	var ans int
	for i := 0; i < 2022; i++ {
		ans = cave.dropRock()
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 1 (%dms): Rock Height = %d", end-start, ans)
	return ans
}

func solvePart2(wind string) int {
	start := time.Now().UnixMilli()
	ans := len(wind)
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 2 (%dms): Rock Height = %d", end-start, ans)
	return ans
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
	yOffset := c.topRock() + 4

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
	return len(c.rows) - 1
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
