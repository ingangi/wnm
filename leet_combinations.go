package main

import "fmt"

//Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//
//Example:
//
//Input: n = 4, k = 2
//Output:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
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

func combinations(n int, k int)  {
	left := make([]int, n)
	for i:=0; i < n; i++ {
		left[i] = i+1
	}

	for idx := 0; idx <= n-k; idx++ {
		fmt.Println(combin(left[idx:], k))
	}
}

func main()  {
	n := 4
	k := 2
	combinations(n, k)
}