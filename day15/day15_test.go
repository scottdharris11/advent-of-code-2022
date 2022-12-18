package day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
	"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
	"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
	"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
	"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
	"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
	"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
	"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
	"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
	"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
	"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
	"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
	"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
	"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 26, solvePart1(testInput, 10))
	assert.Equal(t, 5607466, solvePart1(utils.ReadLines("day15", "day-15-input.txt"), 2000000))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 56000011, solvePart2(testInput, 20))
	assert.Equal(t, 12543202766584, solvePart2(utils.ReadLines("day15", "day-15-input.txt"), 4000000))
}

func TestNoBeaconPoints(t *testing.T) {
	tests := []struct {
		ranges   []Range
		expected int
	}{
		{nil, 0},
		{[]Range{
			{-10, -5},
			{-3, 3},
			{10, 15},
		}, 19},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.ranges), func(t *testing.T) {
			assert.Equal(t, tt.expected, NoBeaconPoints(tt.ranges))
		})
	}
}

func TestMarkBeacon(t *testing.T) {
	before := []Range{
		{10, 15},
		{20, 25},
		{40, 45},
	}
	tests := []struct {
		beacon Point
		after  []Range
	}{
		{Point{9, 11}, before},
		{Point{16, 11}, before},
		{Point{46, 11}, before},
		{Point{10, 11}, []Range{{11, 15}, {20, 25}, {40, 45}}},
		{Point{15, 11}, []Range{{10, 14}, {20, 25}, {40, 45}}},
		{Point{22, 11}, []Range{{10, 15}, {20, 21}, {23, 25}, {40, 45}}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			b := make([]Range, len(before))
			copy(b, before)
			after := MarkBeacon(tt.beacon, tt.beacon.y, b)
			assert.Equal(t, tt.after, after)
		})
	}
}

func TestMarkNoBeacon(t *testing.T) {
	sensor := NewSensor("Sensor at x=8, y=7: closest beacon is at x=2, y=10")
	ranges := MarkNoBeacon(*sensor, 11, 0, nil)
	assert.Equal(t, []Range{{3, 13}}, ranges)
}

func TestRecordRange(t *testing.T) {
	before := []Range{
		{10, 15},
		{20, 25},
		{40, 45},
	}
	tests := []struct {
		start    int
		end      int
		maxCoord int
		after    []Range
	}{
		{5, 7, 0, []Range{{5, 7}, {10, 15}, {20, 25}, {40, 45}}},
		{50, 52, 0, []Range{{10, 15}, {20, 25}, {40, 45}, {50, 52}}},
		{5, 10, 0, []Range{{5, 15}, {20, 25}, {40, 45}}},
		{5, 11, 0, []Range{{5, 15}, {20, 25}, {40, 45}}},
		{18, 23, 0, []Range{{10, 15}, {18, 25}, {40, 45}}},
		{5, 18, 0, []Range{{5, 18}, {20, 25}, {40, 45}}},
		{5, 18, 0, []Range{{5, 18}, {20, 25}, {40, 45}}},
		{5, 23, 0, []Range{{5, 25}, {40, 45}}},
		{5, 26, 0, []Range{{5, 26}, {40, 45}}},
		{5, 50, 0, []Range{{5, 50}}},
		{12, 15, 0, []Range{{10, 15}, {20, 25}, {40, 45}}},
		{12, 18, 0, []Range{{10, 18}, {20, 25}, {40, 45}}},
		{12, 30, 0, []Range{{10, 30}, {40, 45}}},
		{12, 50, 0, []Range{{10, 50}}},
		{55, 65, 50, before},
		{-4, -2, 50, before},
		{-4, 5, 50, []Range{{0, 5}, {10, 15}, {20, 25}, {40, 45}}},
		{49, 55, 50, []Range{{10, 15}, {20, 25}, {40, 45}, {49, 50}}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			b := make([]Range, len(before))
			copy(b, before)
			after := recordRange(b, tt.maxCoord, tt.start, tt.end)
			assert.Equal(t, tt.after, after)
		})
	}
}

func TestPoint_ManhattanDistance(t *testing.T) {
	tests := []struct {
		point1   Point
		point2   Point
		expected int
	}{
		{Point{2, 18}, Point{-2, 15}, 7},
		{Point{9, 16}, Point{10, 16}, 1},
		{Point{8, 7}, Point{2, 10}, 9},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v vs %v", tt.point1, tt.point2), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.point1.ManhattanDistance(tt.point2))
			assert.Equal(t, tt.expected, tt.point2.ManhattanDistance(tt.point1))
		})
	}
}

func TestNewSensor(t *testing.T) {
	tests := []struct {
		input    string
		expected Sensor
	}{
		{testInput[0], Sensor{Point{2, 18}, Point{-2, 15}}},
		{testInput[1], Sensor{Point{9, 16}, Point{10, 16}}},
		{testInput[2], Sensor{Point{13, 2}, Point{15, 3}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, *NewSensor(tt.input))
		})
	}
}
