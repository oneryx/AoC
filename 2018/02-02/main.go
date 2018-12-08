package main

import (
	"fmt"

	"github.com/oneryx/AoC/2018/util"
)

// https://adventofcode.com/2018/day/2
func main() {
	strs := util.ReadAsSlice("input.txt")
	for i, str := range strs {
		b, common := alikeAny(str, strs[i+1:])
		if b {
			fmt.Printf("line %v, common: %v\n", i, common)
			return
		}
	}
	fmt.Println("nothing found")
}

// find any string in strs alike str
// return true, common part if found
func alikeAny(str string, strs []string) (bool, string) {
	for _, s := range strs {
		b, common := alike(str, s)
		if b {
			fmt.Printf("found string '%v' similar to input string '%v'\n", s, str)
			return true, common
		}
	}
	return false, ""
}

// check whether two strings are alike (only 1 char difference is allowed)
// if they are alike, return the common part
func alike(str1, str2 string) (bool, string) {
	// the two strings have the same length
	l := len(str1)
	diff := 0
	common := ""
	for i := 0; i < l; i++ {
		if str1[i] != str2[i] {
			diff++
			if diff > 1 {
				return false, ""
			}
		} else {
			common += string(str1[i])
		}
	}
	return true, common
}
