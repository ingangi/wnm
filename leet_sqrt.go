package main

import "fmt"

//Implement int sqrt(int x).
//
//Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
//
//Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
//
//Example 1:
//
//Input: 4
//Output: 2

func sqrt(x int) int {

	if x < 2 {
		return x
	}

	l := 1
	r := x
	ret := l
	for r >= l {
		m := (l+r) >> 1
		if m * m > x {
			r = m-1
		} else {
			ret = m
			l = m+1
		}
	}
	return ret
}

func main()  {
	x := 15
	fmt.Println(x, "sqrt is", sqrt(x))
}