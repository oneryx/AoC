package main

import (
	"fmt"
	"strings"

	"github.com/oneryx/AoC/2018/util"
)

// https://adventofcode.com/2018/day/2#part2
func main() {
	// line count that has letter repeats exactly two times
	twoCount := 0
	// line count that has letter repeats exactly three times
	threeCount := 0
	strs := util.ReadAsSlice("input.txt")

	for _, str := range strs {
		hasTwo, hasThree := hasLetterAppearTwoOrThreeTimes(str)
		if hasTwo {
			twoCount++
		}
		if hasThree {
			threeCount++
		}
	}
	fmt.Println(twoCount * threeCount)
}

// check whether a single letter appear in str exactly two or three times
// for return result, the first boolean is for two times and the second is for three times appearance
func hasLetterAppearTwoOrThreeTimes(str string) (bool, bool) {
	hasTwo := false
	hasThree := false
	m := map[string]bool{}
	for _, ascii := range str {
		c := string(ascii)
		if m[c] {
			// already counted for this letter
			continue
		}
		count := strings.Count(str, c)
		m[c] = true
		if count == 2 {
			hasTwo = true
		}
		if count == 3 {
			hasThree = true
		}
		if hasTwo && hasThree {
			// both found, no point to continue
			break
		}
	}
	return hasTwo, hasThree
}
