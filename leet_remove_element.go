package main

import "fmt"

//Given an array nums and a value val, remove all instances of that value in-place and return the new length.
//
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
//The order of elements can be changed. It doesn't matter what you leave beyond the new length.
//
//Example 1:
//
//Given nums = [3,2,2,3], val = 3,
//
//Your function should return length = 2, with the first two elements of nums being 2.
//
//It doesn't matter what you leave beyond the returned length.

func removeElem(arr []int, elm int) int {
	w := 0

	len := len(arr)

F:	for r:=0 ;r < len && w < len ; r++{
		if arr[r] == elm {
			for {
				r++
				if r > len -1 {
					break F
				}

				if arr[r] != elm {
					break
				}
			}

			// 交换w r
			arr[w], arr[r] = arr[r], arr[w]
			r--
		}
		w++
	}

	fmt.Println(arr, w)
	return w
}

func main()  {
	arr := []int{3,3,1,2,3,2,3}
	removeElem(arr, 3)
}