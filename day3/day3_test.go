package day3

import (
	"advent-of-code-2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 157, solvePart1(testValues))
	assert.Equal(t, 7597, solvePart1(utils.ReadLines("day3", "day-3-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 70, solvePart2(testValues))
	assert.Equal(t, 2607, solvePart2(utils.ReadLines("day3", "day-3-input.txt")))
}

func TestNewRuckSack(t *testing.T) {
	tests := []struct {
		input    string
		expected RuckSack
	}{
		{"aBcDeFgHiJ", RuckSack{
			[]rune{'a', 'B', 'c', 'D', 'e'},
			[]rune{'F', 'g', 'H', 'i', 'J'},
		}},
		{"aBcDeFgHia", RuckSack{
			[]rune{'a', 'B', 'c', 'D', 'e'},
			[]rune{'F', 'g', 'H', 'i', 'a'},
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, newRuckSack(tt.input))
		})
	}
}

func TestRuckSackMisplaced(t *testing.T) {
	tests := []struct {
		input    RuckSack
		expected []rune
	}{
		{RuckSack{
			[]rune{'a', 'B', 'c', 'D', 'e'},
			[]rune{'F', 'g', 'H', 'i', 'J'},
		}, nil},
		{RuckSack{
			[]rune{'a', 'B', 'c', 'D', 'e'},
			[]rune{'F', 'g', 'H', 'i', 'a'},
		}, []rune{'a'}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.misplaced())
		})
	}
}

func TestRuckSackItems(t *testing.T) {
	r := RuckSack{
		container1: []rune{'a', 'b', 'c'},
		container2: []rune{'A', 'B', 'C'},
	}
	assert.Equal(t, []rune{'a', 'b', 'c', 'A', 'B', 'C'}, r.items())
}

func TestRuckSackContainsItem(t *testing.T) {
	r := RuckSack{
		container1: []rune{'a', 'b', 'c'},
		container2: []rune{'A', 'B', 'C'},
	}
	tests := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{'z', false},
		{'A', true},
		{'Z', false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, r.containsItem(tt.input))
		})
	}
}

func TestGroupBadge(t *testing.T) {
	g := Group{[]RuckSack{
		{[]rune{'a', 'b', 'c'}, []rune{'A', 'B', 'C'}},
		{[]rune{'d', 'b', 'f'}, []rune{'D', 'E', 'F'}},
		{[]rune{'g', 'b', 'i'}, []rune{'G', 'H', 'I'}},
	}}
	assert.Equal(t, 'b', g.badge())
}

func TestPriority(t *testing.T) {
	tests := []struct {
		input    rune
		expected int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, priority(tt.input))
		})
	}
}
