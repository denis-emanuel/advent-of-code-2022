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
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
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
		myChoice := rpsInputs[input[1]]

		roundScore += shapeScore[myChoice]
		roundScore += calcResultPoints(myChoice, opponentChoice)

		finalScore += roundScore
		// fmt.Fprintf(os.Stdout, "score is: %d", roundScore)
	}

	fmt.Fprintf(os.Stdout, "Final score: %d", finalScore)
}

func calcResultPoints(myChoice string, opponentsChoice string) int {
	if shapeScore[myChoice] == shapeScore[opponentsChoice] {
		return 3
	}

	if shapeScore[myChoice] == shapeScore[opponentsChoice]+1 ||
		shapeScore[myChoice] == shapeScore[opponentsChoice]-2 {
		return 6
	}

	return 0
}
