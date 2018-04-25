package main

import "fmt"

//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
//Example 1:
//
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//Example 2:
//
//Input: "cbbd"
//Output: "bb"

// 思路： 遍历每个字符，对该字符像两端延伸，寻找到最长的回文
// 需要区分奇数回文 和 偶数回文
func longgestOfChar(s string, begin int, end int, start *int, stop *int) {
	for {
		if begin < 0 || end > len(s)-1 {
			break
		}

		if s[begin] == s[end] {
			begin--
			end++
		} else {
			break
		}
	}

	// 回缩到有效范围
	begin++
	end--
	if end - begin > *stop - *start {
		*stop = end
		*start = begin
	}
}

func longgestOfString(s string) int {

	start := 0
	stop := 0
	for i, _ := range s {
		longgestOfChar(s, i, i, &start, &stop)		//奇数
		longgestOfChar(s, i, i+1, &start, &stop)	//偶数
	}

	fmt.Println(s[start:stop+1])
	return stop-start+1
}

func main() {
	s := "cbbd"
	fmt.Println("longgestOfString is", longgestOfString(s))
}