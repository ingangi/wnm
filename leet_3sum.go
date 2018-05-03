package main

import (
	"fmt"
	"sort"
	"math"
)

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

//Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
// Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//For example, given array S = {-1 2 1 -4}, and target = 1.
//
//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
func threeSumClosest(arr []int, target int) int {
	ret := 0
	if len(arr) < 3 {
		return ret
	}

	// 为了提高效率，先对arr进行升序排序
	// 再通过两端指针往中间收缩的方式，找到合适的值
	sort.Sort(sort.IntSlice(arr))

	ret = arr[0] + arr[1] + arr[2]
	diff := math.Abs(float64(ret-target))
	for i:=0; i<len(arr); i++ {
		left := i+1
		right := len(arr)-1

		for ;left<right; {
			tmp := arr[i]+arr[left]+arr[right]
			tmpdiff := math.Abs(float64(tmp-target))
			if tmpdiff < diff {
				ret = tmp
				diff = tmpdiff
			} else if tmp < target {
				left++
			} else if tmp > target {
				right--
			} else {
				return tmp
			}
		}
	}

	return ret
}

func main()  {
	//arr := []int{-1, 0, 1, 2, -1, -4}
	//threeSum(arr)

	arr := []int{-1, -4, 2, 1, -2}
	target := 1
	fmt.Println("arr for target:", arr, target, "is", threeSumClosest(arr, target))
}