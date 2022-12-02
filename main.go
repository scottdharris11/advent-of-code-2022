package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
)

type Solver interface {
	Solve()
}

func main() {
	solvers := []Solver{
		day1.Puzzle{}, day2.Puzzle{},
	}
	for _, solver := range solvers {
		solver.Solve()
	}
}
