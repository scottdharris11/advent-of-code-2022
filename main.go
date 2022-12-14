package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day12"
	"advent-of-code-2022/day13"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
)

type Solver interface {
	Solve()
}

func main() {
	solvers := []Solver{
		day1.Puzzle{}, day2.Puzzle{}, day3.Puzzle{}, day4.Puzzle{}, day5.Puzzle{},
		day6.Puzzle{}, day7.Puzzle{}, day8.Puzzle{}, day9.Puzzle{}, day10.Puzzle{},
		day11.Puzzle{}, day12.Puzzle{}, day13.Puzzle{},
	}
	for _, solver := range solvers {
		solver.Solve()
	}
}
