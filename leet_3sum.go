package main

import "fmt"

//Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
//Note:
//
//The solution set must not contain duplicate triplets.
//
//Example:
//
//Given array nums = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func twoSumForTarget(arr []int, target int) (output [][]int) {

	tmpmap := make(map[int]int)
	for i,v := range arr {
		j, exist := tmpmap[v]
		if exist {
			output = append(output,[]int{arr[j], arr[i]})
		} else {
			tmpmap[target-v] = i
		}
	}

	return 
}

func threeSum(arr []int)  {
	// 从左往右遍历，第i个元素+后面的twosum结果为0  则输出结果
	l := len(arr)
	for i,v := range arr {
		if i >= l-2 {
			break
		}

		output := twoSumForTarget(arr[i+1:], 0-v)
		if len(output) > 0 {
			for _,p := range output {
				fmt.Println(v, p[0], p[1])
			}
		}
	}
}

func main()  {
	arr := []int{-1, 0, 1, 2, -1, -4}
	threeSum(arr)
}