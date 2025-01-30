package day16

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
	"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
	"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
	"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
	"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
	"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
	"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
	"Valve HH has flow rate=22; tunnel leads to valve GG",
	"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
	"Valve JJ has flow rate=21; tunnel leads to valve II",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1651, solvePart1(testInput))
	assert.Equal(t, 1580, solvePart1(utils.ReadLines("day16", "day-16-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 1707, solvePart2(testInput))
	assert.Equal(t, 0, solvePart2(utils.ReadLines("day16", "day-16-input.txt")))
}

func TestValve_PressureRelief(t *testing.T) {
	tests := []struct {
		valve   Valve
		minutes int
		relief  int
	}{
		{Valve{id: "BB", flow: 13}, 28, 364},
		{Valve{id: "CC", flow: 2}, 26, 52},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.relief, tt.valve.PressureRelief(tt.minutes))
		})
	}
}
