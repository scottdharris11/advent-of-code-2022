package day6

import (
	"advent-of-code-2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1760, solvePart1(utils.ReadLines("day6", "day-6-input.txt")[0]))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 2974, solvePart2(utils.ReadLines("day6", "day-6-input.txt")[0]))
}

func TestFirstUniqueRange(t *testing.T) {
	tests := []struct {
		s         string
		rangeSize int
		expected  int
	}{
		{"abc", 4, -1},
		{"abababab", 4, -1},
		{"abababcd", 4, 8},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.expected, firstUniqueRange(tt.s, tt.rangeSize))
		})
	}
}

func TestDupCharacters(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"abcdd", true},
		{"abcd", false},
		{"abac", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.expected, dupCharacters(tt.s))
		})
	}
}
