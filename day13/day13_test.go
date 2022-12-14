package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestNewPacket(t *testing.T) {
	tests := []struct {
		input  string
		output Packet
	}{
		{"[]", Packet{value: -1}},
		{"[3]", Packet{value: 3}},
		{"[1,1,3,1,1]", Packet{value: -1, values: []*Packet{
			{value: 1},
			{value: 1},
			{value: 3},
			{value: 1},
			{value: 1},
		}}},
		{"[[1],[2,3,4]]", Packet{value: -1, values: []*Packet{
			{value: 1},
			{value: -1, values: []*Packet{
				{value: 2},
				{value: 3},
				{value: 4},
			}},
		}}},
		{"[[1],4]", Packet{value: -1, values: []*Packet{
			{value: 1},
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
