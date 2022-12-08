package day7

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day7", "day-7-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	tree := parseTree(lines)
	ans := sumDirectoriesUnderSize(tree, 100000)
	end := time.Now().UnixMilli()
	log.Printf("Day 7, Part 1 (%dms): Dir Sum = %d", end-start, ans)
	return ans
}

func sumDirectoriesUnderSize(dir *Item, limit int) int {
	ans := 0
	s := dir.Size()
	if s < limit {
		ans += s
	}
	for _, c := range dir.children {
		if c.directory {
			ans += sumDirectoriesUnderSize(c, limit)
		}
	}
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	tree := parseTree(lines)
	total := 70000000
	used := tree.Size()
	free := 30000000 - (total - used)
	dir := findSmallestDirToDelete(tree, free, tree)
	ans := dir.Size()
	end := time.Now().UnixMilli()
	log.Printf("Day 7, Part 2 (%dms): Smallest Dir = %d", end-start, ans)
	return ans
}

func findSmallestDirToDelete(dir *Item, free int, smallest *Item) *Item {
	r := smallest
	s := dir.Size()
	if s < free {
		return r
	}
	if s > free && s < smallest.Size() {
		r = dir
	}
	for _, c := range dir.children {
		if c.directory {
			r = findSmallestDirToDelete(c, free, r)
		}
	}
	return r
}

type Item struct {
	name      string
	size      int
	directory bool
	parent    *Item
	children  map[string]*Item
}

func (i *Item) Size() int {
	s := i.size
	if i.directory {
		for _, c := range i.children {
			s += c.Size()
		}
	}
	return s
}

func (i *Item) Directory(name string) *Item {
	if name == ".." {
		return i.parent
	}
	return i.children[name]
}

func parseTree(lines []string) *Item {
	root := &Item{
		name:      "/",
		directory: true,
		children:  map[string]*Item{},
	}
	var currentDir *Item
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "$ cd "):
			d := line[strings.LastIndex(line, " ")+1:]
			switch d {
			case "/":
				currentDir = root
			default:
				currentDir = currentDir.Directory(d)
			}
		case line == "$ ls":
			continue
		default:
			s := strings.Split(line, " ")
			switch s[0] {
			case "dir":
				currentDir.children[s[1]] = &Item{
					name:      s[1],
					directory: true,
					parent:    currentDir,
					children:  map[string]*Item{},
				}
			default:
				currentDir.children[s[1]] = &Item{
					name:   s[1],
					parent: currentDir,
					size:   utils.Number(s[0]),
				}
			}
		}
	}
	return root
}
