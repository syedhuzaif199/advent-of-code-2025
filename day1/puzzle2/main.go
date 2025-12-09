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
				newPos := position
				newPos -= mag
				newPos %= 100
				if newPos < 0 {
					newPos += 100
				}
				fmt.Printf("L%d: circled back %d times\n", mag, mag/100)
				count += mag / 100
				if position != 0 && newPos > position || newPos == 0 {
					fmt.Println("remainder passed 0 as newPos =", newPos, "and pos =", position)
					count++
				}
				position = newPos
			} else if dir == 'R' {
				newPos := position
				newPos += mag
				newPos %= 100
				fmt.Printf("R%d: circled back %d times\n", mag, mag/100)
				count += mag / 100
				if newPos < position {
					fmt.Println("remainder passed 0 as newPos =", newPos, "and pos =", position)
					count++
				}
				position = newPos
			} else {
				log.Fatal("Unexpected character:", wordStr[0])
			}
			word = []byte{}
			continue
		}
		word = append(word, b)
	}

	fmt.Println(count)
	fmt.Println(5 / 2)
}
