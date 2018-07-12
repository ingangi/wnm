package main

import "fmt"

//You are climbing a stair case. It takes n steps to reach to the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//Note: Given n will be a positive integer.
//
//Example 1:
//
//Input: 2
//Output: 2
//Explanation: There are two ways to climb to the top.
//1. 1 step + 1 step
//2. 2 steps

var t int

func climbStair(left int)  {
	if left <= 0 {
		t++
		return
	}

	if left > 1 {
		climbStair(left - 2)
	}

	climbStair(left - 1)
}

func main()  {
	left := 4
	climbStair(left)
	fmt.Println("climbStair", left, t)
}