package main

import (
	"fmt"
)

//Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
//
//Example:
//
//Input: S = "ADOBECODEBANC", T = "ABC"
//Output: "BANC"
//Note:
//
//If there is no such window in S that covers all characters in T, return the empty string "".
//If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

func minimumWindowSubstring(S string, T string) string {

	// 将T中每个字符对应的ASC码作为下标，记录到数组，值代表字符的个数
	// asc码最多256个
	need := [256]int{}		//表示T中每个字符需要的个数
	found := [256]int{}		//遍历过程中，找到的个数

	for _,c := range T {
		need[int(c)]++
	}

	// 寻找过程, 用2个指针，
	// 移动第一个指针，从头开始遍历，找到第一个全包含的位置，我们把它定为end指针
	// end定了后，移动另一个指针，直到相对于end的最小窗口（就是无法再往右了），我们定为start
	start, end := 0, 0
	minstart, minend := 0, 0	//这两个做对比更小用
	foundcount := 0		//当这个值等于len(T)时，说明已经满足包含所有字符的要求
	for i,c := range S {
		if need[c] == 0 {
			// 不是目标字符
			continue
		}

		if found[c] <= need[c] {
			// 是需要的字符  并且这个字符还欠缺
			foundcount++
		}
		found[c]++

		if foundcount < len(T) {
			continue
		}

		// 到这里说明  已经包含了T的所有字符
		end = i
		// 开始将start向右移动
		for ; start < end; start++ {
			index := S[start]
			if need[index] == 0 {
				continue
			}
			if found[index] > need[index] {
				// 说明这个字符还有多余  可以继续向右
				found[index]--
				continue
			}

			break
		}

		// 到这里 找到了一个window  对比下是不是更小的window
		if minend > minstart && (minend - minstart) < (end - start) || start >= end{
		} else {
			minend = end
			minstart = start
		}
	}

	result := ""
	if minend > minstart {
		result = S[minstart:minend+1]
	}

	return result
}

func main()  {
	S := "ADOBECAODEBANC"
	T := "ABC"
	fmt.Println(S, T, "minimumWindowSubstring is", minimumWindowSubstring(S, T))
}