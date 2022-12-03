package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
		first, second := identifyCompartments(line)
		commonRune := findCommonRune(first, second)
		priority := getPriority(commonRune)
		total += priority
	}

	return total
}

func solvePart2(lines []string) int {
	total := 0
	trioOfElf := make([]string, 0)

	for i, line := range lines {
		// Need to start index at 1 for the remainder/modulo check
		i++
		trioOfElf = append(trioOfElf, line)
		if i%3 == 0 {
			// Slice is "full", process then empty if we're done with the previous trio
			commonRune := findCommonRuneInSlice(trioOfElf)
			total += getPriority(commonRune)
			trioOfElf = nil
		}
	}

	return total
}

func findCommonRuneInSlice(rucksacks []string) rune {
	// Find common runes between the first two strings, then check for those in the last string
	commonRunes := findCommonRunes(rucksacks[0], rucksacks[1])
	commonRunes = findCommonRunes(commonRunes, rucksacks[2])

	return rune(commonRunes[0])
}

func findCommonRunes(first, second string) string {
	commonRunes := ""
	for _, char := range first {
		if strings.Contains(second, string(char)) {
			commonRunes += string(char)
		}
	}

	return commonRunes
}

func findCommonRune(first, second string) rune {
	for _, char := range first {
		if strings.Contains(second, string(char)) {
			return char
		}
	}

	return '?'
}

func getPriority(char rune) int {
	offset := 0
	if unicode.IsUpper(char) {
		// we want to start at 26
		offset = int('A') - 1 - 26
	} else {
		// we want to start at 1
		offset = int('a') - 1
	}

	return int(char) - offset
}

func identifyCompartments(rucksackContent string) (first, second string) {
	return rucksackContent[:len(rucksackContent)/2], rucksackContent[len(rucksackContent)/2:]
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
