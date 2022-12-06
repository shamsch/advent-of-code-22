package main

import (
	"fmt"
	"github.com/shamsch/advent-of-code-22/helper"
)

func main() {
	lines := helper.ReadInputFileLines("day6/input.txt")
	
	for _, line := range lines {
		allChars := breakLineIntoChars(line)
		part1 := findMarker(allChars, 4)
		part2 := findMarker(allChars, 14)
		fmt.Println(part1)
		fmt.Println(part2)
	}
}

func checkIfAllElementsAreUnique(list []string) bool {
	uniqueElements := map[string]bool{}
	for _, val := range list {
		if _, ok := uniqueElements[val]; ok {
			return false
		}
		uniqueElements[val] = true
	}
	return true
}

func findMarker(chars []string, numOfUniqueChar int) int {
	counter := 0
	listOfChars := []string{}
	for index, char := range chars {
		if len(listOfChars)==numOfUniqueChar{
			result := checkIfAllElementsAreUnique(listOfChars)
			if result {
				return index
			}else {
				listOfChars = listOfChars[1:]
			}
		}
		listOfChars = append(listOfChars, char)
		counter++
	}
	return 0
}

func breakLineIntoChars(line string) []string {
	chars := []string{}
	for _, char := range line {
		chars = append(chars, string(char))
	}
	return chars
}