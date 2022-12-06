package main

import (
	"fmt"
)

func main() {
	scanner := ReadScanner("../input.txt")

	positions := GetCraneInitialPos(scanner)

	movements := GetInputMovements(scanner)

	finalPositions := MoveStuff(positions, movements)

	upperBoxes := GetUpperBoxes(finalPositions)

	fmt.Println(upperBoxes)
}
