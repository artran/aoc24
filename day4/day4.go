package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	xBound, yBound int
	word           = "XMAS"
	lenWord        = len(word)
)

// x and y are used to set the direction in which word needs to be searched.
var (
	x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y = []int{-1, 0, 1, -1, 1, -1, 0, 1}
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	xBound = len(grid)
	yBound = len(grid[0])
	wordsFound := SearchWord(grid)

	fmt.Printf("Total: %d\n", wordsFound)
}

func SearchWord(grid [][]byte) int {
	ans := 0
	for i := 0; i < xBound; i++ {
		for j := 0; j < yBound; j++ {
			// add the number of matches from this coordinate
			ans += searchWord2d(grid, i, j, word)
		}
	}

	return ans
}

// Search for the given word in all 8 directions from the coordinate.
func searchWord2d(grid [][]byte, row, col int, word string) int {
	// return early if the given coordinate does not match with first index char.
	if grid[row][col] != word[0] {
		return 0
	}

	// This loop will search in all the 8 directions one by one.
	found := 0
	for dir := 0; dir < 8; dir++ {
		// Initialize starting point for current direction
		currX, currY := row+x[dir], col+y[dir]
		k := 1

		for k < lenWord {
			// break if out of bounds
			if currX >= xBound || currX < 0 || currY >= yBound || currY < 0 {
				break
			}

			// break if character doesn't match expected
			if grid[currX][currY] != word[k] {
				break
			}

			// If we get here we matched the current character so count it
			k += 1

			// Move in specific direction
			currX += x[dir]
			currY += y[dir]
		}

		// If all characters matched, then value of k must be equal to length of word
		if k == lenWord {
			found++
		}
	}

	return found
}
