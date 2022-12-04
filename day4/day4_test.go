package day4

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testValues = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 2, solvePart1(testValues))
	assert.Equal(t, 538, solvePart1(utils.ReadLines("day4", "day-4-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 4, solvePart2(testValues))
	assert.Equal(t, 792, solvePart2(utils.ReadLines("day4", "day-4-input.txt")))
}

func TestPairAssignmentFullOverlap(t *testing.T) {
	tests := []struct {
		pair     PairAssignment
		expected bool
	}{
		{PairAssignment{assign1: SectionAssignment{2, 4}, assign2: SectionAssignment{6, 8}}, false},
		{PairAssignment{assign1: SectionAssignment{6, 6}, assign2: SectionAssignment{4, 6}}, true},
		{PairAssignment{assign1: SectionAssignment{2, 8}, assign2: SectionAssignment{3, 7}}, true},
		{PairAssignment{assign1: SectionAssignment{5, 7}, assign2: SectionAssignment{7, 9}}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.pair.fullOverlap())
		})
	}
}

func TestPairAssignmentPartialOverlap(t *testing.T) {
	tests := []struct {
		pair     PairAssignment
		expected bool
	}{
		{PairAssignment{assign1: SectionAssignment{2, 4}, assign2: SectionAssignment{6, 8}}, false},
		{PairAssignment{assign1: SectionAssignment{6, 6}, assign2: SectionAssignment{4, 6}}, true},
		{PairAssignment{assign1: SectionAssignment{2, 8}, assign2: SectionAssignment{3, 7}}, true},
		{PairAssignment{assign1: SectionAssignment{5, 7}, assign2: SectionAssignment{7, 9}}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.pair.partialOverlap())
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		begin1   int
		end1     int
		begin2   int
		end2     int
		expected bool
	}{
		{2, 4, 6, 8, false},
		{2, 3, 4, 5, false},
		{5, 7, 7, 9, false},
		{2, 8, 3, 7, true},
		{3, 7, 2, 8, false},
		{6, 6, 4, 6, false},
		{4, 6, 6, 6, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.expected, contains(tt.begin1, tt.end1, tt.begin2, tt.end2))
		})
	}
}

func TestOverlap(t *testing.T) {
	tests := []struct {
		begin1   int
		end1     int
		begin2   int
		end2     int
		expected bool
	}{
		{2, 4, 6, 8, false},
		{2, 3, 4, 5, false},
		{5, 7, 7, 9, true},
		{7, 9, 5, 7, true},
		{2, 8, 3, 7, true},
		{3, 7, 2, 8, true},
		{6, 6, 4, 6, true},
		{4, 6, 6, 6, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.expected, overlap(tt.begin1, tt.end1, tt.begin2, tt.end2))
		})
	}
}

func TestParseAssignments(t *testing.T) {
	pairs := parseAssignments([]string{
		"2-4,6-8",
		"2-3,4-5",
	})
	assert.Equal(t, []PairAssignment{
		{
			assign1: SectionAssignment{min: 2, max: 4},
			assign2: SectionAssignment{min: 6, max: 8},
		},
		{
			assign1: SectionAssignment{min: 2, max: 3},
			assign2: SectionAssignment{min: 4, max: 5},
		},
	}, pairs)
}
