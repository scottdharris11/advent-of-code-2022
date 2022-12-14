package day13

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day13", "day-13-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	pair := 1
	for i := 0; i < len(lines); i += 3 {
		left := NewPacket(lines[i])
		right := NewPacket(lines[i+1])
		if left.Ordered(*right) {
			ans += pair
		}
		pair++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 1 (%dms): Ordered = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}

type Packet struct {
	value  int
	values []*Packet
}

func (p Packet) Ordered(p2 Packet) bool {
	return p.compare(p2) <= 0
}

func (p Packet) compare(p2 Packet) int {
	if p.value >= 0 && p2.value >= 0 {
		switch {
		case p.value < p2.value:
			return -1
		case p.value > p2.value:
			return 1
		default:
			return 0
		}
	}

	left := p.values
	if p.value >= 0 {
		left = []*Packet{{value: p.value}}
	}

	right := p2.values
	if p2.value >= 0 {
		right = []*Packet{{value: p2.value}}
	}

	for i, l := range left {
		if i+1 > len(right) {
			return 1
		}
		r := right[i]
		c := l.compare(*r)
		if c == 0 {
			continue
		}
		return c
	}

	return 0
}

func NewPacket(s string) *Packet {
	stack := utils.Stack{}
	var p *Packet
	var begin int
	var pop bool
	for i, r := range s {
		switch r {
		case '[':
			np := &Packet{value: -1}
			if p != nil {
				stack.Push(p)
				p.values = append(p.values, np)
			}
			p = np
			begin = i + 1
		case ']':
			pop = true
			fallthrough
		case ',':
			if begin < i {
				v := utils.Number(s[begin:i])
				p.values = append(p.values, &Packet{value: v})
			}
			begin = i + 1
		}

		if pop {
			pop = false
			pp := stack.Pop()
			if pp != nil {
				if pp, ok := pp.(*Packet); ok {
					p = pp
				}
			}
		}
	}
	return p
}
