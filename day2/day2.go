package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	split := strings.Split(line, " ")

	prev, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	direction := 0 // 1 -> increasing, -1 -> decreasing, 0 -> not set yet
	for _, nextStr := range split[1:] {
		next, err := strconv.Atoi(nextStr)
		if err != nil {
			panic(err)
		}

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
