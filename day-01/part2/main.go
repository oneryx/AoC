package main

import (
	"errors"
	"fmt"

	util "github.com/oneryx/AoC/2018/util"
)

var sum = 0

// records appearances of each number
var m = map[int]int{0: 1}

// https://adventofcode.com/2018/day/1#part2
func main() {
	ints := util.ReadAsIntSlice("../input.txt")
	for {
		result, err := loop(ints)
		if err == nil {
			fmt.Println(result)
			break
		}
	}
}

func loop(ints []int) (int, error) {
	for _, v := range ints {
		sum += v
		m[sum] = m[sum] + 1
		if m[sum] == 2 {
			return sum, nil
		}
	}
	return 0, errors.New("not found yet")
}
