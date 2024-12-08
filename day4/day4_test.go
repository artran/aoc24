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

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})

	t.Run("with corners", func(t *testing.T) {
		data := [][]byte{
			[]byte("XMAS..SAMX"),
			[]byte("MM......MM"),
			[]byte("A.A....A.A"),
			[]byte("S..S..S..S"),
			[]byte(".........."),
			[]byte(".........."),
			[]byte("S..S..S..S"),
			[]byte("A.A....A.A"),
			[]byte("MM......MM"),
			[]byte("XMAS..SAMX"),
		}
		expected := 12

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})

	t.Run("with corners reversed", func(t *testing.T) {
		data := [][]byte{
			[]byte("SAMX..XMAS"),
			[]byte("AA......AA"),
			[]byte("M.M....M.M"),
			[]byte("X..X..X..X"),
			[]byte(".........."),
			[]byte(".........."),
			[]byte("X..X..X..X"),
			[]byte("M.M....M.M"),
			[]byte("AA......AA"),
			[]byte("SAMX..XMAS"),
		}
		expected := 12

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})

	t.Run("with horizontal", func(t *testing.T) {
		data := [][]byte{
			[]byte("XMAS......"),
			[]byte("..XMAS...."),
			[]byte(".XMAS....."),
			[]byte("....XMAS.."),
			[]byte("...XMAS..."),
			[]byte("......XMAS"),
			[]byte(".....XMAS."),
			[]byte("XMAS......"),
			[]byte(".........."),
			[]byte(".........."),
		}
		expected := 8

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})

	t.Run("with horizontal reversed", func(t *testing.T) {
		data := [][]byte{
			[]byte("SAMX......"),
			[]byte("..SAMX...."),
			[]byte(".SAMX....."),
			[]byte("....SAMX.."),
			[]byte("...SAMX..."),
			[]byte("......SAMX"),
			[]byte(".....SAMX."),
			[]byte("SAMX......"),
			[]byte(".........."),
			[]byte(".........."),
		}
		expected := 8

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})

	t.Run("with vertical", func(t *testing.T) {
		data := [][]byte{
			[]byte("X....X...."),
			[]byte("M.X..M...."),
			[]byte("AXM..A...."),
			[]byte("SMAX.S...."),
			[]byte(".ASM......"),
			[]byte(".S.AX.X..."),
			[]byte("XX.SM.M..."),
			[]byte("MM..A.A..."),
			[]byte("AA..S.A..."),
			[]byte("SS........"),
		}
		expected := 8

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})

	t.Run("with crossing", func(t *testing.T) {
		data := [][]byte{
			[]byte("...X......"),
			[]byte("..XMAS...."),
			[]byte(".XMAS....."),
			[]byte("XMAS......"),
			[]byte(".AS.....S."),
			[]byte(".S.....SA."),
			[]byte("......SAMX"),
			[]byte(".....SAMX."),
			[]byte("....SAMX.."),
			[]byte("......X..."),
		}
		expected := 12

		found := SearchWord(data, "XMAS")

		checkResult(t, found, expected)
	})
}
