package main

import "fmt"

//Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
//
//Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
//
//Note: You are not suppose to use the library's sort function for this problem.
//
//Example:
//
//Input: [2,0,2,1,1,0]
//Output: [0,0,1,1,2,2]
//Follow up:
//
//A rather straight forward solution is a two-pass algorithm using counting sort.
//First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
//Could you come up with a one-pass algorithm using only constant space?

func sortColorsOnePass(colors []int)  {
	if len(colors) == 0 {
		return
	}

	// 以1为轴  0在左边  2在右边
	left := 0
	right := len(colors) - 1
	fix1 := -1
	for {
		if colors[left] == 2 {
			colors[left], colors[right] = colors[right], colors[left]
			right--
		} else {
			if colors[left] == 1 && fix1 == -1 {
				fix1 = left
			} else if fix1 >= 0 && colors[left] == 0 {
				colors[left], colors[fix1] = colors[fix1], colors[left]
				left, fix1 = fix1+1, -1
				continue
			}

			left++
		}

		if left > right {
			break
		}
	}
}

func main()  {
	colors := []int{1,0,1,1,2,0,1}
	sortColorsOnePass(colors)
	fmt.Println(colors)
}
