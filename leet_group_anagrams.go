package main

import "fmt"

//Given an array of strings, group anagrams together.
//
//Example:
//
//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//Note:
//
//All inputs will be in lowercase.
//The order of your output does not matter.

func groupAnagrams(strs []string) map[int][]string {
	ret := make(map[int][]string)

	for _,str := range strs {
		// 为每个字符串生成特征值，使用位运算
		intv := 0
		for _,c := range str {
			intv |= 1 << uint32(c-'a')
		}
		ret[intv] = append(ret[intv], str)
	}
	return ret
}

func main()  {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}