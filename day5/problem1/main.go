package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"github.com/shamsch/advent-of-code-22/helper"
)

var positionWithCrates = []string{"1", "5", "9", "13", "17", "21", "25", "29", "33"} 

type position struct {
	crateNumber int
	crate string
}

func parseMovingInstruction(instruction string) (string, string, string) {
	instructionParts := strings.Split(instruction, " ")
	amount := instructionParts[1]
	from := instructionParts[3]
	to := instructionParts[5]
	return amount, from, to
}

func moveTheCrates(amount int, from string, to string, cratesAndPositions map[string][]string) {
	crates := helper.CloneASlice((cratesAndPositions)[from])
	// remove the crates from the from position
	// take from start of the slice
	(cratesAndPositions)[from] = crates[amount:]
	// add the crates to the to position
	// append to the beginning of the slice
	(cratesAndPositions)[to] = append(helper.ReverseSlice(helper.CloneASlice(crates[:amount])), (cratesAndPositions)[to]...)
}

func getTheTopMostCrate(cratesAndPositions map[string][]string) string {
	topMostCrate := []position{}
	topMostCrateAsString := ""

	// get the top most crate
	for idx, val := range cratesAndPositions {
		crateNumber, _ := strconv.Atoi(idx)
		topMostCrate = append(topMostCrate, position{crateNumber, val[0]})
	}
	
	// sort the the positions of the crates in ascending order
	sort.Slice(topMostCrate, func(i, j int) bool {
		return topMostCrate[i].crateNumber < topMostCrate[j].crateNumber
	})

	// append the top most crate to the string
	for _, val := range topMostCrate {
		topMostCrateAsString += val.crate
	}

	return topMostCrateAsString
}

func main() {
	lines := helper.ReadInputFileLines("day5/input.txt")

	reachedEnd := false // reached end of the crates positions 
	cratesAndPositions := make(map[string][]string) // map of crates and their positions
	for _, line := range lines {
		if reachedEnd {
			amount, from, to := parseMovingInstruction(line)
			amountInt, _ := strconv.Atoi(amount)
			moveTheCrates(amountInt, from, to, cratesAndPositions)
		}
		if line == "" {
			reachedEnd = true  
		}
		if !reachedEnd && line[1]!='1'{
			cratePosition := 1
			for idx, val := range line{
				// if the current position is a crate position
				if helper.SliceContains(positionWithCrates, strconv.Itoa(idx)) {
					if string(val) != " "{
						cratesAndPositions[strconv.Itoa(cratePosition)] = append(cratesAndPositions[strconv.Itoa(cratePosition)], string(val))
					}	
					cratePosition++
				}
			}
		}
		
	}
	
	fmt.Println(getTheTopMostCrate(cratesAndPositions))
}