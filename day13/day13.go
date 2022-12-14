package day13

import (
	"log"
	"reflect"
	"sort"
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
	packets := orderedPackets(append(lines, []string{"[[2]]", "[[6]]"}...))
	div1Marker := NewPacket("[[2]]")
	div2Marker := NewPacket("[[6]]")
	div1Idx := 0
	div2Idx := 0
	for i, p := range packets {
		if reflect.DeepEqual(p, div1Marker) {
			div1Idx = i + 1
			continue
		}
		if reflect.DeepEqual(p, div2Marker) {
			div2Idx = i + 1
			break
		}
	}
	ans := div1Idx * div2Idx
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 2 (%dms): Decoder Key = %d", end-start, ans)
	return ans
}

type Packet struct {
	value  int
	values []*Packet
}

func (p Packet) Ordered(p2 Packet) bool {
	return p.compare(p2) < 0
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

	if len(left) < len(right) {
		return -1
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

func orderedPackets(lines []string) []*Packet {
	var packets []*Packet
	for _, l := range lines {
		if l == "" {
			continue
		}
		packets = append(packets, NewPacket(l))
	}

	oList := packetList(packets)
	sort.Sort(oList)
	return oList
}

type packetList []*Packet

func (p packetList) Len() int {
	return len(p)
}

func (p packetList) Less(i, j int) bool {
	return p[i].compare(*p[j]) < 0
}

func (p packetList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
