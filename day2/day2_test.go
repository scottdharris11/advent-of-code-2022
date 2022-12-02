package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testValues = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 15, solvePart1(testValues))
	assert.Equal(t, 10310, solvePart1(utils.ReadLines("day2", "day-2-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 12, solvePart2(testValues))
	assert.Equal(t, 14859, solvePart2(utils.ReadLines("day2", "day-2-input.txt")))
}

func TestScore(t *testing.T) {
	tests := []struct {
		play     Play
		opp      Play
		expected int
	}{
		{Rock, Rock, 4},
		{Rock, Paper, 1},
		{Rock, Scissors, 7},
		{Paper, Rock, 8},
		{Paper, Paper, 5},
		{Paper, Scissors, 2},
		{Scissors, Rock, 3},
		{Scissors, Paper, 9},
		{Scissors, Scissors, 6},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d vs %d", tt.play, tt.opp), func(t *testing.T) {
			assert.Equal(t, tt.expected, score(tt.play, tt.opp))
		})
	}
}

func TestRuneToPlay(t *testing.T) {
	tests := []struct {
		input    rune
		expected Play
	}{
		{'A', Rock},
		{'B', Paper},
		{'C', Scissors},
		{'X', Rock},
		{'Y', Paper},
		{'Z', Scissors},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, runeToPlay(tt.input))
		})
	}
}

func TestRuneToOutcome(t *testing.T) {
	tests := []struct {
		input    rune
		expected int
	}{
		{'X', Loss},
		{'Y', Tie},
		{'Z', Win},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, runeToOutcome(tt.input))
		})
	}
}

func TestPlayForOutcome(t *testing.T) {
	tests := []struct {
		oppPlay  Play
		outcome  int
		expected Play
	}{
		{Rock, Win, Paper},
		{Rock, Tie, Rock},
		{Rock, Loss, Scissors},
		{Paper, Win, Scissors},
		{Paper, Tie, Paper},
		{Paper, Loss, Rock},
		{Scissors, Win, Rock},
		{Scissors, Tie, Scissors},
		{Scissors, Loss, Paper},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d against %d", tt.outcome, tt.oppPlay), func(t *testing.T) {
			assert.Equal(t, tt.expected, playForOutcome(tt.oppPlay, tt.outcome))
		})
	}
}
