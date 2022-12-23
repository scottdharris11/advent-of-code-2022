package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadLines read lines from data file as strings
func ReadLines(dir string, filename string) []string {
	file := openFile(dir, filename)
	defer closeFile(file)

	var values []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values
}

// ReadIntegerGrid reads grid of integers from set of lines
func ReadIntegerGrid(lines []string) [][]int {
	var grid [][]int
	for _, line := range lines {
		var gridRow []int
		for _, r := range line {
			gridRow = append(gridRow, int(r-'0'))
		}
		grid = append(grid, gridRow)
	}
	return grid
}

// ReadIntegersFromLine parses integers on a single line from csv format
func ReadIntegersFromLine(line string, sep string) []int {
	sValues := strings.Split(line, sep)
	var values []int
	for _, sValue := range sValues {
		values = append(values, Number(sValue))
	}
	return values
}

// Number parses the supplied string into number. Will return 0
// if value is not a number
func Number(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

// build platform independent file path
func buildPath(dir string, filename string) string {
	return filepath.Join(".", dir, filename)
}

// open filename supplied.  look in current directly if not found in supplied directory
func openFile(dir string, filename string) *os.File {
	file, err := os.Open(buildPath(dir, filename))
	if err != nil {
		file, err = os.Open(buildPath(".", filename))
		if err != nil {
			log.Fatalln(err)
		}
	}
	return file
}

// close file and log any errors
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println("Error closing file: ", err)
	}
}
