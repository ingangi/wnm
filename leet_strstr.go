package main

import "fmt"

//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

func strstr(haystack string, needle string) int {

	if len(haystack) == 0 {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	ih := 0
	in := 0
	index := -1
	for ;ih < len(haystack) && in < len(needle); ih++ {
		fmt.Println("cmp", string(haystack[ih]), ih, string(needle[in]), in)
		if haystack[ih] == needle[in] {
			if index == -1 {
				index = ih
			}
			in++
		} else if index != -1 {
			index = -1
			in = 0
			ih--	//已经匹配到过，又发现后面的不匹配，要回退这个不匹配的字符与needle首字符重新比较，否则会错过一个
		}
	}
	return index
}

func main() {

	haystack := "bbzbzx"
	needle := "bzx"
	fmt.Println(haystack, needle, "strstr is", strstr(haystack, needle))
}