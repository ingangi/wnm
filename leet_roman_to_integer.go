package main

import "fmt"

//Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
//
//Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.Instead,
//the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
//I can be placed before V (5) and X (10) to make 4 and 9.
//X can be placed before L (50) and C (100) to make 40 and 90.
//C can be placed before D (500) and M (1000) to make 400 and 900.
//Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

//Input: "MCMXCIV"
//Output: 1994
//Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

var romMap = map[string]int{
	"I":1,
	"V":5,
	"X":10,
	"L":50,
	"C":100,
	"D":500,
	"M":1000,
}

func romanToInt(rom string) int {
	var ret int

	// 从右向左遍历字符串，取字符对应的数值累加
	// 当遍历到I时，若累积量小于5，说明I必定在右侧，为1，否则在左侧为-1
	// X,C同理
	l := len(rom)
	for i:=l-1; i >= 0; i-- {
		v,ok := romMap[string(rom[i])]
		isneg := 1
		if ok {
			switch string(rom[i]) {
			case "I":
				if ret >= 5 {
					isneg = -1
				}
			case "X":
				if ret >= 50 {
					isneg = -1
				}
			case "C":
				if ret >= 500 {
					isneg = -1
				}
			}
		}
		ret = ret + v * isneg
	}

	return ret
}

func main() {

	rom := "MCMXCIV"
	fmt.Println(rom, "to int:", romanToInt(rom))
}