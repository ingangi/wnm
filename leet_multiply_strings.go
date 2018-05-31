package main

import (
	"fmt"
)

//Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
//
//Example 1:
//
//Input: num1 = "2", num2 = "3"
//Output: "6"
//Example 2:
//
//Input: num1 = "123", num2 = "456"
//Output: "56088"

func multiplyStrings(num1 string, num2 string) (sum []int) {
	len1 := len(num1)
	len2 := len(num2)

	sum = make([]int, len1+len2)	//两数相乘结果的位数 不会大于两个数各自位数之和
	for i:=len1-1; i>=0; i-- {
		tmp1 := int(num1[i] - '0')
		car := 0	//进位

		for j:=len2-1; j>=0; j-- {
			tmp2 := int(num2[j] - '0')

			sumt := sum[i+j+1] + tmp1 * tmp2 + car
			sum[i+j+1] = sumt % 10
			car = sumt / 10
		}

		sum[i] += car
	}
	fmt.Println(sum)

	return sum
}

func main()  {
	num1 := "123"
	num2 := "456"
	fmt.Println(num1, num2, "multiplyStrings is", multiplyStrings(num1, num2))
}
