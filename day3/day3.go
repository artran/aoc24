package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	total := DeCorrupt(strings.Join(lines, ""))

	fmt.Printf("Total: %d", total)
}

func DeCorrupt(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
	found := re.FindAllStringSubmatch(data, -1)

	enabled := true
	total := 0
	for _, match := range found {
		if match[0] == "don't()" {
			enabled = false
			continue
		}
		if match[0] == "do()" {
			enabled = true
			continue
		}

		// Add to the total only if calculation is enabled
		if !enabled {
			continue
		}
		lhs, _ := strconv.Atoi(match[1])
		rhs, _ := strconv.Atoi(match[2])
		total += lhs * rhs
	}
	return total
}
