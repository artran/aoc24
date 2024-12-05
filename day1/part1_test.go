package main

import (
	"slices"
	"testing"
)

func TestDistances(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	expected := 11

	actual := Distances(list1, list2)

	if actual != expected {
		t.Errorf("got %d want %d", actual, expected)
	}
}

func TestLineSplitting(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	expectedList1 := []int{3, 4, 2, 1, 3, 3}
	expectedList2 := []int{4, 3, 5, 3, 9, 3}

	list1, list2 := SplitInput(input)

	if slices.Equal(list1, expectedList1) == false {
		t.Errorf("First list is wrong: got %d, want %d", list1, expectedList1)
	}

	if slices.Equal(list2, expectedList2) == false {
		t.Errorf("Second list is wrong: got %d, want %d", list2, expectedList2)
	}
}
