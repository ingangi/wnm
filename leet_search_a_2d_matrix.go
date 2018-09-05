package main

import "fmt"

//Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//
//Integers in each row are sorted from left to right.
//The first integer of each row is greater than the last integer of the previous row.
//Example 1:
//
//Input:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 3
//Output: true
//Example 2:
//
//Input:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 13
//Output: false

func find2dMatrix(matrix [][]int, v int) bool {
	m := len(matrix)
	n := len(matrix[0])
	max := m * n
	if max == 0 {
		return false
	}

	high := max - 1
	low := 0

	for {
		// 将二维数组想象成一个连续的一维数组，根据题意是已经排好序的
		i := (high + low) / 2
		r := i / n
		c := i % n
		if matrix[r][c] == v {
			return true
		}

		if matrix[r][c] > v {
			high = i - 1
		} else {
			low = i + 1
		}

		if low > high {
			return false
		}
	}
	return false
}

func main()  {
	matrix := [][]int{[]int{1, 3, 5, 7},[]int{10, 11, 16, 20},[]int{23, 30, 34, 50}}
	v := 3
	fmt.Println(matrix, v, "result is", find2dMatrix(matrix, v))
}