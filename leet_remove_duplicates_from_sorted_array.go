package main

import "fmt"

//Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
//
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
//Example 1:
//
//Given nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
//
//It doesn't matter what you leave beyond the returned length.

func removeDuplicates(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	w, last := 0, 0
	for r,v := range arr {
		if v != last {
			last = v
			if r != w {
				arr[w] = v
			}
			w++
		}
	}
	fmt.Println(arr, w)
	return w
}

func main() {
	arr := []int{1,1,2,4,5,5,3}
	removeDuplicates(arr)
}