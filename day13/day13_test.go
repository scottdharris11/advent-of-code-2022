package day13

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testLines = []string{
	"[1,1,3,1,1]",
	"[1,1,5,1,1]",
	"",
	"[[1],[2,3,4]]",
	"[[1],4]",
	"",
	"[9]",
	"[[8,7,6]]",
	"",
	"[[4,4],4,4]",
	"[[4,4],4,4,4]",
	"",
	"[7,7,7,7]",
	"[7,7,7]",
	"",
	"[]",
	"[3]",
	"",
	"[[[]]]",
	"[[]]",
	"",
	"[1,[2,[3,[4,[5,6,7]]]],8,9]",
	"[1,[2,[3,[4,[5,6,0]]]],8,9]",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 13, solvePart1(testLines))
	assert.Equal(t, 6187, solvePart1(utils.ReadLines("day13", "day-13-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 140, solvePart2(testLines))
	assert.Equal(t, 23520, solvePart2(utils.ReadLines("day13", "day-13-input.txt")))
}

func TestPacket_Ordered(t *testing.T) {
	tests := []struct {
		left    Packet
		right   Packet
		ordered bool
	}{
		{*NewPacket("[1,1,5,1,1]"), *NewPacket("[1,1,5,1,1]"), false},
		{*NewPacket("[1,1,3,1,1]"), *NewPacket("[1,1,5,1,1]"), true},
		{*NewPacket("[[1],[2,3,4]]"), *NewPacket("[[1],4]"), true},
		{*NewPacket("[9]"), *NewPacket("[[8,7,6]]"), false},
		{*NewPacket("[[4,4],4,4]"), *NewPacket("[[4,4],4,4,4]"), true},
		{*NewPacket("[7,7,7,7]"), *NewPacket("[7,7,7]"), false},
		{*NewPacket("[]"), *NewPacket("[3]"), true},
		{*NewPacket("[[[]]]"), *NewPacket("[[]]"), false},
		{*NewPacket("[1,[2,[3,[4,[5,6,7]]]],8,9]"), *NewPacket("[1,[2,[3,[4,[5,6,0]]]],8,9]"), false},
		{*NewPacket("[[[2,3],3],[10,4,[[10,8,4,6,9],5],[8,[0,9,5,5,4],3]]]"), *NewPacket("[[0],[[[1]]],[1,9,2,2,[[4],4,8]],[]]"), false},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.ordered, tt.left.Ordered(tt.right))
		})
	}
}

func TestNewPacket(t *testing.T) {
	tests := []struct {
		input  string
		output Packet
	}{
		{"[]", Packet{value: -1}},
		{"[3]", Packet{value: -1, values: []*Packet{
			{value: 3},
		}}},
		{"[1,1,3,1,1]", Packet{value: -1, values: []*Packet{
			{value: 1},
			{value: 1},
			{value: 3},
			{value: 1},
			{value: 1},
		}}},
		{"[[1],[2,3,4]]", Packet{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 1},
			}},
			{value: -1, values: []*Packet{
				{value: 2},
				{value: 3},
				{value: 4},
			}},
		}}},
		{"[[1],4]", Packet{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 1},
			}},
			{value: 4},
		}}},
		{"[[8,7,6]]", Packet{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 8},
				{value: 7},
				{value: 6},
			}},
		}}},
		{"[[4,4],4,4]", Packet{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 4},
				{value: 4},
			}},
			{value: 4},
			{value: 4},
		}}},
		{"[[[]]]", Packet{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: -1},
			}},
		}}},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", Packet{value: -1, values: []*Packet{
			{value: 1},
			{value: -1, values: []*Packet{
				{value: 2},
				{value: -1, values: []*Packet{
					{value: 3},
					{value: -1, values: []*Packet{
						{value: 4},
						{value: -1, values: []*Packet{
							{value: 5},
							{value: 6},
							{value: 7},
						}},
					}},
				}},
			}},
			{value: 8},
			{value: 9},
		}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.output, *NewPacket(tt.input))
		})
	}
}

func TestOrderedPackets(t *testing.T) {
	assert.Equal(t, []*Packet{
		{value: -1},
		{value: -1, values: []*Packet{
			{value: -1},
		}},
		{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: -1},
			}},
		}},
		{value: -1, values: []*Packet{
			{value: 1},
			{value: 1},
			{value: 3},
			{value: 1},
			{value: 1},
		}},
		{value: -1, values: []*Packet{
			{value: 1},
			{value: 1},
			{value: 5},
			{value: 1},
			{value: 1},
		}},
		{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 1},
			}},
			{value: -1, values: []*Packet{
				{value: 2},
				{value: 3},
				{value: 4},
			}},
		}},
		{value: -1, values: []*Packet{
			{value: 1},
			{value: -1, values: []*Packet{
				{value: 2},
				{value: -1, values: []*Packet{
					{value: 3},
					{value: -1, values: []*Packet{
						{value: 4},
						{value: -1, values: []*Packet{
							{value: 5},
							{value: 6},
							{value: 0},
						}},
					}},
				}},
			}},
			{value: 8},
			{value: 9},
		}},
		{value: -1, values: []*Packet{
			{value: 1},
			{value: -1, values: []*Packet{
				{value: 2},
				{value: -1, values: []*Packet{
					{value: 3},
					{value: -1, values: []*Packet{
						{value: 4},
						{value: -1, values: []*Packet{
							{value: 5},
							{value: 6},
							{value: 7},
						}},
					}},
				}},
			}},
			{value: 8},
			{value: 9},
		}},
		{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 1},
			}},
			{value: 4},
		}},
		{value: -1, values: []*Packet{
			{value: 3},
		}},
		{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 4},
				{value: 4},
			}},
			{value: 4},
			{value: 4},
		}},
		{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 4},
				{value: 4},
			}},
			{value: 4},
			{value: 4},
			{value: 4},
		}},
		{value: -1, values: []*Packet{
			{value: 7},
			{value: 7},
			{value: 7},
		}},
		{value: -1, values: []*Packet{
			{value: 7},
			{value: 7},
			{value: 7},
			{value: 7},
		}},
		{value: -1, values: []*Packet{
			{value: -1, values: []*Packet{
				{value: 8},
				{value: 7},
				{value: 6},
			}},
		}},
		{value: -1, values: []*Packet{
			{value: 9},
		}},
	}, orderedPackets(testLines))
}
