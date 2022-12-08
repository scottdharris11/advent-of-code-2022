package day7

import (
	"advent-of-code-2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLines = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 95437, solvePart1(testLines))
	assert.Equal(t, 1723892, solvePart1(utils.ReadLines("day7", "day-7-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 24933642, solvePart2(testLines))
	assert.Equal(t, 8474158, solvePart2(utils.ReadLines("day7", "day-7-input.txt")))
}

func TestItem_Size(t *testing.T) {
	tests := []struct {
		name     string
		item     Item
		expected int
	}{
		{"just a file", Item{size: 300}, 300},
		{"directory with files", Item{directory: true, children: map[string]*Item{
			"a": {size: 200},
			"b": {size: 100},
			"c": {size: 400},
		}}, 700},
		{"nested directories", Item{directory: true, children: map[string]*Item{
			"a": {size: 200},
			"b": {size: 100},
			"c": {size: 400},
			"d": {directory: true, children: map[string]*Item{
				"e": {size: 300},
				"f": {size: 900},
			}},
		}}, 1900},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.item.Size())
		})
	}
}

func TestItem_Directory(t *testing.T) {
	root := &Item{name: "/", directory: true}
	dir1 := &Item{name: "dir1", directory: true, parent: root}
	dir2 := &Item{name: "dir2", directory: true, parent: root}
	dir3 := &Item{name: "dir3", directory: true, parent: dir1}
	root.children = map[string]*Item{
		"a.txt": {name: "a.txt", size: 100, parent: root},
		"dir1":  dir1,
		"dir2":  dir2,
	}
	dir1.children = map[string]*Item{
		"b.txt": {name: "b.txt", size: 200, parent: dir1},
		"dir3":  dir3,
		"c.txt": {name: "c.txt", size: 300, parent: dir1},
	}
	dir2.children = map[string]*Item{
		"d.txt": {name: "d.txt", size: 400, parent: dir2},
	}

	assert.Equal(t, dir1, root.Directory("dir1"))
	assert.Equal(t, root, dir1.Directory(".."))
	assert.Equal(t, dir2, root.Directory("dir2"))
	assert.Equal(t, dir3, dir1.Directory("dir3"))
	assert.Equal(t, dir1, dir3.Directory(".."))
	assert.Nil(t, dir3.Directory("dir4"))
}

func TestParseTree(t *testing.T) {
	root := parseTree(testLines)

	assert.Equal(t, "/", root.name)
	assert.True(t, root.directory)
	assert.Equal(t, 4, len(root.children))

	assert.Equal(t, "a", root.children["a"].name)
	assert.True(t, root.children["a"].directory)
	assert.Equal(t, 4, len(root.children["a"].children))

	assert.Equal(t, "e", root.children["a"].children["e"].name)
	assert.True(t, root.children["a"].children["e"].directory)
	assert.Equal(t, 1, len(root.children["a"].children["e"].children))

	assert.Equal(t, "i", root.children["a"].children["e"].children["i"].name)
	assert.False(t, root.children["a"].children["e"].children["i"].directory)
	assert.Equal(t, 584, root.children["a"].children["e"].children["i"].size)

	assert.Equal(t, "f", root.children["a"].children["f"].name)
	assert.False(t, root.children["a"].children["f"].directory)
	assert.Equal(t, 29116, root.children["a"].children["f"].size)

	assert.Equal(t, "g", root.children["a"].children["g"].name)
	assert.False(t, root.children["a"].children["g"].directory)
	assert.Equal(t, 2557, root.children["a"].children["g"].size)

	assert.Equal(t, "h.lst", root.children["a"].children["h.lst"].name)
	assert.False(t, root.children["a"].children["h.lst"].directory)
	assert.Equal(t, 62596, root.children["a"].children["h.lst"].size)

	assert.Equal(t, "b.txt", root.children["b.txt"].name)
	assert.False(t, root.children["b.txt"].directory)
	assert.Equal(t, 14848514, root.children["b.txt"].size)

	assert.Equal(t, "c.dat", root.children["c.dat"].name)
	assert.False(t, root.children["c.dat"].directory)
	assert.Equal(t, 8504156, root.children["c.dat"].size)

	assert.Equal(t, "d", root.children["d"].name)
	assert.True(t, root.children["d"].directory)
	assert.Equal(t, 4, len(root.children["d"].children))

	assert.Equal(t, "j", root.children["d"].children["j"].name)
	assert.False(t, root.children["d"].children["j"].directory)
	assert.Equal(t, 4060174, root.children["d"].children["j"].size)

	assert.Equal(t, "d.log", root.children["d"].children["d.log"].name)
	assert.False(t, root.children["d"].children["d.log"].directory)
	assert.Equal(t, 8033020, root.children["d"].children["d.log"].size)

	assert.Equal(t, "d.ext", root.children["d"].children["d.ext"].name)
	assert.False(t, root.children["d"].children["d.ext"].directory)
	assert.Equal(t, 5626152, root.children["d"].children["d.ext"].size)

	assert.Equal(t, "k", root.children["d"].children["k"].name)
	assert.False(t, root.children["d"].children["k"].directory)
	assert.Equal(t, 7214296, root.children["d"].children["k"].size)
}
