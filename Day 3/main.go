package main

import (
	"bufio"
	"fmt"
)

var (
	linesPerGroup = 3
)

func main() {
	scanner := ReadScanner("input.txt")

	var score int

	for scanner.Scan() {
		lines := ReadLines(scanner, linesPerGroup)
		var appearances [53]int

		for i, line := range lines {
			for j := 0; j < len(line); j++ {
				letter := int(line[j])
				letterScore := CalcLetterScore(letter)

				if appearances[letterScore] == i {
					appearances[letterScore]++
				}

				if appearances[letterScore] == linesPerGroup {
					score += letterScore
					break
				}
			}
		}

	}
	fmt.Println(score)
}

func ReadLines(scanner *bufio.Scanner, noOfLines int) []string {
	lines := make([]string, 0)

	//Scan() already called once before this function is called
	line := scanner.Text()
	lines = append(lines, line)

	for i := 0; i < noOfLines-1; i++ {
		scanner.Scan()
		line = scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func CalcLetterScore(letterAscii int) int {
	if letterAscii > int('Z') {
		return letterAscii - int('a') + 1
	} else {
		return letterAscii - int('A') + 1 + 26
	}
}
