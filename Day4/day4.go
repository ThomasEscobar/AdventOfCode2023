package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	total := 0

	for _, line := range lines {
		first, second := identifyPairs(line)
		if checkIfPairContained(first, second) {
			total++
		}
	}

	return total
}

func solvePart2(lines []string) int {
	total := 0

	for _, line := range lines {
		first, second := identifyPairs(line)
		if checkIfPairsOverlap(first, second) {
			total++
		}
	}

	return total
}

func checkIfPairsOverlap(first, second []int) bool {
	if first[1] < second[0] || first[0] > second[1] {
		return false
	}

	return true
}

func checkIfPairContained(first, second []int) bool {
	if (first[0] <= second[0] && first[1] >= second[1]) || (first[0] >= second[0] && first[1] <= second[1]) {
		return true
	}

	return false
}

func identifyPairs(line string) (first, second []int) {
	firstStrings := strings.Split(strings.Split(line, ",")[0], "-")
	secondStrings := strings.Split(strings.Split(line, ",")[1], "-")

	first = make([]int, len(firstStrings))
	second = make([]int, len(secondStrings))

	first[0], _ = strconv.Atoi(firstStrings[0])
	first[1], _ = strconv.Atoi(firstStrings[1])
	second[0], _ = strconv.Atoi(secondStrings[0])
	second[1], _ = strconv.Atoi(secondStrings[1])

	return first, second
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
