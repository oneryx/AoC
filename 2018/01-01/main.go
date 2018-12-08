package main

import (
	"github.com/oneryx/AoC/2018/util"
)

// https://adventofcode.com/2018/day/1
func main() {
	sum := 0
	ints := util.ReadAsIntSlice("input.txt")
	for _, v := range ints {
		sum += v
	}
	println(sum)
}
