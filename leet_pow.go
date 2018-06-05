package main

import "fmt"
import "time"

//Implement pow(x, n), which calculates x raised to the power n (xn).
//
//Example 1:
//
//Input: 2.00000, 10
//Output: 1024.00000

// 硬算
//func mypow(x float64, n int) float64 {
//
//	res := float64(1)
//	for i:=0;i<n ;i++  {
//		res = res * x
//	}
//	return res
//}

// 二分法
//x^n = x^(n/2) * x^(n/2) --- n 为偶数
//x^n = x * x^(n/2) * x^(n/2) --- n 为奇数
func mypow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	tmp := mypow(x, n/2)
	if n & 1 == 0 {
		return tmp * tmp
	} else {
		return x * tmp * tmp
	}
}

func mypow_test()  {

	start := time.Now().Nanosecond()

	for i:=0;i<1000000 ;i++  {
		mypow(2,10)
	}

	end := time.Now().Nanosecond()
	fmt.Println(end - start)
}

func main()  {

	x := float64(2)
	n := 10
	fmt.Println(mypow(x,n))

	//mypow_test()
}