package main

import (
	"regexp"
	"strconv"
)

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
