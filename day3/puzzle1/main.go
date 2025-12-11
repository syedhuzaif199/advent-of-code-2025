package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")
	dieIfErr(err)
	banks := []string{}
	bankBuilder := []byte{}
	for _, b := range content {
		if b == '\n' {
			if len(bankBuilder) != 0 {
				log.Fatal("expected bank to be empty, but found", string(bankBuilder))
			}
			continue
		}

		if b == '\r' {
			banks = append(banks, string(bankBuilder))
			bankBuilder = []byte{}
			continue
		}
		bankBuilder = append(bankBuilder, b)
	}

	totalJoltage := 0
	for _, bank := range banks {
		first := 0
		for i := 0; i < len(bank)-1; i++ {
			if bank[i] > bank[first] {
				first = i
			}
		}
		second := first + 1
		for i := first + 1; i < len(bank); i++ {
			if bank[i] > bank[second] {
				second = i
			}
		}

		joltage, err := strconv.Atoi(string([]byte{bank[first], bank[second]}))
		dieIfErr(err)
		totalJoltage += joltage
	}

	fmt.Println(totalJoltage)
}

func dieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
