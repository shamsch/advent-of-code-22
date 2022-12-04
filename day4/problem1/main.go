package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shamsch/advent-of-code-22/helper"
)

type Pairs struct {
	Start int
	End int
}

func comparePairs(pair1 Pairs, pair2 Pairs) bool {
	if pair1.Start <= pair2.Start && pair1.End >= pair2.End { // does pair1 contain pair2
		return true
	} else if pair1.Start >= pair2.Start && pair1.End <= pair2.End { // does pair2 contain pair1
		return true
	} else {
		return false
	}
}

func main() {
	lines := helper.ReadInputFileLines("day4/input.txt")
	sum := 0

	for _, line := range lines {
		// separate the line by , 
		pairs := strings.Split(line, ",")
		
		// then separate the pairs by -
		firstPair := strings.Split(pairs[0], "-")
		secondPair := strings.Split(pairs[1], "-")

		// then convert the strings to ints
		firstPairStart, _ := strconv.Atoi(firstPair[0])
		firstPairEnd, _ := strconv.Atoi(firstPair[1])

		secondPairStart, _ := strconv.Atoi(secondPair[0])
		secondPairEnd, _ := strconv.Atoi(secondPair[1])
		
		// then create a pair struct
		firstPairStruct := Pairs{firstPairStart, firstPairEnd}
		secondPairStruct := Pairs{secondPairStart, secondPairEnd}
		

		// then compare the pairs
		if comparePairs(firstPairStruct, secondPairStruct) {
			sum++
		}
	}

	fmt.Println(sum)
}

