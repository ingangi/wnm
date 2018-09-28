package main

import (
	"fmt"
)

//
//Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
//
//For example, given the following triangle
//
//[
//[2],
//[3,4],
//[6,5,7],
//[4,1,8,3]
//]
//The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
//
//Note:
//
//Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

func dfsHelper(triangle [][]int, maxrow int, curlow int, lastindex int, presum int) int {
	if curlow > maxrow {
		return presum
	}
	curindex := 0
	if triangle[curlow][lastindex] > triangle[curlow][lastindex+1] {
		curindex = lastindex+1
	} else {
		curindex = lastindex
	}
	return dfsHelper(triangle, maxrow, curlow+1, curindex, presum + triangle[curlow][curindex])
}

func triangleMinPath(triangle [][]int) int {
	return dfsHelper(triangle, len(triangle) - 1, 1, 0, triangle[0][0])
}

func main()  {
	triangle := [][]int{[]int{2},[]int{3,1},[]int{6,5,1},[]int{4,1,8,1}}
	fmt.Println(triangleMinPath(triangle))
}