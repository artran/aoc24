package main

import (
	"testing"
)

func TestStability(t *testing.T) {
	data := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	expected := []bool{true, false, false, false, false, true}

	for idx, line := range data {
		stability := Stability(line)
		if stability != expected[idx] {
			t.Errorf("Incorrect stability. Expected %t but got got %t", expected[idx], stability)
		}
	}
}
