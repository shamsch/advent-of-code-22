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
	for i:=20; i<=220; i+=40{
		strength += valueArray[i-1]*i
	}
	fmt.Println(strength)
}

func calculateValue()  {
	
}
