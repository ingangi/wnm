package main

import (
	"fmt"
)

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func largestSubstringLen(s string) int {

	// 遍历一次
	// 使用map记录遍历的每个字符，比较是否出现过
	// 记录当前最长子串长途
	l := 0
	m := make(map[int32]int)
	len := 0
	for i,c := range s {
		_, exit := m[c]
		if exit {
			if len > l {
				l = len
			}
			len = 1

		} else {
			m[c] = i
			len++
		}
	}
	return l
}

func main()  {
	fmt.Println("len is", largestSubstringLen("1234ssqwertyu556"))
}