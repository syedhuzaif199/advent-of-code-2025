package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	PLUS_OP    = -1
	PRODUCT_OP = -2
)

func main() {
	content, err := os.ReadFile("input.txt")
	dieIfErr(err)
	puzzle1(content)
	puzzle2(content)
}

func puzzle1(content []byte) {

	word := []byte{}
	row := []int{}
	numbers := [][]int{}
	for _, b := range content {
		if b == '\n' {
			continue
		}
		if b == '\r' && len(row) != 0 {
			if len(word) != 0 {
				num := wordToNum(word)
				row = append(row, num)
				word = []byte{}
			}
			numbers = append(numbers, row)
			row = []int{}
			word = []byte{}
			continue
		}
		if b == ' ' {
			if len(word) == 0 {
				continue
			}
			num := wordToNum(word)
			row = append(row, num)
			word = []byte{}
			continue
		}

		word = append(word, b)
	}

	total := 0
	ops := numbers[len(numbers)-1]
	numbers = numbers[:len(numbers)-1]
	for i, op := range ops {
		if op == PLUS_OP {
			acc := 0
			for _, row := range numbers {
				acc += row[i]
			}
			total += acc
		} else if op == PRODUCT_OP {
			acc := 1
			for _, row := range numbers {
				acc *= row[i]
			}
			total += acc
		} else {
			log.Fatalf("invalid operator: %c\n", op)
		}
	}

	fmt.Println("Puzzle1:", total)
}

func puzzle2(content []byte) {
	rows := strings.Split(string(content), "\r\n")
	// discard last empty row
	rows = rows[:len(rows)-1]

	// get the actual last row
	lastRow := rows[len(rows)-1]

	rows = rows[:len(rows)-1]

	ops := []byte{}
	for i := range lastRow {
		c := lastRow[i]
		if c == '+' || c == '*' {
			ops = append(ops, c)
		}
	}

	// the number of problems is equal to the number of operators encountered
	frontier := 0
	total := 0
	for _, op := range ops {
		sol := 0
		if op == '*' {
			sol = 1
		}
		for {
			// assuming all rows are of equal length
			if frontier >= len(rows[0]) {
				break
			}

			done := true

			num := 0
			for _, row := range rows {
				if row[frontier] != ' ' {
					done = false
					num = 10*num + asciiByteToInt(row[frontier])
				}
			}
			frontier++
			if done {
				break
			}

			if op == '*' {
				sol *= num
			} else if op == '+' {
				sol += num
			}

		}
		total += sol
	}
	fmt.Println("Puzzle2:", total)
}

func dieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func wordToNum(word []byte) int {
	wordStr := string(word)
	wordStr = strings.TrimSpace(wordStr)
	if len(wordStr) == 1 && (wordStr[0] == '*' || wordStr[0] == '+') {
		if wordStr[0] == '+' {
			return -1
		} else {
			return -2
		}
	}
	num, err := strconv.Atoi(wordStr)
	dieIfErr(err)
	return num
}

func asciiByteToInt(b byte) int {
	if b < '0' || b > '9' {
		log.Fatalf("Invalid character %c is not a number\n", b)
	}

	return int(b - '0')
}
