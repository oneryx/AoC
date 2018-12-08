package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadAsSlice read file as slice of string
func ReadAsSlice(filepath string) []string {
	result := []string{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

// ReadAsIntSlice read file as slice of int
func ReadAsIntSlice(filepath string) []int {
	ints := []int{}
	strs := ReadAsSlice(filepath)
	for _, str := range strs {
		ints = append(ints, Atoi(str))
	}
	return ints
}

// Atoi wrapper of strconv.Atoi, just handled error
func Atoi(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
