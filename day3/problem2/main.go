package main

import (
	"fmt"
	"github.com/shamsch/advent-of-code-22/helper"
)

var values = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}


type Character struct {
	Count int
	LastFoundInLine int
}

func findCommonCharacter(lines []string) string {
	var characters = make(map[string]Character)
	
	for i, line := range lines {
		for _, character := range line {
			char := string(character)
			if _, ok := characters[char]; ok {
				if characters[char].LastFoundInLine != i {
					characters[char] = Character{characters[char].Count + 1, i}
				}
			} else {
				characters[char] = Character{1, i}
			}
		}
	}

	for char, character := range characters {
		if character.Count == len(lines) {
			return char
		}
	}

	return ""
}

func main() {
	lines := helper.ReadInputFileLines("day3/input.txt")
	sum := 0
	groupOfThreeLines := []string{}
	for _, line := range lines {
		groupOfThreeLines = append(groupOfThreeLines, line)
		if len(groupOfThreeLines) == 3 {
			sum += values[findCommonCharacter(groupOfThreeLines)]
			groupOfThreeLines = []string{}
		}
	}
	fmt.Println(sum)
}