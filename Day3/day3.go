package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	lines, err := readLines("example.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %s", err)
	}

	// Part 1
	solutionPart1 := solvePart1(lines)
	fmt.Printf("Solution for part 1 is %d\n", solutionPart1)

	// Part 2
	solutionPart2 := solvePart2(lines)
	fmt.Printf("Solution for part 2 is %d\n", solutionPart2)
}

func solvePart1(lines []string) int {
	engineParts := make(map[int]int)
	var engineIndex int
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		fmt.Printf("Line %d: %s\n", y, line)
		for x := 0; x < len(line); x++ {
			tmpString := ""
			char := line[x]
			fmt.Printf("Character %d: %c\n", x, char)
			if unicode.IsNumber(rune(char)) {
				tmpString += string(char)
			} else if tmpString != "" {
				tmpNumber, _ := strconv.Atoi(tmpString)
				engineParts[engineIndex] = tmpNumber
			}
		}
	}

	return -1
}

func solvePart2(lines []string) int {

	return -1
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
