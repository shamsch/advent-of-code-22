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

func main () {
	// read a text file
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd+"/../input.txt")
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
		if moveMap[elfMove] > moveMap[myMove] {
			fmt.Println("Elf wins", elfMove, ">", myMove)
			elfScore += roundWinPnt
		} else if moveMap[elfMove] < moveMap[myMove] {
			fmt.Println("I win", myMove, ">", elfMove)
			myScore += roundWinPnt
		} else if moveMap[elfMove] == moveMap[myMove] {
			fmt.Println("Draw", elfMove, "=", myMove)
			elfScore += roundDrawPnt
			myScore += roundDrawPnt
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