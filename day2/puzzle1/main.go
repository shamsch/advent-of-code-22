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
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var rockPaperScissors = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
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
	elfScore := 0

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// separate the line into two moves
		moves := scanner.Text()

		elfMove:= string(moves[0])
		myMove := string(moves[2])

		// calculate the score
		if rockPaperScissors[elfMove] == rockPaperScissors[myMove] {
			elfScore += roundDrawPnt
			myScore += roundDrawPnt
		} else if rockPaperScissors[elfMove] == "Rock" && rockPaperScissors[myMove] == "Paper" {
			elfScore += roundLosePnt
			myScore += roundWinPnt
		} else if rockPaperScissors[elfMove] == "Rock" && rockPaperScissors[myMove] == "Scissors" {
			elfScore += roundWinPnt
			myScore += roundLosePnt
		} else if rockPaperScissors[elfMove] == "Paper" && rockPaperScissors[myMove] == "Rock" {
			elfScore += roundWinPnt
			myScore += roundLosePnt
		} else if rockPaperScissors[elfMove] == "Paper" && rockPaperScissors[myMove] == "Scissors" {
			elfScore += roundLosePnt
			myScore += roundWinPnt
		} else if rockPaperScissors[elfMove] == "Scissors" && rockPaperScissors[myMove] == "Rock" {
			elfScore += roundLosePnt
			myScore += roundWinPnt
		} else if rockPaperScissors[elfMove] == "Scissors" && rockPaperScissors[myMove] == "Paper" {
			elfScore += roundWinPnt
			myScore += roundLosePnt
		}

		// add the score of the move to the player's score
		elfScore += moveMap[elfMove]
		myScore += moveMap[myMove]
	}

	// print my final score
	fmt.Println("my score: ", myScore)

	// check for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}