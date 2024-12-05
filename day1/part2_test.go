package main

import "testing"

func TestSimilarity(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	expected := 31

	similarity := Similarity(list1, list2)

	if similarity != expected {
		t.Errorf("got %d want %d", similarity, expected)
	}
}
