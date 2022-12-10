package day10

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCommands = []string{
	"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8", "addx 13", "addx 4",
	"noop", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5",
	"addx -1", "addx -35", "addx 1", "addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop",
	"noop", "addx 21", "addx -15", "noop", "noop", "addx -3", "addx 9", "addx 1", "addx -3", "addx 8",
	"addx 1", "addx 5", "noop", "noop", "noop", "noop", "noop", "addx -36", "noop", "addx 1", "addx 7",
	"noop", "noop", "noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop", "noop", "addx 1", "noop",
	"noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7", "noop", "addx 1", "addx -33", "noop",
	"noop", "noop", "addx 2", "noop", "noop", "noop", "addx 8", "noop", "addx -1", "addx 2", "addx 1",
	"noop", "addx 17", "addx -9", "addx 1", "addx 1", "addx -3", "addx 11", "noop", "noop", "addx 1", "noop",
	"addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1", "addx 3", "addx 26", "addx -30", "addx 12",
	"addx -1", "addx 3", "addx 1", "noop", "noop", "noop", "addx -9", "addx 18", "addx 1", "addx 2", "noop",
	"noop", "addx 9", "noop", "noop", "noop", "addx -1", "addx 2", "addx -37", "addx 1", "addx 3", "noop",
	"addx 15", "addx -21", "addx 22", "addx -6", "addx 1", "noop", "addx 2", "addx 1", "noop", "addx -10",
	"noop", "noop", "addx 20", "addx 1", "addx 2", "addx 2", "addx -6", "addx -11", "noop", "noop", "noop",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 13140, solvePart1(testCommands))
	assert.Equal(t, 16060, solvePart1(utils.ReadLines("day10", "day-10-input.txt")))
}

var crtExpected1 = "##..##..##..##..##..##..##..##..##..##..\n" +
	"###...###...###...###...###...###...###.\n" +
	"####....####....####....####....####....\n" +
	"#####.....#####.....#####.....#####.....\n" +
	"######......######......######......####\n" +
	"#######.......#######.......#######.....\n"

var crtExpected2 = "###...##...##..####.#..#.#....#..#.####.\n" +
	"#..#.#..#.#..#.#....#.#..#....#..#.#....\n" +
	"###..#..#.#....###..##...#....####.###..\n" +
	"#..#.####.#....#....#.#..#....#..#.#....\n" +
	"#..#.#..#.#..#.#....#.#..#....#..#.#....\n" +
	"###..#..#..##..####.#..#.####.#..#.#....\n"

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, crtExpected1, solvePart2(testCommands))
	assert.Equal(t, crtExpected2, solvePart2(utils.ReadLines("day10", "day-10-input.txt")))
}

func TestStatus_SignalStrength(t *testing.T) {
	tests := []struct {
		status      Status
		sigStrength int
	}{
		{Status{cycle: 20, register: 21}, 420},
		{Status{cycle: 60, register: 19}, 1140},
		{Status{cycle: 100, register: 18}, 1800},
		{Status{cycle: 140, register: 21}, 2940},
		{Status{cycle: 180, register: 16}, 2880},
		{Status{cycle: 220, register: 18}, 3960},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.sigStrength, tt.status.signalStrength())
		})
	}
}

func TestCPU_ExecuteOp(t *testing.T) {
	m := make(chan Status, 10)
	c := NewCPU(m)
	c.ExecuteOp(1, 0)
	c.ExecuteOp(2, 3)
	c.ExecuteOp(2, -5)
	close(m)

	v, ok := <-m
	assert.Equal(t, Status{1, 1}, v)
	assert.True(t, ok)

	v, ok = <-m
	assert.Equal(t, Status{2, 1}, v)
	assert.True(t, ok)

	v, ok = <-m
	assert.Equal(t, Status{3, 1}, v)
	assert.True(t, ok)

	v, ok = <-m
	assert.Equal(t, Status{4, 4}, v)
	assert.True(t, ok)

	v, ok = <-m
	assert.Equal(t, Status{5, 4}, v)
	assert.True(t, ok)

	_, ok = <-m
	assert.False(t, ok)
	assert.Equal(t, -1, c.register)
}

func TestMonitor_MonitorCPU_SigStrength(t *testing.T) {
	s := make(chan Status, 20)
	o := make(chan int)
	m := NewMonitor(s, o, make(chan string))

	go m.monitorCPU()
	s <- Status{cycle: 1, register: 1}
	s <- Status{cycle: 20, register: 21}
	s <- Status{cycle: 22, register: 21}
	s <- Status{cycle: 60, register: 19}
	s <- Status{cycle: 61, register: 19}
	s <- Status{cycle: 99, register: 18}
	s <- Status{cycle: 100, register: 18}
	s <- Status{cycle: 140, register: 21}
	s <- Status{cycle: 141, register: 21}
	s <- Status{cycle: 180, register: 16}
	s <- Status{cycle: 185, register: 16}
	s <- Status{cycle: 220, register: 18}
	close(s)

	assert.Equal(t, 13140, <-o)
}
