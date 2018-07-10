package main

import (
	"fmt"
)

//Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
//
//The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
//
//You may assume the integer does not contain any leading zero, except the number 0 itself.

func plusOne(input []int) []int {

	carry := 0

	for i:=len(input)-1; i>=0; i-- {
		plus := carry
		if i == len(input)-1 {
			plus = 1
		}

		input[i] += plus
		if input[i] > 9 {
			input[i] %= 10
			carry = 1
		} else {
			carry = 0
			break
		}
	}

	// 说明溢出了 要增加位数  最高位为carry
	if carry > 0 {
		input = append([]int{carry}, input...)
	}

	return input
}

func main()  {
	input := []int{9,9,9}
	fmt.Println("add one is", plusOne(input))
}