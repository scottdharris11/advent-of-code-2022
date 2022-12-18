package day15

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	// assert.Equal(t, 0, solvePart2(testInput, 10))
	// assert.Equal(t, 0, solvePart2(utils.ReadLines("day15", "day-15-input.txt"), 2000000))
}

func TestCave_NoBeaconPoints(t *testing.T) {
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
			cave := NewCave(11)
			cave.noBeaconRanges[11] = tt.ranges
			assert.Equal(t, tt.expected, cave.NoBeaconPoints(11))
		})
	}
}

func TestCave_MarkBeacon(t *testing.T) {
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
			cave := NewCave(tt.beacon.y)
			b := make([]Range, len(before))
			copy(b, before)
			cave.noBeaconRanges[tt.beacon.y] = b
			cave.MarkBeacon(tt.beacon)
			assert.Equal(t, tt.after, cave.noBeaconRanges[tt.beacon.y])
		})
	}
}

func TestCave_MarkNoBeacon(t *testing.T) {
	cave := NewCave(11)
	sensor := NewSensor("Sensor at x=8, y=7: closest beacon is at x=2, y=10")
	cave.MarkNoBeacon(*sensor)

	// assert.Nil(t, cave.noBeaconRanges[-3])
	// assert.Equal(t, []Range{{8, 8}}, cave.noBeaconRanges[-2])
	// assert.Equal(t, []Range{{7, 9}}, cave.noBeaconRanges[-1])
	// assert.Equal(t, []Range{{6, 10}}, cave.noBeaconRanges[0])
	// assert.Equal(t, []Range{{5, 11}}, cave.noBeaconRanges[1])
	// assert.Equal(t, []Range{{4, 12}}, cave.noBeaconRanges[2])
	// assert.Equal(t, []Range{{3, 13}}, cave.noBeaconRanges[3])
	// assert.Equal(t, []Range{{2, 14}}, cave.noBeaconRanges[4])
	// assert.Equal(t, []Range{{1, 15}}, cave.noBeaconRanges[5])
	// assert.Equal(t, []Range{{0, 16}}, cave.noBeaconRanges[6])
	// assert.Equal(t, []Range{{-1, 17}}, cave.noBeaconRanges[7])
	// assert.Equal(t, []Range{{0, 16}}, cave.noBeaconRanges[8])
	// assert.Equal(t, []Range{{1, 15}}, cave.noBeaconRanges[9])
	// assert.Equal(t, []Range{{2, 14}}, cave.noBeaconRanges[10])
	assert.Equal(t, []Range{{3, 13}}, cave.noBeaconRanges[11])
	// assert.Equal(t, []Range{{4, 12}}, cave.noBeaconRanges[12])
	// assert.Equal(t, []Range{{5, 11}}, cave.noBeaconRanges[13])
	// assert.Equal(t, []Range{{6, 10}}, cave.noBeaconRanges[14])
	// assert.Equal(t, []Range{{7, 9}}, cave.noBeaconRanges[15])
	// assert.Equal(t, []Range{{8, 8}}, cave.noBeaconRanges[16])
	// assert.Nil(t, cave.noBeaconRanges[17])
}

func TestCave_RecordRange(t *testing.T) {
	before := []Range{
		{10, 15},
		{20, 25},
		{40, 45},
	}
	tests := []struct {
		start int
		end   int
		after []Range
	}{
		{5, 7, []Range{{5, 7}, {10, 15}, {20, 25}, {40, 45}}},
		{50, 52, []Range{{10, 15}, {20, 25}, {40, 45}, {50, 52}}},
		{5, 10, []Range{{5, 15}, {20, 25}, {40, 45}}},
		{5, 11, []Range{{5, 15}, {20, 25}, {40, 45}}},
		{18, 23, []Range{{10, 15}, {18, 25}, {40, 45}}},
		{5, 18, []Range{{5, 18}, {20, 25}, {40, 45}}},
		{5, 18, []Range{{5, 18}, {20, 25}, {40, 45}}},
		{5, 23, []Range{{5, 25}, {40, 45}}},
		{5, 26, []Range{{5, 26}, {40, 45}}},
		{5, 50, []Range{{5, 50}}},
		{12, 15, []Range{{10, 15}, {20, 25}, {40, 45}}},
		{12, 18, []Range{{10, 18}, {20, 25}, {40, 45}}},
		{12, 30, []Range{{10, 30}, {40, 45}}},
		{12, 50, []Range{{10, 50}}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			cave := NewCave(3)
			b := make([]Range, len(before))
			copy(b, before)
			cave.noBeaconRanges[3] = b
			cave.recordRange(3, tt.start, tt.end)
			assert.Equal(t, tt.after, cave.noBeaconRanges[3])
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
