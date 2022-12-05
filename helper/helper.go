package helper

import (
	"bufio"
	"os"
	"sort"
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

func SliceContains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ReverseSlice[T comparable] (slice []T) []T {
	sort.SliceStable(slice, func(i, j int) bool {
        return i > j
    })
	return slice
}