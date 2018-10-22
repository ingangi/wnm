package main

import "fmt"

//Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.
//
//Note:
//
//Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
//
//Example 1:
//
//Input: [2,2,3,2]
//Output: 3
//Example 2:
//
//Input: [0,1,0,1,0,1,99]
//Output: 99

// 通过二进制位来解，目标是确定res的每个bit是0还是1
// 因为其余的数字都出现了三次，所以0-31的每个bit为1的个数之和应该能被3整除，否则该bit在res中为1
// 时间复杂度 O(32n)
func singleNumber(arr []int) (res int) {

	res = 0
	for i := 0; i < 32; i++ {
		count := 0    //第i bit为1的元素个数之和
		for _,v := range arr {
			if v & (1 << uint32(i)) > 0 {
				count++
			}
		}
		if count % 3 > 0 {
			res |= 1 << uint32(i)
		}
	}
	return
}

func main()  {
	arr := []int{0,1,0,1,0,1,99}
	fmt.Println(arr, "singleNumber:", singleNumber(arr))
}