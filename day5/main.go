package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	dieIfErr(err)
	splitted := strings.Split(string(content), "\r\n\r\n")
	idRanges := strings.Split(splitted[0], "\r\n")
	availableIds := strings.Split(splitted[1], "\r\n")
	availableIds = availableIds[:len(availableIds)-1] // discard the extra CRLF at the end
	ranges := [][]int64{}
	for _, idRange := range idRanges {
		lowerUpper := strings.Split(idRange, "-")
		lower, err := strconv.ParseInt(lowerUpper[0], 0, 64)
		dieIfErr(err)
		upper, err := strconv.ParseInt(lowerUpper[1], 0, 64)
		dieIfErr(err)
		newRange := []int64{lower, upper}
		ranges = append(ranges, newRange)
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var curLow int64 = 0
	// initialize curUp to -1 to make the range [curLow, curUp] invalid ([0, -1])
	// also, when incrementing count, initially, curUp - curLow + 1 will evaluate to 0 (-1 - 0 + 1 = 0)
	var curUp int64 = -1
	var count int64 = 0
	if len(ranges) > 0 {
		curLow = ranges[0][0]
		curUp = ranges[0][1]
	}
	coalescedRanges := [][]int64{}
	for i, r := range ranges {
		if i == 0 {
			continue
		}
		lower := r[0]
		upper := r[1]
		if lower <= curUp {
			curUp = max(curUp, upper)
		} else {
			count += curUp - curLow + 1
			n := []int64{curLow, curUp}
			coalescedRanges = append(coalescedRanges, n)
			curLow = lower
			curUp = upper
		}
	}
	count += curUp - curLow + 1

	coalescedRanges = append(coalescedRanges, []int64{curLow, curUp})

	sort.Slice(availableIds, func(i, j int) bool {
		val1, err := strconv.ParseInt(availableIds[i], 0, 64)
		dieIfErr(err)
		val2, err := strconv.ParseInt(availableIds[j], 0, 64)
		dieIfErr(err)
		return val1 < val2
	})

	idIdx, rangeIdx := 0, 0
	availableFreshCount := 0
	for idIdx < len(availableIds) && rangeIdx < len(coalescedRanges) {
		availableId, err := strconv.ParseInt(availableIds[idIdx], 0, 64)
		dieIfErr(err)
		lower := coalescedRanges[rangeIdx][0]
		upper := coalescedRanges[rangeIdx][1]

		if availableId >= lower && availableId <= upper {
			availableFreshCount++
			idIdx++
		} else if availableId > upper {
			rangeIdx++
		} else {
			idIdx++
		}
	}
	fmt.Println("Available fresh ingredients:", availableFreshCount)
	fmt.Println("Total fresh ingredients:", count)
}

func dieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
