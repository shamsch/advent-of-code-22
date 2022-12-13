package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shamsch/advent-of-code-22/helper"
)

func main() {
	lines := helper.ReadInputFileLines("day_10/input.txt")
	strength := 0
	value := 0 
	valueAndLine := make(map[int]int)
	for idx, line := range lines {
		lineNumber := idx+1
		if line != "noop"{
			// separate the line into multiple parts
			parts := strings.Split(line, " ")
			lineStrength, _ := strconv.Atoi(parts[1])
			valueAndLine[lineNumber] = lineStrength
		}
	}
	value = calculateValue(20, valueAndLine)
	fmt.Println(strength, value)
}

func calculateValue(cycle int, valueAndLine map[int]int) int {
	value := 0
	tempSlice := []int{}
	sortedKeys := helper.SortMapByKeys(valueAndLine)
	 half := cycle/2
	for _, line := range sortedKeys {
		fmt.Println(line, valueAndLine[line])
		if len(tempSlice) == 2{
			value += tempSlice[0]
			tempSlice = tempSlice[1:]
		}
		if line == half {
			fmt.Println("half -- line", cycle, line,)

			break;
		} else {
			tempSlice = append(tempSlice, valueAndLine[line])
			fmt.Println(tempSlice, value)
		}
	}
	return value+1
}
