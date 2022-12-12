package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shamsch/advent-of-code-22/helper"
)

func main() {
	lines := helper.ReadInputFileLines("day10/input.txt")
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
	value = calculateValue(5, valueAndLine)
	fmt.Println(strength, value)
}

func calculateValue(cycle int, valueAndLine map[int]int) int {
	value := 0
	for i:=1; i<=cycle; i++ {
		for line, strength := range valueAndLine {
			if line+2 == i {
				value += strength
			}
		}
	}
	return value+1
}
