package day13

import (
	"advent-of-code-2022/utils"
)

type Packet struct {
	value  int
	values []*Packet
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
				if p.value >= 0 {
					p.values = append(p.values, &Packet{value: p.value})
					p.value = -1
				}
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
				switch {
				case p.value == -1 && len(p.values) == 0:
					p.value = v
				case p.value >= 0:
					p.values = append(p.values, &Packet{value: p.value})
					p.value = -1
					fallthrough
				default:
					p.values = append(p.values, &Packet{value: v})
				}
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
