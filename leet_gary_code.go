package main

import "fmt"

//The gray code is a binary numeral system where two successive values differ in only one bit.
//
//Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.
//
//Example 1:
//
//Input: 2
//Output: [0,1,3,2]
//Explanation:
//00 - 0
//01 - 1
//11 - 3
//10 - 2
//
//For a given n, a gray code sequence may not be uniquely defined.
//For example, [0,2,3,1] is also a valid gray code sequence.
//
//00 - 0
//10 - 2
//11 - 3
//01 - 1

//利用模拟法来产生格雷码的时候有一个镜面定理，即当所有三位的格雷码可以由所有二位格雷码＋将二位所有格雷码从从最后一个到第一个依次最左边加１构成．
//例如二位的格雷码依次是00, 01, 11, 10．而三位的格雷码就是继承二位的格雷码＋逆序二位格雷码并在左边加１：110, 111, 101, 100．
//这样我们就可以得到所有的３位格雷为00, 01, 11, 10, 110, 111, 101, 100.

func garyCode(n int) (res []int) {
	res = append(res, 0)
	for i:=0; i < n; i++ {
		prelen := len(res)
		for j:=0; j<prelen; j++ {
			res = append(res, 1 << uint(i) + res[j])
		}
	}
	return 
}

func main()  {
	n := 2
	fmt.Println(n, ":", garyCode(n))
}