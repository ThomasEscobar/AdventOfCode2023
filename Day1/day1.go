package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	lines, err := readLines("input.txt")
	// Aiming for 142
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
	// take 1st number and last number of each line
	// put them together (not added !)
	// each value add them together
	sum := 0
	for _, line := range lines {
		//fmt.Printf("The current line is: %s\n", line)

		// Finding the first number in the line by going through each character "forwards"
		firstNumber := 0
		for _, character := range line {
			if unicode.IsNumber(character) {
				//fmt.Printf("The character %c is a number\n", character)
				firstNumber = int(character) - '0'
				break
			}
		}
		//fmt.Printf("The first number is %d\n", firstNumber)

		// Finding the last number in the line by going through each character "backwards"
		lastNumber := 0
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				lastNumber = int(line[i]) - '0'
				break
			}
		}
		//fmt.Printf("The last number is %d\n", lastNumber)

		numbersTogether := fmt.Sprintf("%d%d", firstNumber, lastNumber)
		//fmt.Printf("The numbers together are %s\n", numbersTogether)

		numbersTogetherAsANumber, _ := strconv.Atoi(numbersTogether)
		sum += numbersTogetherAsANumber
	}

	return sum
}

func solvePart2(lines []string) int {
	// Modify each line to replace the word/number with the number/number ! Then run first solution
	var modifiedLines []string
	for _, line := range lines {
		fmt.Printf("The original line is %s\n", line)

		// Start by replacing the overlapping word/numbers !
		// twone -> 21
		// zerone -> 01
		// eightwo -> 82
		// eightree -> 83
		// oneight -> 18
		// threeight -> 38
		// fiveight -> 58
		// nineight -> 98
		// sevenine -> 79
		re21 := regexp.MustCompile("twone")
		re01 := regexp.MustCompile("zerone")
		re82 := regexp.MustCompile("eightwo")
		re83 := regexp.MustCompile("eightree")
		re18 := regexp.MustCompile("oneight")
		re38 := regexp.MustCompile("threeight")
		re58 := regexp.MustCompile("fiveight")
		re98 := regexp.MustCompile("nineight")
		re79 := regexp.MustCompile("sevenine")

		modifiedLine := re21.ReplaceAllString(line, "21")
		modifiedLine = re01.ReplaceAllString(modifiedLine, "01")
		modifiedLine = re82.ReplaceAllString(modifiedLine, "82")
		modifiedLine = re83.ReplaceAllString(modifiedLine, "83")
		modifiedLine = re18.ReplaceAllString(modifiedLine, "18")
		modifiedLine = re38.ReplaceAllString(modifiedLine, "38")
		modifiedLine = re58.ReplaceAllString(modifiedLine, "58")
		modifiedLine = re98.ReplaceAllString(modifiedLine, "98")
		modifiedLine = re79.ReplaceAllString(modifiedLine, "79")

		// Then, replace the "normal" word/numbers
		re0 := regexp.MustCompile("zero")
		re1 := regexp.MustCompile("one")
		re2 := regexp.MustCompile("two")
		re3 := regexp.MustCompile("three")
		re4 := regexp.MustCompile("four")
		re5 := regexp.MustCompile("five")
		re6 := regexp.MustCompile("six")
		re7 := regexp.MustCompile("seven")
		re8 := regexp.MustCompile("eight")
		re9 := regexp.MustCompile("nine")

		modifiedLine = re0.ReplaceAllString(modifiedLine, "0")
		modifiedLine = re1.ReplaceAllString(modifiedLine, "1")
		modifiedLine = re2.ReplaceAllString(modifiedLine, "2")
		modifiedLine = re3.ReplaceAllString(modifiedLine, "3")
		modifiedLine = re4.ReplaceAllString(modifiedLine, "4")
		modifiedLine = re5.ReplaceAllString(modifiedLine, "5")
		modifiedLine = re6.ReplaceAllString(modifiedLine, "6")
		modifiedLine = re7.ReplaceAllString(modifiedLine, "7")
		modifiedLine = re8.ReplaceAllString(modifiedLine, "8")
		modifiedLine = re9.ReplaceAllString(modifiedLine, "9")

		fmt.Printf("The modified line is %s\n", modifiedLine)

		modifiedLines = append(modifiedLines, modifiedLine)
	}

	return solvePart1(modifiedLines)
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
