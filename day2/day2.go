package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	stabiltyCount := 0
	for _, line := range lines {
		if Stability(line) {
			stabiltyCount += 1
		}
	}

	fmt.Printf("Total: %d", stabiltyCount)
}

func Stability(line string) bool {
	reports := stringToReports(line)

	// First try the unchanged reports
	if stabilityOfReports(reports) {
		return true
	}

	// If that failed try removing each element in turn
	for idx := 0; idx < len(reports); idx++ {
		copied := slices.Clone(reports)
		copied = slices.Delete(copied, idx, idx+1)
		s := stabilityOfReports(copied)
		if s {
			return true
		}
	}

	return false
}

func stringToReports(line string) []int {
	split := strings.Split(line, " ")
	reports := make([]int, len(split))
	var err error

	for i, s := range split {
		reports[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}

	return reports
}

func stabilityOfReports(reports []int) bool {
	prev := reports[0]

	direction := 0 // 1 -> increasing, -1 -> decreasing, 0 -> not set yet
	for _, next := range reports[1:] {
		diff := 0
		if next > prev {
			if direction == -1 {
				return false
			}
			direction = 1
			diff = next - prev
		} else {
			if direction == 1 {
				return false
			}
			direction = -1
			diff = prev - next
		}

		if diff > 3 || diff < 1 {
			return false
		}

		prev = next
	}

	return true
}
