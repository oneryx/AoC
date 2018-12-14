package main

import (
	"fmt"
	"io/ioutil"
)

// https://adventofcode.com/2018/day/5#part2
func main() {
	bytes, _ := ioutil.ReadFile("../input.txt")
	alp := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	shortest := len(bytes)
	for _, b := range alp {
		result := lenAfterReact(bytes, b)
		if result < shortest {
			shortest = result
		}
	}
	fmt.Println(shortest)
}

func lenAfterReact(bytes []byte, x byte) int {
	s := []byte{}
	for _, b := range bytes {
		if b == x || b == x+32 {
			continue
		}
		end := len(s) - 1
		if len(s) == 0 || (s[end]-b != 32 && b-s[end] != 32) {
			s = append(s, b)
		} else {
			s = s[:end]
		}
	}
	return len(s)
}
