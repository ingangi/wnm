package main

import "fmt"

//Given an array nums of n integers and an integer target, are there elements a, b, c,
//and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
//Note:
//
//The solution set must not contain duplicate quadruplets.
//
//Example:
//
//Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A solution set is:
//[
//[-1,  0, 0, 1],
//[-2, -1, 1, 2],
//[-2,  0, 0, 2]
//]

// arr中所有2个数相加为target的组合
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

// arr中所有3个数相加为target的组合
func threeSumForTarget(arr []int, target int) (output [][]int)  {
	l := len(arr)
	for i,v := range arr {
		if i >= l-2 {
			break
		}

		tmp := twoSumForTarget(arr[i+1:], target-v)
		if len(tmp) > 0 {
			for _,p := range tmp {
				output = append(output,[]int{v, p[0], p[1]})
			}
		}
	}
	return
}

func fourSumForTarget(arr []int, target int) (output [][]int)  {
	l := len(arr)
	for i,v := range arr {
		if i >= l-3 {
			break
		}
		tmp := threeSumForTarget(arr[i+1:], target-v)
		if len(tmp) > 0 {
			for _,p := range tmp {
				output = append(output,[]int{v, p[0], p[1], p[2]})
			}
		}
	}
	return
}

func main()  {
	arr := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(arr, "fourSumForTarget 0:", fourSumForTarget(arr, 0))
}