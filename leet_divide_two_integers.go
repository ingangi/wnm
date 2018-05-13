package main

import (
	"fmt"
	"math"
)

//Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
//
//Return the quotient after dividing dividend by divisor.
//
//The integer division should truncate toward zero.
//
//Example 1:
//
//Input: dividend = 10, divisor = 3
//Output: 3
//Example 2:
//
//Input: dividend = 7, divisor = -3
//Output: -2
//Note:
//
//Both dividend and divisor will be 32-bit signed integers.
//The divisor will never be 0.
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

func devide(dvd int32, dvs int32) int32 {

	// 溢出情况
	if dvs ==0 || dvd == math.MinInt32 && dvs == -1 {
		return math.MaxInt32
	}

	// 是否负数
	isNeg := (dvd < 0) != (dvs < 0)

	if dvd < 0 {
		dvd = -dvd
	}
	if dvs < 0 {
		dvs = -dvs
	}

	var res int32
	for ;dvd >= dvs ;  {
		tmp := dvs
		mul := int32(1)

		// 左移一位相当于乘以2
		for ; dvd >= tmp << 1; {
			tmp <<= 1
			mul <<= 1
		}
		dvd -= tmp
		res += mul
	}
	if isNeg {
		res = -res
	}
	return res
}

func main() {
	dvd := -9
	dvs := 1
	fmt.Println(dvd, dvs, "devide is", devide(int32(dvd), int32(dvs)))
}