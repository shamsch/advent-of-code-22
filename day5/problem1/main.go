package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shamsch/advent-of-code-22/helper"
)
var positionWithCrates = []string{"1", "5", "9", "13", "17", "21", "25", "29", "33"} 

func parseMovingInstruction(instruction string) (string, string, string) {
	instructionParts := strings.Split(instruction, " ")
	amount := instructionParts[1]
	from := instructionParts[3]
	to := instructionParts[5]
	return amount, from, to
}

func main() {
	lines := helper.ReadInputFileLines("day5/input.txt")

	reachedEnd := false // reached end of the crates positions 
	cratesAndPositions := make(map[string][]string) // map of crates and their positions
	for _, line := range lines {
		if reachedEnd {
			amount, from, to := parseMovingInstruction(line)
			amountInt, _ := strconv.Atoi(amount)
			fmt.Println(amountInt, from, to)
		}
		if line == "" {
			reachedEnd = true  
		}
		if !reachedEnd && line[1]!='1'{
			cratePosition := 1
			for idx, val := range line{
				// if the current position is a crate position
				if helper.SliceContains(positionWithCrates, strconv.Itoa(idx)) {
					cratesAndPositions[strconv.Itoa(cratePosition)] = append(cratesAndPositions[strconv.Itoa(cratePosition)], string(val))
					cratePosition++
				}
			}
		}
		
	}
	for key, val := range cratesAndPositions {
		fmt.Println(key, val)
	}
}