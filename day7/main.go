package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	content, err := os.ReadFile("input.txt")
	dieIfErr(err)
	grid := [][]byte{}
	row := []byte{}
	rowIdx := 0
	colIdx := 0
	splits := 0
	splitterActive := false
	for _, c := range content {
		if c == '\n' {
			continue
		}
		if c == '\r' {
			grid = append(grid, row)
			row = []byte{}
			rowIdx++
			colIdx = 0
			splitterActive = false
			continue
		}
		if c == 'S' {
			row = append(row, '|')
		} else if c == '^' {
			row = append(row, '^')
			if rowIdx > 0 {
				if grid[rowIdx-1][colIdx] == '|' {
					splitterActive = true
					splits++
					if colIdx > 0 {
						row[colIdx-1] = '|'
					}
				}
			}
		} else {
			if c != '.' {
				log.Fatalf("invalid character %c found\n", c)
			}
			if colIdx > 0 && row[colIdx-1] == '^' && splitterActive {
				splitterActive = false
				row = append(row, '|')
			} else if rowIdx > 0 && grid[rowIdx-1][colIdx] == '|' {
				row = append(row, '|')
			} else {
				row = append(row, '.')
			}
		}

		colIdx++
	}

	fmt.Println("Splits:", splits)
}

func puzzle2() {
	const (
		splitter = -1
		empty    = 0
	)
	content, err := os.ReadFile("input.txt")
	dieIfErr(err)
	grid := [][]int{}
	row := []int{}
	rowIdx := 0
	colIdx := 0
	activeSplitterCount := 0
	for _, c := range content {
		if c == '\n' {
			continue
		}
		if c == '\r' {
			grid = append(grid, row)
			row = []int{}
			rowIdx++
			colIdx = 0
			activeSplitterCount = 0
			continue
		}
		if c == 'S' {
			row = append(row, 1)
		} else if c == '^' {
			row = append(row, splitter)
			if rowIdx > 0 {
				if grid[rowIdx-1][colIdx] > 0 {
					activeSplitterCount = grid[rowIdx-1][colIdx]
					if colIdx > 0 {
						row[colIdx-1] += activeSplitterCount
					}
				}
			}
		} else {
			if c != '.' {
				log.Fatalf("invalid character %c found\n", c)
			}
			row = append(row, empty)
			if colIdx > 0 && row[colIdx-1] == splitter && activeSplitterCount > 0 {
				row[colIdx] += activeSplitterCount
				activeSplitterCount = 0
			}
			if rowIdx > 0 && grid[rowIdx-1][colIdx] > 0 {
				row[colIdx] += grid[rowIdx-1][colIdx]
			}
		}

		colIdx++
	}

	timelines := 0
	for _, col := range grid[len(grid)-1] {
		timelines += col
	}

	fmt.Println("Timelines:", timelines)

}

func dieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
