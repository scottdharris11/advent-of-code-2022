package day18

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2022/utils"
)

var testInput = []string{
	"2,2,2",
	"1,2,2",
	"3,2,2",
	"2,1,2",
	"2,3,2",
	"2,2,1",
	"2,2,3",
	"2,2,4",
	"2,2,6",
	"1,2,5",
	"3,2,5",
	"2,1,5",
	"2,3,5",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 64, solvePart1(testInput))
	assert.Equal(t, 3470, solvePart1(utils.ReadLines("day18", "day-18-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 58, solvePart2(testInput))
	assert.Equal(t, 1986, solvePart2(utils.ReadLines("day18", "day-18-input.txt")))
}

func TestCube_MarkAdjacent(t *testing.T) {
	tests := []struct {
		cube1    *Cube
		cube2    *Cube
		adjacent []bool
	}{
		{
			&Cube{x: 1, y: 1, z: 1},
			&Cube{x: 2, y: 1, z: 1},
			[]bool{false, true, false, false, false},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			tt.cube1.MarkAdjacent(tt.cube2)
			for i, a := range tt.adjacent {
				if a {
					switch i {
					case 0:
						assert.Equal(t, tt.cube2, tt.cube1.adjacent[0])
						assert.Equal(t, tt.cube1, tt.cube2.adjacent[1])
					case 1:
						assert.Equal(t, tt.cube2, tt.cube1.adjacent[1])
						assert.Equal(t, tt.cube1, tt.cube2.adjacent[0])
					case 2:
						assert.Equal(t, tt.cube2, tt.cube1.adjacent[2])
						assert.Equal(t, tt.cube1, tt.cube2.adjacent[3])
					case 3:
						assert.Equal(t, tt.cube2, tt.cube1.adjacent[3])
						assert.Equal(t, tt.cube1, tt.cube2.adjacent[2])
					case 4:
						assert.Equal(t, tt.cube2, tt.cube1.adjacent[4])
						assert.Equal(t, tt.cube1, tt.cube2.adjacent[5])
					case 5:
						assert.Equal(t, tt.cube2, tt.cube1.adjacent[5])
						assert.Equal(t, tt.cube1, tt.cube2.adjacent[4])
					}
				}
			}
		})
	}
}

func TestCube_ExposedSides(t *testing.T) {
	tests := []struct {
		cube     Cube
		expected int
	}{
		{Cube{}, 6},
		{Cube{adjacent: [6]*Cube{{}, nil, nil, nil, nil, nil}}, 5},
		{Cube{adjacent: [6]*Cube{{}, {}, nil, nil, nil, nil}}, 4},
		{Cube{adjacent: [6]*Cube{{}, {}, {}, nil, nil, nil}}, 3},
		{Cube{adjacent: [6]*Cube{{}, {}, {}, {}, nil, nil}}, 2},
		{Cube{adjacent: [6]*Cube{{}, {}, {}, {}, {}, nil}}, 1},
		{Cube{adjacent: [6]*Cube{{}, {}, {}, {}, {}, {}}}, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.cube.ExposedSides())
		})
	}
}
