package main

import "fmt"

//Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
//
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
//Example 1:
//
//Given nums = [1,1,1,2,2,3],
//
//Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
//
//It doesn't matter what you leave beyond the returned length.

func remoteDuplicates2(arr []int) int {

	w := 0
	pre := -999
	repeats := 0
	wstop := false
	for i,v := range arr {
		if !wstop {
			w = i
		}

		if v != pre {
			repeats = 1
			if i != w {
				arr[w] = v
				w++
			}
		} else {
			repeats++
			if repeats > 2 {
				if w == i {
					wstop = true
				}
			} else {
				if i != w {
					arr[w] = v
					w++
				}
			}

		}

		pre = v
	}

	if !wstop {
		w++
	}
	return w
}

func main()  {
	nums := []int{1,1,1,2,2,2,2,2,2,3,3,4,4,4,4,5,5}
	len := remoteDuplicates2(nums)
	fmt.Println(len, nums[0:len])
}