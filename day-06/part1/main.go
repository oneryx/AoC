package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/oneryx/AoC/util"
)

// each row stands for a point(x, y, area), area =-1 means infinite area, default area 0
var matrix [][3]int

// rows of matrix (or points count)
var count int

// https://adventofcode.com/2018/day/6
func main() {
	s := util.ReadAsSlice("../input.txt")
	for _, l := range s {
		arr := strings.Split(l, ",")
		matrix = append(matrix, [3]int{util.Atoi(strings.Trim(arr[0], " ")), util.Atoi(strings.Trim(arr[1], " ")), 0})
	}
	count = len(matrix)

	// find edges
	minX, maxX, minY, maxY := edges()

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			process(x, y, minX, maxX, minY, maxY)
		}
	}

	maxArea := 0
	for i := 0; i < count; i++ {
		// fmt.Printf("%v\n", matrix[i])
		if matrix[i][2] > maxArea {
			maxArea = matrix[i][2]
		}
	}

	fmt.Println(maxArea)
}

// distance from point (x, y) to point matrix[i]
func distance(x, y, i int) int {
	return util.Abs(x-matrix[i][0]) + util.Abs(y-matrix[i][1])
}

// find min/max x/y from all points
func edges() (minX, maxX, minY, maxY int) {
	minX, maxX, minY, maxY = matrix[0][0], matrix[0][0], matrix[0][1], matrix[0][1]
	for i := 1; i < count; i++ {
		if matrix[i][0] < minX {
			minX = matrix[i][0]
		}
		if matrix[i][0] > maxX {
			maxX = matrix[i][0]
		}
		if matrix[i][1] < minY {
			minY = matrix[i][1]
		}
		if matrix[i][1] > maxY {
			maxY = matrix[i][1]
		}
	}
	return minX, maxX, minY, maxY
}

// find closet point in matrix to point (x, y) and increase that point's area by 1 if found
// if the point is found having inifinite area, set area to -1
func process(x, y, minX, maxX, minY, maxY int) {
	minDistance := math.MaxInt32
	closestPointIndex := -1
	for i := 0; i < len(matrix); i++ {
		d := distance(x, y, i)
		if d < minDistance {
			minDistance = d
			closestPointIndex = i
		} else if d == minDistance {
			// if more than one points have the same minial distance to point (x, y), the area belongs to none of them
			closestPointIndex = -1
		}
	}
	if closestPointIndex == -1 || matrix[closestPointIndex][2] == -1 {
		// if current block has same distance to multiple points, there is no closest one, so no point should claim this bock
		// if the closest point was determined having infinite area
		return
	}
	if x == minX || x == maxX || y == minY || y == maxY {
		// if current point is on edge line, the close point will have infinite area, set area to -1
		matrix[closestPointIndex][2] = -1
		return
	}
	// increase area
	matrix[closestPointIndex][2] = matrix[closestPointIndex][2] + 1
}
