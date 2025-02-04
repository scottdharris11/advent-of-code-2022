package day25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"1=-0-2",
	"12111",
	"2=0=",
	"21",
	"2=01",
	"111",
	"20012",
	"112",
	"1=-1=",
	"1-12",
	"12",
	"1=",
	"122",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, "2=-1=0", solvePart1(testInput))
	assert.Equal(t, "2----0=--1122=0=0021", solvePart1(utils.ReadLines("day25", "day-25-input.txt")))
}

func TestStoI(t *testing.T) {
	tests := []struct {
		sanfu   string
		decimal int
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.sanfu, func(t *testing.T) {
			assert.Equal(t, tt.decimal, stoi(tt.sanfu))
		})
	}
}

func TestItoS(t *testing.T) {
	tests := []struct {
		snafu   string
		decimal int
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.decimal), func(t *testing.T) {
			assert.Equal(t, tt.snafu, itos(tt.decimal))
		})
	}
}
