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
	score := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		oppPlay := translateEncryptedMove(fields[0])
		myPlay := translateEncryptedMove(fields[1])
		score += resolveRockPaperScissorsRound(oppPlay, myPlay)
	}

	return score
}

func solvePart2(lines []string) int {
	score := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		oppPlay := translateEncryptedMove(fields[0])
		targetOutcome := fields[1]
		myPlay := choosePlayToGetTargetOutcome(oppPlay, targetOutcome)
		score += resolveRockPaperScissorsRound(oppPlay, myPlay)
	}

	return score
}

func choosePlayToGetTargetOutcome(oppPlay int, targetOutcome string) int {
	myPlay := 0

	switch targetOutcome {
	case "X":
		// Loose
		myPlay = oppPlay - 1
		if myPlay == 0 {
			myPlay = 3
		}
	case "Y":
		// Draw
		myPlay = oppPlay
	case "Z":
		// Win
		myPlay = oppPlay + 1
		if myPlay == 4 {
			myPlay = 1
		}
	default:
		log.Fatalf("The outcome %s is unknown and couldn't be processed", targetOutcome)
	}

	return myPlay
}

func resolveRockPaperScissorsRound(oppPlay, myPlay int) int {
	roundScore := myPlay

	if oppPlay == myPlay {
		// Draw
		roundScore += 3
	} else if myPlay == oppPlay+1 || myPlay == 1 && oppPlay == 3 {
		// Win
		roundScore += 6
	} else {
		// Loose
	}

	return roundScore
}

func translateEncryptedMove(letter string) int {
	// Converting to score value instead to save a step later
	switch letter {
	case "A", "X":
		//return "rock"
		return 1
	case "B", "Y":
		//return "paper"
		return 2
	case "C", "Z":
		//return "scissors"
		return 3
	default:
		log.Fatalf("The move %s is unknown and couldn't be processed", letter)
		return -1
	}
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
