package main

import (
	"fmt"
	"strings"
)

func main() {
	scanner := ReadScanner("input.txt")

	var score int

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)

		firstHalf := line[:length/2]
		secondHalf := line[length/2:]

		for i := 0; i < length/2; i++ {
			letter := fmt.Sprintf("%c", firstHalf[i])

			isCommon := strings.Contains(secondHalf, letter)
			if isCommon {
				if letter > "Z" {
					score += int(firstHalf[i]) - int('a') + 1
				} else {
					score += int(firstHalf[i]) - int('A') + 1 + 26
				}
				break
			}
		}
	}

	fmt.Println(score)
}
