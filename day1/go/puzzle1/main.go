package main

import (
	"os";
	"bufio";
	"strconv";
	"fmt";
)

var sum = 0
var maximum = 0

func main() {
	// read a text file
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd+"/day1/go/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse the line and add it to the sum
		// check if line is empty
		if len(scanner.Text()) > 0 {
			if line, err := strconv.Atoi(scanner.Text()); err == nil {
				sum += line
			}
		} else {
			if sum > maximum {
				maximum = sum
			}
			sum = 0
		}
	}

	
	// check for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(maximum)
}
