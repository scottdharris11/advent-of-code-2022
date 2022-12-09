package day9

import (
	"fmt"
	"log"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day9", "day-9-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	tail := NewKnot(nil)
	head := NewKnot(tail)
	for _, line := range lines {
		s := strings.Split(line, " ")
		move(head, s[0], utils.Number(s[1]))
	}
	ans := tail.PositionsVisited()
	end := time.Now().UnixMilli()
	log.Printf("Day 9, Part 1 (%dms): Tail Locations = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	tail := NewKnot(nil)
	prev := tail
	for i := 0; i < 9; i++ {
		prev = NewKnot(prev)
	}
	head := prev
	for _, line := range lines {
		s := strings.Split(line, " ")
		move(head, s[0], utils.Number(s[1]))
	}
	ans := tail.PositionsVisited()
	end := time.Now().UnixMilli()
	log.Printf("Day 9, Part 2 (%dms): Tail Locations = %d", end-start, ans)
	return ans
}

func move(knot *Knot, dir string, steps int) {
	for i := 0; i < steps; i++ {
		switch dir {
		case "U":
			knot.Move(1, 0)
		case "D":
			knot.Move(-1, 0)
		case "R":
			knot.Move(0, 1)
		case "L":
			knot.Move(0, -1)
		}
	}
}

func NewKnot(linked *Knot) *Knot {
	k := Knot{}
	k.linked = linked
	k.visited = make(map[string]int)
	k.visited["0:0"]++
	return &k
}

type Knot struct {
	x       int
	y       int
	visited map[string]int
	linked  *Knot
}

func (r *Knot) Move(upDown int, rightLeft int) {
	r.y += upDown
	r.x += rightLeft
	r.visited[fmt.Sprintf("%d:%d", r.x, r.y)]++
	if r.linked != nil {
		r.linked.Align(*r)
	}
}

func (r *Knot) PositionsVisited() int {
	return len(r.visited)
}

func (r *Knot) Align(h Knot) {
	adjacent, xDiff, yDiff := r.adjacent(h)
	if adjacent {
		return
	}

	switch {
	case xDiff == 0:
		yMove := yDiff - 1
		if yDiff < 0 {
			yMove = yDiff + 1
		}
		r.Move(yMove, 0)
	case yDiff == 0:
		xMove := xDiff - 1
		if xDiff < 0 {
			xMove = xDiff + 1
		}
		r.Move(0, xMove)
	case xDiff > 0 && yDiff > 0:
		r.Move(1, 1)
	case xDiff < 0 && yDiff < 0:
		r.Move(-1, -1)
	case xDiff > 0 && yDiff < 0:
		r.Move(-1, 1)
	case xDiff < 0 && yDiff > 0:
		r.Move(1, -1)
	}
}

func (r *Knot) adjacent(h Knot) (bool, int, int) {
	xDiff := h.x - r.x
	if h.x < r.x {
		xDiff = (r.x - h.x) * -1
	}
	yDiff := h.y - r.y
	if h.y < r.y {
		yDiff = (r.y - h.y) * -1
	}
	return xDiff >= -1 && xDiff <= 1 && yDiff >= -1 && yDiff <= 1, xDiff, yDiff
}
