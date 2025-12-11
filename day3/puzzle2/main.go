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

	var totalJoltage uint64 = 0
	for _, bank := range banks {
		previous := 0
		var joltage uint64 = 0
		// assuming all banks are at least 12 batteries long
		for k := range 12 {
			for i := previous; i < len(bank)-(12-k-1); i++ {
				if bank[i] > bank[previous] {
					previous = i
				}
			}
			num, err := strconv.Atoi(string(bank[previous]))
			dieIfErr(err)
			joltage = joltage*10 + uint64(num)
			previous++
		}
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
