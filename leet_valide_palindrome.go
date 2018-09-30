package main

import "fmt"

//Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//
//Note: For the purpose of this problem, we define empty string as valid palindrome.
//
//Example 1:
//
//Input: "A man, a plan, a canal: Panama"
//Output: true
//Example 2:
//
//Input: "race a car"
//Output: false

// 判断是否是字母和数字，并且把字母转换为小写
func isAlphanumericAndLowcap(c int32) (bool, int32) {
	if c >= '0' && c <= '9' {
		return true,c
	}
	if c >= 'A' && c <= 'Z' {
		return true, c+'a'-'A'
	}
	if c >= 'a' && c <= 'z' {
		return true,c
	}
	return false,c
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	left, right := 0, l-1
	lc, rc, isAlph := int32(0), int32(0), false
	for left <= right {
		isAlph, lc = isAlphanumericAndLowcap(int32(s[left]))
		if !isAlph {
			left++
			continue
		}
		isAlph, rc = isAlphanumericAndLowcap(int32(s[right]))
		if !isAlph {
			right--
			continue
		}
		if lc != rc {
			return false
		}
		left++
		right--
	}
	return true
}

func main()  {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
