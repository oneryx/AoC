package main

import (
	"fmt"
	"strings"

	"github.com/oneryx/AoC/util"
)

var matrix = [][]int{}

// https://adventofcode.com/2018/day/3
func main() {
	strs := util.ReadAsSlice("../input.txt")
	w, h := matrixSize(strs)
	fmt.Printf("calculated matrix size: [%v, %v]\n", w, h)

	matrix = make([][]int, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]int, w)
	}

	for _, str := range strs {
		_, left, top, width, heigth := parseLine(str)
		updateMatrix(left, top, width, heigth)
	}

	square := 0
	for i := 0; i < w-1; i++ {
		for j := 0; j < h-1; j++ {
			if matrix[i][j] > 1 {
				// if number is greater than 1, it means the block is claimed by more than once
				square++
			}
		}
	}

	// printMatrix()

	fmt.Println(square)
}

// process each line of claim and update matrix
func updateMatrix(left, top, width, height int) {
	for i := left; i < left+width; i++ {
		for j := top; j < top+height; j++ {
			matrix[i][j] = matrix[i][j] + 1
		}
	}
}

func parseLine(str string) (id string, left, top, width, height int) {
	atIndex := strings.Index(str, "@")
	id = str[0 : atIndex-1]

	leftTopStr := str[atIndex+2 : strings.Index(str, ":")]
	leftTopArr := strings.Split(leftTopStr, ",")
	left = util.Atoi(leftTopArr[0])
	top = util.Atoi(leftTopArr[1])

	widthHeightStr := str[strings.Index(str, ":")+2:]
	widthHeightArr := strings.Split(widthHeightStr, "x")
	width = util.Atoi(widthHeightArr[0])
	height = util.Atoi(widthHeightArr[1])

	return id, left, top, width, height
}

// get a proper matrix size according to each claim
func matrixSize(strs []string) (int, int) {
	maxW := 0
	maxH := 0
	for _, str := range strs {
		_, left, top, width, height := parseLine(str)
		needW := left + width + 1
		needH := top + height + 1
		if needW > maxW {
			maxW = needW
		}
		if needH > maxH {
			maxH = needH
		}
	}
	return maxW, maxH
}

func printMatrix() {
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}
