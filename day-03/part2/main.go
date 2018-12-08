package main

import (
	"fmt"
	"strings"

	"github.com/oneryx/AoC/2018/util"
)

type claim struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

// https://adventofcode.com/2018/day/3#part2
func main() {
	strs := util.ReadAsSlice("../input.txt")
	claims := []claim{}
	for _, str := range strs {
		claims = append(claims, parseLine(str))
	}

	for _, claim := range claims {
		if !overlapAny(claim, claims) {
			fmt.Println(claim.id)
		}
	}
}

func overlapAny(c claim, claims []claim) bool {
	for _, claim := range claims {
		if c.id == claim.id {
			// skip the claim itself
			continue
		}
		if overlap(c, claim) {
			return true
		}
	}
	return false
}

// check whether two claim has any overlap
func overlap(c1, c2 claim) bool {
	if c1.left+c1.width < c2.left {
		return false
	}
	if c1.top+c1.height < c2.top {
		return false
	}
	if c2.left+c2.width < c1.left {
		return false
	}
	if c2.top+c2.height < c1.top {
		return false
	}
	return true
}

func parseLine(str string) claim {
	c := claim{}
	atIndex := strings.Index(str, "@")
	c.id = str[0 : atIndex-1]

	leftTopStr := str[atIndex+2 : strings.Index(str, ":")]
	leftTopArr := strings.Split(leftTopStr, ",")
	c.left = util.Atoi(leftTopArr[0])
	c.top = util.Atoi(leftTopArr[1])

	widthHeightStr := str[strings.Index(str, ":")+2:]
	widthHeightArr := strings.Split(widthHeightStr, "x")
	c.width = util.Atoi(widthHeightArr[0])
	c.height = util.Atoi(widthHeightArr[1])

	return c
}
