package common

import (
	"bufio"
	"os"
)

func ReadScanner(fileName string) *bufio.Scanner {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	return scanner
}
