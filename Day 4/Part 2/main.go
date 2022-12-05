package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	scanner := ReadScanner("../input.txt")

	var pairs int

	for scanner.Scan() {
		input := scanner.Text()

		intervals := GetInputIntervals(input)

		if CalcIntervalOverlaps(intervals) {
			pairs++
			fmt.Println(intervals)
		}

	}

	fmt.Println(pairs)
}

func GetInputIntervals(input string) [4]int {
	interval1 := strings.Split(input, ",")[0]
	interval2 := strings.Split(input, ",")[1]

	start1, _ := strconv.Atoi(strings.Split(interval1, "-")[0])
	end1, _ := strconv.Atoi(strings.Split(interval1, "-")[1])
	start2, _ := strconv.Atoi(strings.Split(interval2, "-")[0])
	end2, _ := strconv.Atoi(strings.Split(interval2, "-")[1])

	return [4]int{start1, end1, start2, end2}
}

func CalcIntervalOverlaps(intervals [4]int) bool {
	if intervals[0] >= intervals[2] && intervals[0] <= intervals[3] {
		return true
	}

	if intervals[1] >= intervals[2] && intervals[1] <= intervals[3] {
		return true
	}

	if intervals[0] >= intervals[2] && intervals[1] <= intervals[3] {
		return true
	}

	if intervals[0] <= intervals[2] && intervals[1] >= intervals[3] {
		return true
	}

	return false
}
