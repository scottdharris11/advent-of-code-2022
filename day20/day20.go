package day20

import (
	"log"
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

	list := utils.WrappedList{}
	var items []*utils.WrappedListItem
	for _, line := range lines {
		items = append(items, list.Add(utils.Number(line)))
	}

	for _, item := range items {
		list.Move(item, item.Value.(int))
	}

	zeroIdx := 0
	for _, value := range list.Items() {
		if value.(int) == 0 {
			break
		}
		zeroIdx++
	}

	var i1000 = list.ItemAtIndex(zeroIdx + 1000).(int)
	var i2000 = list.ItemAtIndex(zeroIdx + 2000).(int)
	var i3000 = list.ItemAtIndex(zeroIdx + 3000).(int)
	ans := i1000 + i2000 + i3000
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 1 (%dms): Coordinate = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := len(lines)
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 2 (%dms): Answer = %d", end-start, ans)
	return ans
}
