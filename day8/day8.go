package day8

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day8", "day-8-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	grid := utils.ReadIntegerGrid(lines)
	ans := treesVisible(grid)
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 1 (%dms): Trees Visible = %d", end-start, ans)
	return ans
}

func treesVisible(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	v := (height * 2) + ((width - 2) * 2)
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if visible(grid, x, y) {
				v++
			}
		}
	}
	return v
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	grid := utils.ReadIntegerGrid(lines)
	ans := topScore(grid)
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 2 (%dms): Top Scenic Score = %d", end-start, ans)
	return ans
}

func topScore(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	topScore := 0
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			s := scenicScore(grid, x, y)
			if s > topScore {
				topScore = s
			}
		}
	}
	return topScore
}

func visible(grid [][]int, x int, y int) bool {
	if ok, _ := visibleUp(grid, x, y); ok {
		return true
	}
	if ok, _ := visibleDown(grid, x, y); ok {
		return true
	}
	if ok, _ := visibleLeft(grid, x, y); ok {
		return true
	}
	if ok, _ := visibleRight(grid, x, y); ok {
		return true
	}
	return false
}

func scenicScore(grid [][]int, x int, y int) int {
	_, u := visibleUp(grid, x, y)
	_, d := visibleDown(grid, x, y)
	_, l := visibleLeft(grid, x, y)
	_, r := visibleRight(grid, x, y)
	return u * d * l * r
}

func visibleUp(grid [][]int, x int, y int) (bool, int) {
	height := grid[y][x]
	up := y - 1
	cnt := 0
	for up >= 0 {
		cnt++
		if grid[up][x] >= height {
			return false, cnt
		}
		up--
	}
	return true, cnt
}

func visibleDown(grid [][]int, x int, y int) (bool, int) {
	height := grid[y][x]
	down := y + 1
	cnt := 0
	for down < len(grid) {
		cnt++
		if grid[down][x] >= height {
			return false, cnt
		}
		down++
	}
	return true, cnt
}

func visibleLeft(grid [][]int, x int, y int) (bool, int) {
	height := grid[y][x]
	left := x - 1
	cnt := 0
	for left >= 0 {
		cnt++
		if grid[y][left] >= height {
			return false, cnt
		}
		left--
	}
	return true, cnt
}

func visibleRight(grid [][]int, x int, y int) (bool, int) {
	height := grid[y][x]
	right := x + 1
	cnt := 0
	for right < len(grid[0]) {
		cnt++
		if grid[y][right] >= height {
			return false, cnt
		}
		right++
	}
	return true, cnt
}
