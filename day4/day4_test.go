package main

import "testing"

func TestWordSearch(t *testing.T) {
	data := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	expected := 18

	found := SearchText(data)

	if found != expected {
		t.Errorf("Found %d, expected %d", found, expected)
	}
}
