package helper


import (
	"bufio"
	"os"
)

func ReadInputFileLines(location string) []string {
	file, err := os.Open(location)
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