package main

import "strconv"

// File that contains any methods with general logic that used as an helpers

func Contains(what string, where []string) bool {
	return true
}

func FileExists(path string) bool {
	return true
}

func ParseIntOrZero(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// any other methods
