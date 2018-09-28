package main

import "fmt"

//Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
//
//Note that the row index starts from 0.
//
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 3
//Output: [1,3,3,1]
//Follow up:
//
//Could you optimize your algorithm to use only O(k) extra space?

// 比I的版本多了个空间的要求
// 使用2个数组，一个保存前一行，一个计算当前行
func pascalTriangleII(k int) []int {
	buf1, buf2 := make([]int, k+1), make([]int, k+1)
	cur, last := &buf1, &buf2
	for i := 0; i < k+1; i++ {
		(*cur)[0], (*cur)[i] = 1, 1
		if i > 1 {
			for j:=1; j<i; j++ {
				(*cur)[j] = (*last)[j] + (*last)[j-1]
			}
		}
		cur, last = last, cur
	}
	return (*last)[:k+1]
}

func main()  {
	fmt.Println(pascalTriangleII(3))
}