package main

import "testing"

func TestWordSearch(t *testing.T) {
	data := [][]byte{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	}
	expected := 18

	found := SearchWord(data)

	if found != expected {
		t.Errorf("Found %d, expected %d", found, expected)
	}
}
