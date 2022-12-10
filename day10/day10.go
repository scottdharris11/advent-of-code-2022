package day10

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day10", "day-10-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()

	status := make(chan Status)
	output := make(chan int)
	m := NewMonitor(status, output, make(chan string, 5))
	c := NewCPU(status)

	go m.monitorCPU()
	executeOperations(lines, c)
	close(status)

	ans := <-output
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 1 (%dms): Signal Strength = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) string {
	start := time.Now().UnixMilli()
	status := make(chan Status)
	output := make(chan string)
	m := NewMonitor(status, make(chan int, 5), output)
	c := NewCPU(status)

	go m.monitorCPU()
	executeOperations(lines, c)
	close(status)

	ans := <-output
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 2 (%dms): Output = \n%s", end-start, ans)
	return ans
}

func executeOperations(ops []string, c *CPU) {
	for _, op := range ops {
		if op == "noop" {
			c.ExecuteOp(1, 0)
			continue
		}
		c.ExecuteOp(2, utils.Number(strings.Split(op, " ")[1]))
	}
}

type Status struct {
	cycle    int
	register int
}

func (s *Status) signalStrength() int {
	return s.cycle * s.register
}

func NewCPU(monitor chan Status) *CPU {
	return &CPU{cycle: 0, register: 1, monitor: monitor}
}

type CPU struct {
	cycle    int
	register int
	monitor  chan Status
}

func (c *CPU) ExecuteOp(cycles int, registerAdjust int) {
	for i := 0; i < cycles; i++ {
		c.cycle++
		if c.cycle == 1 {
			c.register = 1
		}
		if c.monitor != nil {
			c.monitor <- Status{cycle: c.cycle, register: c.register}
		}
		if i == cycles-1 {
			c.register += registerAdjust
		}
	}
}

func NewMonitor(status chan Status, strength chan int, crt chan string) *Monitor {
	return &Monitor{iStatus: status, oStrength: strength, oCrt: crt}
}

type Monitor struct {
	iStatus   chan Status
	oStrength chan int
	oCrt      chan string
	pixels    [6][40]rune
}

func (m *Monitor) monitorCPU() {
	row, col, sigSum := 0, 0, 0
	for s := range m.iStatus {
		m.pixels[row][col], row, col = m.pixel(s, row, col)
		if s.cycle == 20 || (s.cycle-20)%40 == 0 {
			sigSum += s.signalStrength()
		}
	}
	m.oStrength <- sigSum
	m.oCrt <- m.render()
}

func (m *Monitor) pixel(s Status, row int, col int) (rune, int, int) {
	pixel := '.'
	if col >= s.register-1 && col <= s.register+1 {
		pixel = '#'
	}

	col++
	if col == len(m.pixels[0]) {
		row++
		if row > len(m.pixels) {
			row = 0
		}
		col = 0
	}

	return pixel, row, col
}

func (m *Monitor) render() string {
	sb := strings.Builder{}
	for _, l := range m.pixels {
		sb.WriteString(string(l[:40]))
		sb.WriteRune('\n')
	}
	return sb.String()
}
