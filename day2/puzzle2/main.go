package main

import (
	"os";
	"bufio";
	"fmt";
)

const roundWinPnt = 6
const roundLosePnt = 0
const roundDrawPnt = 3

// map of move and its value
var moveMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var rockPaperScissors = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
}

func calculateScore(elfMove string, command string) int {
	if command == "Win" {
		ptn := 0 
		if elfMove == "A" {
			ptn = moveMap["B"]
		} else if elfMove == "B" {
			ptn = moveMap["C"]
		} else if elfMove == "C" {
			ptn = moveMap["A"]
		}
		return roundWinPnt + ptn
	} else if command == "Lose" {
		ptn := 0 
		if elfMove == "A" {
			ptn = moveMap["C"]
		} else if elfMove == "B" {
			ptn = moveMap["A"]
		} else if elfMove == "C" {
			ptn = moveMap["B"]
		}
		return roundLosePnt + ptn
	} else {
		return roundDrawPnt + moveMap[elfMove]
	}
}

func main () {
	// read a text file
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// player scores
	myScore := 0

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// separate the line into two moves
		moves := scanner.Text()

		elfMove:= string(moves[0])
		command := string(moves[2])

		// calculate the score
		myScore += calculateScore(elfMove, rockPaperScissors[command])
	}

	// print my final score
	fmt.Println("my score: ", myScore)

	// check for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}