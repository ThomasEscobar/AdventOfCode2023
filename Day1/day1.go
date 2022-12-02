package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	elfCals := make(map[int]int)
	elfIndex := 0

	for _, line := range lines {
		if line != "" {
			cal, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Failed to convert to int: %s", err)
			}
			elfCals[elfIndex] += cal
		} else {
			elfIndex++
		}
	}

	richestElfIndex, maxCal := findRichestElf(elfCals)

	fmt.Printf("The richest elf is elf %d and he has %d calories\n", richestElfIndex, maxCal)

	return maxCal
}

func solvePart2(lines []string) int {
	// Sorting seems complicated with maps (or even impossible ? Sorted by index)... Let's re-use the elf index and get the max 3 times !

	elfCals := make(map[int]int)
	elfIndex := 0

	for _, line := range lines {
		if line != "" {
			cal, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Failed to convert to int: %s", err)
			}
			elfCals[elfIndex] += cal
		} else {
			elfIndex++
		}
	}

	// 1st richest elf
	richestElfIndex, maxCal := findRichestElf(elfCals)
	fmt.Printf("The richest elf is elf %d and he has %d calories\n", richestElfIndex, maxCal)
	delete(elfCals, richestElfIndex)

	top3Cals := maxCal

	// 2nd richest elf
	richestElfIndex, maxCal = findRichestElf(elfCals)
	fmt.Printf("The 2nd richest elf is elf %d and he has %d calories\n", richestElfIndex, maxCal)
	delete(elfCals, richestElfIndex)

	top3Cals += maxCal

	// 3rd richest elf
	richestElfIndex, maxCal = findRichestElf(elfCals)
	fmt.Printf("The 3rd richest elf is elf %d and he has %d calories\n", richestElfIndex, maxCal)
	delete(elfCals, richestElfIndex)

	top3Cals += maxCal

	return top3Cals
}

func findRichestElf(elfCals map[int]int) (richestElfIndex, maxCal int) {
	maxCal = 0

	for i := range elfCals {
		if elfCals[i] > maxCal {
			richestElfIndex = i
			maxCal = elfCals[i]
		}
	}

	return richestElfIndex, maxCal
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
