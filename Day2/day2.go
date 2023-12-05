package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	var sum int
	const limitRed, limitGreen, limitBlue = 12, 13, 14

	// Game ID is idx+1, no need to read again
	for idx, line := range lines {
		// Get red max
		maxRed := findColorMaximum(line, "red")
		// Get green max
		maxGreen := findColorMaximum(line, "green")
		// Get blue max
		maxBlue := findColorMaximum(line, "blue")

		// If the maximum for each color is less than the respective limit, game is valid, add ID to sum
		if maxRed <= limitRed && maxGreen <= limitGreen && maxBlue <= limitBlue {
			sum += idx + 1
		}
	}

	return sum
}

func solvePart2(lines []string) int {
	var sum int
	for _, line := range lines {
		// Get red max
		maxRed := findColorMaximum(line, "red")
		// Get green max
		maxGreen := findColorMaximum(line, "green")
		// Get blue max
		maxBlue := findColorMaximum(line, "blue")
		power := maxRed * maxGreen * maxBlue
		sum += power
	}

	return sum
}

func findColorMaximum(line string, color string) int {
	var max int
	reColorCount := regexp.MustCompile(fmt.Sprintf("(\\d+) %s", color))
	matches := reColorCount.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		// For each match, item 0 contains the full matched pattern, item 1 contains the captured group
		count, _ := strconv.Atoi(match[1])
		if count > max {
			max = count
		}
	}

	return max
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
