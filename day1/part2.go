package main

func Similarity(list1, list2 []int) int {
	gathered := gatherTerms(list2)
	total := 0

	for _, val := range list1 {
		total += val * gathered[val]
	}

	return total
}

func gatherTerms(lst []int) map[int]int {
	ret := make(map[int]int, 0)

	for _, key := range lst {
		ret[key] += 1
	}

	return ret
}
