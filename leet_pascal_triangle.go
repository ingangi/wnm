package main

import "fmt"

//Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
//
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 5
//Output:
//[
//[1],
//[1,1],
//[1,2,1],
//[1,3,3,1],
//[1,4,6,4,1]
//]

func pascalTriangle(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		if i > 1 {
			for j:=1; j<i; j++ {
				result[i][j] = result[i-1][j] + result[i-1][j-1]
			}
		}
	}
	return result
}

func main()  {
	fmt.Println(pascalTriangle(10))
}