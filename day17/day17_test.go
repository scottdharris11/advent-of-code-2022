package day17

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 3068, solvePart1(testInput))
	assert.Equal(t, 0, solvePart1(utils.ReadLines("day17", "day-17-input.txt")[0]))
}

func TestRock_CanShiftLeft(t *testing.T) {
	cross := Rock{height: 3, width: 3, pattern: []string{
		" # ",
		"###",
		" # ",
	}}
	angle := Rock{height: 3, width: 3, pattern: []string{
		"  #",
		"  #",
		"###",
	}}

	tests := []struct {
		rock    Rock
		rows    []string
		xOffset int
		left    bool
	}{
		{cross, []string{
			"##     ",
			"#      ",
			"##     ",
			"#######",
		}, 1, false},
		{cross, []string{
			"##     ",
			"#      ",
			"##     ",
			"#######",
		}, 2, true},
		{cross, []string{
			"#      ",
			"##     ",
			"#      ",
			"#######",
		}, 2, false},
		{cross, []string{
			"##     ",
			"#      ",
			"##     ",
			"#######",
		}, 3, true},
		{angle, []string{
			"###    ",
			"##     ",
			"       ",
			"#######",
		}, 1, false},
		{angle, []string{
			"##     ",
			"##     ",
			"       ",
			"#######",
		}, 1, true},
		{angle, []string{
			"##     ",
			"###    ",
			"       ",
			"#######",
		}, 1, false},
		{angle, []string{
			"##     ",
			"##     ",
			"#      ",
			"#######",
		}, 1, false},
		{angle, []string{
			"##     ",
			"##     ",
			"#      ",
			"#######",
		}, 2, true},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.left, tt.rock.CanShiftLeft(tt.xOffset, tt.rows))
		})
	}
}

func TestRock_CanShiftRight(t *testing.T) {
	cross := Rock{height: 3, width: 3, pattern: []string{
		" # ",
		"###",
		" # ",
	}}

	tests := []struct {
		rock    Rock
		rows    []string
		xOffset int
		right   bool
	}{
		{cross, []string{
			"     ##",
			"      #",
			"     ##",
			"#######",
		}, 1, true},
		{cross, []string{
			"     ##",
			"      #",
			"     ##",
			"#######",
		}, 2, true},
		{cross, []string{
			"      #",
			"     ##",
			"      #",
			"#######",
		}, 2, false},
		{cross, []string{
			"     ##",
			"      #",
			"    ###",
			"#######",
		}, 2, false},
		{cross, []string{
			"    ###",
			"      #",
			"     ##",
			"#######",
		}, 2, false},
		{cross, []string{
			"     ##",
			"      #",
			"     ##",
			"#######",
		}, 3, false},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.right, tt.rock.CanShiftRight(tt.xOffset, tt.rows))
		})
	}
}

func TestRock_CanShiftDown(t *testing.T) {
	rock := Rock{height: 3, width: 3, pattern: []string{
		" # ",
		"###",
		" # ",
	}}
	rows := []string{
		"       ",
		"       ",
		"       ",
		"     ##",
	}

	tests := []struct {
		rows    []string
		xOffset int
		down    bool
	}{
		{rows, 0, true},
		{rows, 1, true},
		{rows, 2, true},
		{rows, 3, true},
		{[]string{
			"       ",
			"       ",
			"     # ",
			"     ##",
		}, 3, false},
		{rows, 4, false},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.down, rock.CanShiftDown(tt.xOffset, tt.rows))
		})
	}
}

func TestCave_WindDirection(t *testing.T) {
	cave := NewCave("<<>><", RockPatterns)
	assert.Equal(t, Left, cave.windDirection())
	assert.Equal(t, Left, cave.windDirection())
	assert.Equal(t, Right, cave.windDirection())
	assert.Equal(t, Right, cave.windDirection())
	assert.Equal(t, Left, cave.windDirection())

	assert.Equal(t, Left, cave.windDirection())
	assert.Equal(t, Left, cave.windDirection())
	assert.Equal(t, Right, cave.windDirection())
	assert.Equal(t, Right, cave.windDirection())
	assert.Equal(t, Left, cave.windDirection())

	assert.Equal(t, Left, cave.windDirection())
	assert.Equal(t, Left, cave.windDirection())
	assert.Equal(t, Right, cave.windDirection())
	assert.Equal(t, Right, cave.windDirection())
	assert.Equal(t, Left, cave.windDirection())
}

func TestCave_RockPattern(t *testing.T) {
	cave := NewCave("", RockPatterns)
	assert.Equal(t, RockPatterns[0], cave.rock())
	assert.Equal(t, RockPatterns[1], cave.rock())
	assert.Equal(t, RockPatterns[2], cave.rock())
	assert.Equal(t, RockPatterns[3], cave.rock())
	assert.Equal(t, RockPatterns[4], cave.rock())

	assert.Equal(t, RockPatterns[0], cave.rock())
	assert.Equal(t, RockPatterns[1], cave.rock())
	assert.Equal(t, RockPatterns[2], cave.rock())
	assert.Equal(t, RockPatterns[3], cave.rock())
	assert.Equal(t, RockPatterns[4], cave.rock())

	assert.Equal(t, RockPatterns[0], cave.rock())
	assert.Equal(t, RockPatterns[1], cave.rock())
	assert.Equal(t, RockPatterns[2], cave.rock())
	assert.Equal(t, RockPatterns[3], cave.rock())
	assert.Equal(t, RockPatterns[4], cave.rock())
}

func TestCave_DropRock(t *testing.T) {
	cave := NewCave(testInput, RockPatterns)

	height := cave.dropRock()
	assert.Equal(t, 1, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 4, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"   #   ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 6, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"####   ",
		"  #    ",
		"  #    ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 7, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 9, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
		"    ## ",
		"    ## ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 10, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
		"    ## ",
		"    ## ",
		" ####  ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 13, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
		"    ## ",
		"    ## ",
		" ####  ",
		"  #    ",
		" ###   ",
		"  #    ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 15, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
		"    ## ",
		"    ## ",
		" ####  ",
		"  #    ",
		" ###   ",
		"  #### ",
		"     # ",
		"     # ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 17, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
		"    ## ",
		"    ## ",
		" ####  ",
		"  #    ",
		" ###   ",
		"  #### ",
		"    ## ",
		"    ## ",
		"    #  ",
		"    #  ",
	}, cave.rows)

	height = cave.dropRock()
	assert.Equal(t, 17, height)
	assert.Equal(t, []string{
		"#######",
		"  #### ",
		"   #   ",
		"  ###  ",
		"#####  ",
		"  # #  ",
		"  # #  ",
		"    #  ",
		"    ## ",
		"    ## ",
		" ####  ",
		"  #    ",
		" ###   ",
		"###### ",
		"##  ## ",
		"    ## ",
		"    #  ",
		"    #  ",
	}, cave.rows)
}

func TestCave_ImpactRows(t *testing.T) {
	cave := NewCave("", RockPatterns)
	cave.rows = append(cave.rows, " ####  ")
	cave.rows = append(cave.rows, "   ####")
	cave.rows = append(cave.rows, "####   ")

	tests := []struct {
		yOffset  int
		height   int
		expected []string
	}{
		{8, 4, []string{
			"       ",
			"       ",
			"       ",
			"       ",
			"       ",
		}},
		{4, 4, []string{
			"       ",
			"       ",
			"       ",
			"       ",
			"####   ",
		}},
		{4, 1, []string{
			"       ",
			"####   ",
		}},
		{3, 4, []string{
			"       ",
			"       ",
			"       ",
			"####   ",
			"   ####",
		}},
		{2, 4, []string{
			"       ",
			"       ",
			"####   ",
			"   ####",
			" ####  ",
		}},
		{1, 4, []string{
			"       ",
			"####   ",
			"   ####",
			" ####  ",
			"#######",
		}},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.expected, cave.impactRows(tt.yOffset, tt.height))
		})
	}
}

func TestCave_AddRock(t *testing.T) {
	tests := []struct {
		before  []string
		yOffset int
		xOffset int
		rock    Rock
		after   []string
		size    int
	}{
		{[]string{
			"#######",
		}, 1, 2, Rock{height: 2, width: 2, pattern: []string{
			"##",
			"##",
		}}, []string{
			"#######",
			"  ##   ",
			"  ##   ",
		}, 2},
		{[]string{
			"      #",
			"      #",
			"      #",
			"      #",
			"#######",
		}, 1, 2, Rock{height: 2, width: 2, pattern: []string{
			"##",
			"##",
		}}, []string{
			"#######",
			"  ##  #",
			"  ##  #",
			"      #",
			"      #",
		}, 4},
		{[]string{
			"      #",
			"      #",
			"      #",
			"      #",
			"#######",
		}, 1, 2, Rock{height: 3, width: 3, pattern: []string{
			"  #",
			"  #",
			"###",
		}}, []string{
			"#######",
			"  ### #",
			"    # #",
			"    # #",
			"      #",
		}, 4},
		{[]string{
			"     ##",
			"      #",
			"     ##",
			"#######",
		}, 1, 3, Rock{height: 3, width: 3, pattern: []string{
			" # ",
			"###",
			" # ",
		}}, []string{
			"#######",
			"    ###",
			"   ####",
			"    ###",
		}, 3},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			cave := Cave{width: 7}
			for i := len(tt.before) - 1; i >= 0; i-- {
				cave.rows = append(cave.rows, tt.before[i])
			}
			cave.addRock(tt.xOffset, tt.yOffset, tt.rock)
			assert.Equal(t, tt.after, cave.rows)
			assert.Equal(t, tt.size, cave.topRock())
		})
	}
}
