package day14

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCave(t *testing.T) {
	cave := Cave{
		xOffset: 494,
		grid: [][]rune{
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("          "),
			[]rune("    #   ##"),
			[]rune("    #   # "),
			[]rune("  ###   # "),
			[]rune("        # "),
			[]rune("        # "),
			[]rune("######### "),
		},
	}

	tests := []struct {
		rocks    []Rock
		expected Cave
	}{
		{[]Rock{
			{[]Point{{498, 4}, {498, 6}, {496, 6}}},
			{[]Point{{503, 4}, {502, 4}, {502, 9}, {494, 9}}},
		}, cave},
		{[]Rock{
			{[]Point{{496, 6}, {498, 6}, {498, 4}}},
			{[]Point{{494, 9}, {502, 9}, {502, 4}, {503, 4}}},
		}, cave},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, NewCave(tt.rocks))
		})
	}
}

func TestNewRock(t *testing.T) {
	tests := []struct {
		input string
		rock  Rock
	}{
		{"498,4 -> 498,6 -> 496,6", Rock{[]Point{
			{498, 4},
			{498, 6},
			{496, 6},
		}}},
		{"503,4 -> 502,4 -> 502,9 -> 494,9", Rock{[]Point{
			{503, 4},
			{502, 4},
			{502, 9},
			{494, 9},
		}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.rock, NewRock(tt.input))
		})
	}
}

func TestNewPoint(t *testing.T) {
	tests := []struct {
		input string
		point Point
	}{
		{"498,4", Point{498, 4}},
		{"498,6", Point{498, 6}},
		{"496,6", Point{496, 6}},
		{"503,4", Point{503, 4}},
		{"502,4", Point{502, 4}},
		{"502,9", Point{502, 9}},
		{"494,9", Point{494, 9}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.point, NewPoint(tt.input))
		})
	}
}
