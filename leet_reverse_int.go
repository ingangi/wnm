package main

import "fmt"

//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
//	For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.


func reverse(in int) (out int) {

	for i:=int(1); in != 0; i++{
		out = out*10 + in % 10
		in = in / 10
	}
	return
}

func main() {

	in := int(-120)
	fmt.Println(in, "reverse to", reverse(in))
}