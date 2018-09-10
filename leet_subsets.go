package main

import "fmt"

//Given a set of distinct integers, nums, return all possible subsets (the power set).
//
//Note: The solution set must not contain duplicate subsets.
//
//Example:
//
//Input: nums = [1,2,3]
//Output:
//[
//[3],
//[1],
//[2],
//[1,2,3],
//[1,3],
//[2,3],
//[1,2],
//[]
//]

func combin(left []int, k int) [][]int {
	if k == 1 {
		tmp := make([][]int, len(left))
		for i, v := range left {
			tmp[i] = []int{v}
		}
		return tmp
	}

	c := combin(left[1:], k-1)
	for idx, _ := range c {
		c[idx] = append(c[idx], left[0])
	}
	return c
}

func subsets(left []int)  {
	n := len(left)

	for k := 1; k <= n; k++ {
		for idx := 0; idx <= n-k; idx++ {
			fmt.Println(combin(left[idx:], k))
			if k == 1 {
				break
			}
		}
	}
}

func main()  {
	left := []int{1,2,3,4,5}
	subsets(left)
}