package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	total := 0
	for _, line := range lines {
		total += DeCorrupt(line)
	}

	fmt.Printf("Total: %d", total)
}

func DeCorrupt(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	found := re.FindAllStringSubmatch(data, -1)

	total := 0
	for _, match := range found {
		lhs, _ := strconv.Atoi(match[1])
		rhs, _ := strconv.Atoi(match[2])
		total += lhs * rhs
	}
	return total
}
