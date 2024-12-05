package day1

import "testing"

func TestDistances(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	expected := 11

	actual := Distances(list1, list2)

	if actual != expected {
		t.Errorf("got %d want %d", actual, expected)
	}
}
