package main

import (
	"fmt"
	"math"
)

//Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
//n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.

// 题意
// i为x轴（0,1,2,3,4...）  a为Y轴   N个点(i, ai)  每个点向X轴做垂直的线段  那么两条线段加上它们之间的X轴  构成了一个容器
// 找到最大容积的容器，返回两个a值

func maxContain(a []int) (begin int, end int, s int) {
	if len(a) < 2 {
		return
	}

	b := 0
	e := len(a) - 1

	for ;e>b; {
		si := (e - b) * int(math.Min(float64(a[e]), float64(a[b])))
		if si > s {
			s = si
			begin = b
			end = e
		}

		// O(n)的关键，从两侧往中间搜索各个容器，只收缩线段较短的那一侧，这样不会错过较大的那些容器
		if a[e] > a[b] {
			b++
		} else {
			e--
		}
	}

	return
}

func main() {

	a := []int{3,4,3,8}
	b,e,s := maxContain(a)
	fmt.Println(a,b,e,s)
}