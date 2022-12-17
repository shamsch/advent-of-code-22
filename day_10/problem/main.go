package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shamsch/advent-of-code-22/helper"
)

func main() {
	lines := helper.ReadInputFileLines("day_10/input.txt")
	valueArray := []int{}
	pixelArray := [6][40]string{}
	value := 1 
	strength := 0
	for _, line := range lines {
		if line =="noop"{
			valueArray = append(valueArray, value)
		} else {
			splitline := strings.Split(line, " ")
			valueArray = append(valueArray, value)
			valueArray = append(valueArray, value)
			increment, _ := strconv.Atoi(splitline[1])
			value += increment
		}
	}

	// part 1
	for i:=20; i<=220; i+=40{
		strength += valueArray[i-1]*i
	}
	fmt.Println("part1",strength)

	// part 2, did not understand the question prompt at all 

	secondaryValueSlice := createSecondarySlice(valueArray)
	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			position := row*40 + col + 1
			if math.Abs(float64(secondaryValueSlice[position-1] - col)) <=1 {
				pixelArray[row][col] = "##"
			} else {
				pixelArray[row][col] = ".."
			}

		}
	}
	
	fmt.Println("part2")
	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			fmt.Print(pixelArray[row][col])
		}
		fmt.Println()
	}
	
}

func createSecondarySlice(slice []int) []int {
	secondarySlice := make([]int, 241)
	for i := range secondarySlice {
		secondarySlice[i] = 1
	}
	for i := range slice {
		secondarySlice[i] = slice[i]
	}
	return secondarySlice
}