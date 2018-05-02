package main

import "fmt"

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//Input: ["flower","flow","flight"]
//Output: "fl"
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//Note:
//
//All given inputs are in lowercase letters a-z.

func commonPrefix(str1 string, str2 string) string {

	i := 0
	for ; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			break
		}
	}

	ret := ""
	if i > 0 {
		ret = str1[:i]
	}

	return ret
}

func longestCommonPrefix(strs []string) string {

	ret := ""
	for i:=0; i<len(strs)-1;i++{
		tmp := commonPrefix(strs[i], strs[i+1])
		if tmp == "" {
			return tmp
		}

		if len(tmp) < len(ret) || ret == "" {
			ret = tmp
		}
	}
	return ret
}

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(strs, "longestCommonPrefix is", longestCommonPrefix(strs))
}