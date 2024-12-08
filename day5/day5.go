package main

import (
	"bufio"
	"log"
	"os"
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

	// wordsFound := SearchWord(grid, "XMAS")

	// fmt.Printf("Total: %d\n", wordsFound)
}
