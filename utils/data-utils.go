package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

// Number parses the supplied string into number
func Number(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("bad coordinate received: ", s)
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
