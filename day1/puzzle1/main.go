package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	word := []byte{}
	position := 50
	count := 0
	for _, b := range content {
		if b == '\n' {
			if len(word) != 0 {
				log.Fatal("expected to start a new word, but word is:", string(word))
			}
			continue
		}
		if b == '\r' {
			wordStr := string(word)
			dir := wordStr[0]
			mag, err := strconv.Atoi(wordStr[1:])
			if err != nil {
				log.Fatal(err)
			}
			if dir == 'L' {
				position -= mag
				position %= 100
				if position < 0 {
					position += 100
				}
			} else if dir == 'R' {
				position += mag
				position %= 100
			} else {
				log.Fatal("Unexpected character:", wordStr[0])
			}
			if position == 0 {
				count++
			}
			word = []byte{}
			continue
		}
		word = append(word, b)
	}

	fmt.Println(count)

}
