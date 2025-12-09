package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content = content[:len(content)-2]

	idRanges := strings.Split(string(content), ",")
	invalidIdSum := 0
	for _, idRange := range idRanges {
		lowerUpper := strings.Split(idRange, "-")
		lower, upper := lowerUpper[0], lowerUpper[1]
		if len(lower) == len(upper) && len(lower)&1 != 0 {
			// no invalid id here
			continue
		}

		lowerNum, err := strconv.Atoi(lower)
		dieIfErr(err)
		upperNum, err := strconv.Atoi(upper)
		dieIfErr(err)

		if len(lower)&1 != 0 {
			lowerNum, err := strconv.Atoi(lower)
			if err != nil {
				log.Fatal(err)
			}
			// get the smallest number with even number of digits that is greater than lowerNum
			// which is the same as saying
			// get the smallest power of 10 larger than lowerNum, make that the new value of lowerNum
			lowerNum = int(math.Pow10(len(lower)))
			lower = strconv.Itoa(lowerNum)
		}
		if len(upper)&1 != 0 {
			upperNum, err := strconv.Atoi(upper)
			dieIfErr(err)
			// get the largest number with even number of digits that is smaller than upperNum
			// which is the same as saying
			// get the largest power of 10 smaller than upperNum, and subtract 1 (e.g., make 151 into 100, then subtract 1 to get 99)
			upperNum = int(math.Pow10(len(upper)-1)) - 1
			upper = strconv.Itoa(upperNum)
		}

		firstHalfLower := lower[:len(lower)/2]
		firstHalfUpper := upper[:len(upper)/2]

		firstHalfLowerNum, err := strconv.Atoi(firstHalfLower)
		dieIfErr(err)
		firstHalfUpperNum, err := strconv.Atoi(firstHalfUpper)
		dieIfErr(err)

		secondHalfUpper := upper[len(upper)/2:]
		secondHalfUpperNum, err := strconv.Atoi(secondHalfUpper)
		dieIfErr(err)
		secondHalfLower := lower[len(lower)/2:]
		secondHalfLowerNum, err := strconv.Atoi(secondHalfLower)
		dieIfErr(err)
		if firstHalfLowerNum != firstHalfUpperNum && firstHalfLowerNum >= secondHalfLowerNum {
			invalidIdSum += getInvalidId(firstHalfLowerNum, lowerNum, upperNum, lower, upper)
		}
		for i := firstHalfLowerNum + 1; i < firstHalfUpperNum; i++ {
			invalidIdSum += getInvalidId(i, lowerNum, upperNum, lower, upper)
		}
		if firstHalfUpperNum <= secondHalfUpperNum {
			invalidIdSum += getInvalidId(firstHalfUpperNum, lowerNum, upperNum, lower, upper)
		}
	}

	fmt.Println(invalidIdSum)
}

func getInvalidId(num, lowerNum, upperNum int, lower, upper string) int {
	invalidId := num * int(math.Pow10(len(strconv.Itoa(num)))+1.0)
	invalidStr := strconv.Itoa(invalidId)
	if strings.Compare(invalidStr[:len(invalidStr)/2], invalidStr[len(invalidStr)/2:]) != 0 {
		fmt.Println("Mistake:", invalidStr)
	}

	if invalidId < lowerNum || invalidId > upperNum {
		fmt.Println(invalidStr, "is out of range ["+lower+", "+upper+"]")
	}
	return invalidId
}

func dieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
