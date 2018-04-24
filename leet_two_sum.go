package main

import "fmt"

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(input []int, target int) (output []int) {

	// 思路1  直接循环遍历  O(n2)
	//for i,v := range input {
	//	for j:=i+1;j<len(input);j++ {
	//		if v + input[j] == target {
	//			output = []int{i,j}
	//			return
	//		}
	//	}
	//}

	// 思路2  用一个hashmap将每个元素与target的差值存起来，遍历一遍input就能发现目标值  O(n)
	tmpmap := make(map[int]int)
	for i,v := range input {
		j, exist := tmpmap[v]
		if exist {
			output = []int{j, i}
			return
		} else {
			tmpmap[target-v] = i
		}
	}

	return
}

func main()  {
	input := []int{11, 9, 7, 8, 1, 0}
	target := 9

	output := twoSum(input, target)

	fmt.Println(output)
}