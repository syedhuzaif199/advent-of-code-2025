package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	dieIfErr(err)
	grid := [][]byte{}
	current := []byte{}
	for _, b := range content {
		if b == '\n' {
			continue
		}
		if b == '\r' {
			grid = append(grid, current)
			current = make([]byte, 0, len(current))
			continue
		}
		current = append(current, b)
	}

	total := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			count := 0
			for i := -1; i <= 1; i++ {
				if row+i < 0 || row+i >= len(grid) {
					continue
				}
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if col+j < 0 || col+j >= len(grid[0]) {
						continue
					}
					if grid[row+i][col+j] == '@' {
						count++
					}
				}
			}
			if count < 4 && grid[row][col] == '@' {
				total++
			}
		}
	}
	fmt.Println(total)
}

func dieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
