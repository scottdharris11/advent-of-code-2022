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
		{"directory with files", Item{directory: true, children: []*Item{
			{size: 200},
			{size: 100},
			{size: 400},
		}}, 700},
		{"nested directories", Item{directory: true, children: []*Item{
			{size: 200},
			{size: 100},
			{size: 400},
			{directory: true, children: []*Item{
				{size: 300},
				{size: 900},
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
	root.children = []*Item{
		{name: "a.txt", size: 100, parent: root},
		dir1,
		dir2,
	}
	dir1.children = []*Item{
		{name: "b.txt", size: 200, parent: dir1},
		dir3,
		{name: "c.txt", size: 300, parent: dir1},
	}
	dir2.children = []*Item{
		{name: "d.txt", size: 400, parent: dir2},
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

	assert.Equal(t, "a", root.children[0].name)
	assert.True(t, root.children[0].directory)
	assert.Equal(t, 4, len(root.children[0].children))

	assert.Equal(t, "e", root.children[0].children[0].name)
	assert.True(t, root.children[0].children[0].directory)
	assert.Equal(t, 1, len(root.children[0].children[0].children))

	assert.Equal(t, "i", root.children[0].children[0].children[0].name)
	assert.False(t, root.children[0].children[0].children[0].directory)
	assert.Equal(t, 584, root.children[0].children[0].children[0].size)

	assert.Equal(t, "f", root.children[0].children[1].name)
	assert.False(t, root.children[0].children[1].directory)
	assert.Equal(t, 29116, root.children[0].children[1].size)

	assert.Equal(t, "g", root.children[0].children[2].name)
	assert.False(t, root.children[0].children[2].directory)
	assert.Equal(t, 2557, root.children[0].children[2].size)

	assert.Equal(t, "h.lst", root.children[0].children[3].name)
	assert.False(t, root.children[0].children[3].directory)
	assert.Equal(t, 62596, root.children[0].children[3].size)

	assert.Equal(t, "b.txt", root.children[1].name)
	assert.False(t, root.children[1].directory)
	assert.Equal(t, 14848514, root.children[1].size)

	assert.Equal(t, "c.dat", root.children[2].name)
	assert.False(t, root.children[2].directory)
	assert.Equal(t, 8504156, root.children[2].size)

	assert.Equal(t, "d", root.children[3].name)
	assert.True(t, root.children[3].directory)
	assert.Equal(t, 4, len(root.children[3].children))

	assert.Equal(t, "j", root.children[3].children[0].name)
	assert.False(t, root.children[3].children[0].directory)
	assert.Equal(t, 4060174, root.children[3].children[0].size)

	assert.Equal(t, "d.log", root.children[3].children[1].name)
	assert.False(t, root.children[3].children[1].directory)
	assert.Equal(t, 8033020, root.children[3].children[1].size)

	assert.Equal(t, "d.ext", root.children[3].children[2].name)
	assert.False(t, root.children[3].children[2].directory)
	assert.Equal(t, 5626152, root.children[3].children[2].size)

	assert.Equal(t, "k", root.children[3].children[3].name)
	assert.False(t, root.children[3].children[3].directory)
	assert.Equal(t, 7214296, root.children[3].children[3].size)
}
