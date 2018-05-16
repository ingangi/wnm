package main

import (
	"fmt"
	"math"
)

//Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
//Example 1:
//
//Input: "(()"
//Output: 2
//Explanation: The longest valid parentheses substring is "()"
//Example 2:
//
//Input: ")()())"
//Output: 4
//Explanation: The longest valid parentheses substring is "()()"

// 模拟栈的pop
func popS(slc *[]int) int {
	if len(*slc) == 0 {
		return 0
	}

	c := (*slc)[len(*slc)-1]
	*slc = (*slc)[0:len(*slc)-1]
	return c
}

func longgestParentheses(par string) int {
	if len(par) < 2 {
		return 0
	}

	s := 0
	e := 0
	lmax := 0
	var stack []int

	for i,v := range par {
		if v == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				// 一次结算
				if e > s {
					lmax = int(math.Max(float64(e-s+1), float64(lmax)))
				}
				s = i+1
			}

			if len(stack) > 0 {
				popS(&stack)
				if len(stack) == 0 {
					e = i
				}
			}
		}
	}

	lmax = int(math.Max(float64(e-s+1), float64(lmax)))
	return lmax

}

func main()  {
	par := ")()())()()())("
	fmt.Println(par, "longgestParentheses =", longgestParentheses(par))
}