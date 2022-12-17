package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shamsch/advent-of-code-22/helper"
)

func main() {
	lines := helper.ReadInputFileLines("day_10/input.txt")
	valueArray := []int{}
	value := 1 
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
	fmt.Println(valueArray)
}

func calculateValue()  {
	
}
