package main

import "fmt"

//
//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
//
//You are given a target value to search. If found in the array return true, otherwise return false.
//
//Example 1:
//
//Input: nums = [2,5,6,0,0,1,2], target = 0
//Output: true
//Example 2:
//
//Input: nums = [2,5,6,0,0,1,2], target = 3
//Output: false

func searchInRotatedSortedArray2(nums []int, target int) bool {
	// 二分查找
	// 每次取mid，mid的两边必然有一边是有序的，只要有一边有序，我们就能确定target在mid的哪一侧，实现折半的目的
	left, right := 0, len(nums)-1
	for mid := (left+right)/2; left <= right; mid = (left+right)/2 {
		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[left] { // 左半边有序
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] { // 右半边有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++  //说明mid和left之间都是重复数据，缩小范围
		}
	}
	return false
}

func main()  {
	nums := []int{2,5,6,0,0,1,2}
	target := 0
	fmt.Println(nums, "has", target, "?", searchInRotatedSortedArray2(nums, target))
}