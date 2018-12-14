package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("../input.txt")
	s := []byte{}
	for _, b := range bytes {
		end := len(s) - 1
		if len(s) == 0 || (s[end]-b != 32 && b-s[end] != 32) {
			s = append(s, b)
		} else {
			s = s[:end]
		}
	}
	fmt.Println(len(s))
}
