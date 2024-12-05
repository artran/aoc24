package day1

import (
	"fmt"
	"slices"
)

func Distances(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0

	for idx, val1 := range list1 {
		distance += list2[idx] - val1
	}

	fmt.Println(distance)
	return distance
}
