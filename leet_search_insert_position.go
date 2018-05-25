package main

import "fmt"

//Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//
//Example 1:
//
//Input: [1,3,5,6], 5
//Output: 2

// 二分查找
func searchIndex(arr []int, target int) int {

	l := len(arr)
	high := l - 1
	low := 0
	for ;low <= high ;  {
		mid := (low + high)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func main() {

	arr := []int{1,3,4,5,7,9,22,45,65}
	target := 8
	fmt.Println(arr, "searchIndex", target, searchIndex(arr, target))
}