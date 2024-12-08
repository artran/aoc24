package main

import "testing"

func TestWordSearch(t *testing.T) {
	checkResult := func(t testing.TB, found, expected int) {
		t.Helper()
		if found != expected {
			t.Errorf("Found %d, expected %d", found, expected)
		}
	}

	t.Run("with supplied data", func(t *testing.T) {
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

		// found := SearchWord(data, "XMAS")

		// checkResult(t, found, expected)
	})
}
