package day20

import (
	"log"
	"math"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day20", "day-20-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	coords, zero := parseCoordinates(lines)
	count := len(coords)
	for _, c := range coords {
		move(c, count)
	}
	i1000 := valueAfter(zero, count, 1000)
	i2000 := valueAfter(zero, count, 2000)
	i3000 := valueAfter(zero, count, 3000)
	ans := i1000 + i2000 + i3000
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 1 (%dms): Coordinate = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	coords, zero := parseCoordinates(lines)
	count := len(coords)
	for i := 0; i < 10; i++ {
		for _, c := range coords {
			if i == 0 {
				c.value *= 811589153
			}
			move(c, count)
		}
	}
	i1000 := valueAfter(zero, count, 1000)
	i2000 := valueAfter(zero, count, 2000)
	i3000 := valueAfter(zero, count, 3000)
	ans := i1000 + i2000 + i3000
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 2 (%dms): Coordinate = %d", end-start, ans)
	return ans
}

type Coordinate struct {
	value int
	next  *Coordinate
	prev  *Coordinate
}

func parseCoordinates(lines []string) ([]*Coordinate, *Coordinate) {
	var coords []*Coordinate
	var zero *Coordinate

	first := &Coordinate{value: utils.Number(lines[0])}
	coords = append(coords, first)
	prev := first
	for i := 1; i < len(lines); i++ {
		next := &Coordinate{value: utils.Number(lines[i])}
		if next.value == 0 {
			zero = next
		}
		coords = append(coords, next)
		next.prev = prev
		prev.next = next
		prev = next
	}
	first.prev = prev
	prev.next = first
	return coords, zero
}

func move(c *Coordinate, count int) {
	places := int(math.Abs(float64(c.value)))
	prev := c.value < 0
	if prev {
		places++
	}
	// modulus of count minus 1 since we are moving "positions" and therefore shouldn't
	// count the moving value as an actual position (this was a real bugger)
	places %= count - 1
	if places == 0 {
		return
	}
	current := c
	for i := 0; i < places; i++ {
		if prev {
			current = current.prev
		} else {
			current = current.next
		}
	}

	c.prev.next = c.next
	c.next.prev = c.prev
	c.prev = current
	c.next = current.next
	c.next.prev = c
	c.prev.next = c
}

func valueAfter(c *Coordinate, count int, places int) int {
	current := c
	adjust := places % count
	for i := 0; i < adjust; i++ {
		current = current.next
	}
	return current.value
}
