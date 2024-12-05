package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")
}

func Distances(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0

	for idx, val1 := range list1 {
		val2 := list2[idx]
		if val1 >= val2 {
			distance += val1 - val2
		} else {
			distance += val2 - val1
		}
	}

	fmt.Println(distance)
	return distance
}

func SplitInput(input []string) ([]int, []int) {
	var list1, list2 []int

	for _, value := range input {
		split := strings.Split(value, "   ")
		val1, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		val2, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	return list1, list2
}
