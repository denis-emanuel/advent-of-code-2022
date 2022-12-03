package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shapeScore = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var rpsInputs = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var resultScore = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

func main() {
	// file, _ := os.Open("input-test.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var finalScore int
	for scanner.Scan() {
		var roundScore int

		input := strings.Split(scanner.Text(), " ")
		opponentChoice := rpsInputs[input[0]]
		result := input[1]

		roundScore += resultScore[result]
		roundScore += getMyChoicePoints(opponentChoice, result)

		finalScore += roundScore
		// fmt.Fprintf(os.Stdout, "score is: %d", roundScore)
	}

	fmt.Fprintf(os.Stdout, "Final score: %d", finalScore)
}

func getMyChoicePoints(opponentChoice string, result string) int {
	if resultScore[result] == 3 {
		return shapeScore[opponentChoice]
	}

	if resultScore[result] == 0 {
		if shapeScore[opponentChoice] > 1 {
			return shapeScore[opponentChoice] - 1
		} else {
			return 3
		}
	}

	if shapeScore[opponentChoice] < 3 {
		return shapeScore[opponentChoice] + 1
	} else {
		return 1
	}
}
