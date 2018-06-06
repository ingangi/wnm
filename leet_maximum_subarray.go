package main

import (
	"fmt"
	"math"
)

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
//Example:
//
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.

func maximumSubarr(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	sum := arr[0]
	max := math.MinInt32
	for i:=1; i < len(arr); i++ {

		if sum + arr[i] < arr[i]{
			if sum <= arr[i] {
				sum = arr[i]
			}
		} else {
			sum += arr[i]
		}

		max = int(math.Max(float64(sum), float64(max)))
	}

	fmt.Println(max)
	return max
}

func main()  {
	arr := []int{-2,1,-3,4,-1,-1,2,1,-5,4}
	//arr := []int{-1,-2,-3}
	maximumSubarr(arr)
}