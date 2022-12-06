package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func GetCraneInitialPos(scanner *bufio.Scanner) [][]string {
	inputPos := make([][]string, 0)
	i := 0

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			break
		}

		inputPos = append(inputPos, []string{})

		for j := 1; j < len(input); j += 4 {
			letter := fmt.Sprintf("%c", input[j])

			inputPos[i] = append(inputPos[i], letter)
		}
		i++
	}

	fmt.Println(inputPos)

	rows := len(inputPos[0])
	positions := make([][]string, rows)

	for i := len(inputPos) - 1 - 1; i >= 0; i-- {
		for j := 0; j < len(inputPos[0]); j++ {
			if inputPos[i][j] == " " {
				continue
			}

			// line number == stack number
			positions[j] = append(positions[j], inputPos[i][j])
		}
	}

	fmt.Println(positions)

	return positions
}

type Movements struct {
	NoToMove int
	From     int
	To       int
}

func GetInputMovements(scanner *bufio.Scanner) []Movements {
	movements := make([]Movements, 0)

	for scanner.Scan() {
		input := scanner.Text()
		noToMoveStr := strings.Split(input, "move ")[1]
		fromStr := strings.Split(input, "from ")[1]
		toStr := strings.Split(input, "to ")[1]

		noToMove := GetNumberInputFromSubstr(noToMoveStr)
		from := GetNumberInputFromSubstr(fromStr)
		to := GetNumberInputFromSubstr(toStr)

		movements = append(movements, Movements{NoToMove: noToMove, From: from, To: to})
	}

	return movements
}

func GetNumberInputFromSubstr(substr string) int {
	var numStr string

	for i := 0; i < len(substr); i++ {
		char := fmt.Sprintf("%c", substr[i])

		if char >= "0" && char <= "9" {
			numStr += char

			continue
		}

		break
	}

	num, _ := strconv.Atoi(numStr)

	return num
}

func MoveStuff(positions [][]string, movements []Movements) [][]string {
	for idx, movement := range movements {
		fmt.Println("INDEX:", idx)
		for i := movement.NoToMove; i > 0; i-- {
			fromLength := len(positions[movement.From-1])
			fmt.Println("---------------------------------------")
			fmt.Println(fromLength)
			fmt.Println(movement.From)
			fmt.Println(movement.To)
			fmt.Println(positions[movement.From-1])
			fmt.Println(positions)
			fmt.Println(positions[movement.From-1][fromLength-1:][0])

			upperBox := positions[movement.From-1][fromLength-1:][0]
			positions[movement.To-1] = append(positions[movement.To-1], upperBox)

			positions[movement.From-1] = positions[movement.From-1][:fromLength-1]
		}
	}

	return positions
}

func GetUpperBoxes(positions [][]string) string {
	upperBoxes := ""

	for _, position := range positions {
		lastOccupied := position[0]

		for i := 1; i < len(position); i++ {
			if position[i] == "" {
				break
			}
			lastOccupied = position[i]
		}

		upperBoxes += lastOccupied
	}

	return upperBoxes
}
