package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	stacksOfCratesLines, instructions := parseInput(lines)
	stacksOfCrates := identifyStacks(stacksOfCratesLines)
	executeInstructions(&stacksOfCrates, instructions)

	fmt.Printf("%v", stacksOfCrates)

	return 0
}

func solvePart2(lines []string) int {
	return 0
}

func parseInput(lines []string) ([]string, []string) {
	stacksOfCratesLines := make([]string, 0)
	instructions := make([]string, 0)

	for _, line := range lines {
		if line != "" && !strings.Contains(line, "move") {
			stacksOfCratesLines = append(stacksOfCratesLines, line)
		} else if line != "" {
			instructions = append(instructions, line)
		}
	}

	return stacksOfCratesLines, instructions
}

func identifyStacks(stacksOfCratesLines []string) map[int]Stack {
	stacksOfCrates := make(map[int]Stack)

	for x, stackIndex := range stacksOfCratesLines[len(stacksOfCratesLines)-1] {
		if stackIndex != ' ' {
			tmpStack := make(Stack, 0)
			// -1 because array starts at 0 and -1 because we don't want the last line
			for y := len(stacksOfCratesLines) - 2; y >= 0; y-- {
				var stackChar = string(stacksOfCratesLines[y][x])
				if stackChar != " " {
					tmpStack = append(tmpStack, stackChar)
				}
			}
			stacksOfCrates[convertRuneToDigit(stackIndex)] = tmpStack
		}
	}

	return stacksOfCrates
}

func executeInstructions(p *map[int]Stack, instructions []string) {
	stackOfCrates := *p

	for _, instruction := range instructions {
		fmt.Printf("Before executing instruction '%v': %v\n", instruction, stackOfCrates)

		moveCount := convertRuneToDigit(rune(instruction[5]))
		fromIndex := convertRuneToDigit(rune(instruction[12]))
		toIndex := convertRuneToDigit(rune(instruction[17]))

		for i := 0; i < moveCount; i++ {
			fromStack := stackOfCrates[fromIndex]
			toStack := stackOfCrates[toIndex]
			movingCrate, _ := fromStack.Pop()
			toStack.Push(movingCrate)

			fmt.Printf("Moved %v from %v to %v\n", movingCrate, fromIndex, toIndex)
		}

		fmt.Printf("After executing: %v\n", stackOfCrates)
	}
}

func convertRuneToDigit(r rune) int {
	digit := int(r - '0')
	if digit >= 0 && digit <= 10 {
		return digit
	}
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
