package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines, err := readLines("input.txt")
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
	stacksOfCratesLines := make([]string, 0)
	instructions := make([]string, 0)
	stacksOfCrates := make(map[int]Stack)

	for _, line := range lines {
		if line != "" && !strings.Contains(line, "move") {
			stacksOfCratesLines = append(stacksOfCratesLines, line)
		} else {
			instructions = append(instructions, line)
		}
	}

	for i, char := range stacksOfCratesLines[len(stacksOfCratesLines)-1] {
		if char != ' ' {
			var 
		}
	}

	return 0
}

func solvePart2(lines []string) int {
	return 0
}

// To implement: https://www.educative.io/answers/how-to-implement-a-stack-in-golang

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

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		indexOfTopElement := len(*s) - 1
		topElement := (*s)[indexOfTopElement]
		*s = (*s)[:indexOfTopElement]
		return topElement, true
	}
}
