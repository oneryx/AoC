package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// WriteToFile write string slice to file
func WriteToFile(strs []string, filepath string) {
	f, err := os.Create(filepath)
	defer f.Close()
	CheckErr(err)
	w := bufio.NewWriter(f)
	for _, str := range strs {
		_, err := w.WriteString(str + "\n")
		CheckErr(err)
	}
	w.Flush()
}

// ReadAsSlice read file as slice of string
func ReadAsSlice(filepath string) []string {
	b, err := ioutil.ReadFile(filepath)
	CheckErr(err)
	return strings.Split(string(b), "\n")
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
	CheckErr(err)
	return v
}

// CheckErr log fatal if error
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Abs Abs
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
