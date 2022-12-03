package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("elf-numbers.txt")
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var sum, max1, max2, max3 int
	for scanner.Scan() {
		numStr := scanner.Text()
		num, _ := strconv.Atoi(numStr)

		if numStr != "" {
			sum += num
			continue
		}

		if sum >= max3 && sum <= max2 {
			max3 = sum
		}
		if sum >= max2 && sum <= max1 {
			max3 = max2
			max2 = sum
		}
		if sum > max1 {
			max3 = max2
			max2 = max1
			max1 = sum
		}
		sum = 0
	}

	fmt.Printf("Day1: Max is: %d", max1)
	fmt.Printf("Party II: Sum of first 3 max is: %d", max1+max2+max3)

}
