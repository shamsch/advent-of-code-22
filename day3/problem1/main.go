package main

import (
	"bufio"
	"fmt"
	"os"
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

func readInputFileLines(location string) []string {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd+location)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	lines := readInputFileLines("/../input.txt")
	sum := 0
	for _, line := range lines {
		halfLength := len(line) / 2
		halfLine := line[:halfLength]
		secondHalfLine := line[halfLength:]
		matchFound := false
		for _, char := range halfLine {
			if matchFound {
				break
			}
			for _, secondChar := range secondHalfLine {
				if char == secondChar {
					fmt.Println("Found match: ", string(char), "in", line)
					sum += values[string(char)]
					matchFound = true
					break
				}
			}
		}
	}
	println(sum)
}