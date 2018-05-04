package main

import "fmt"

//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//Example:
//
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//Note:
//
//Although the above answer is in lexicographical order, your answer could be in any order you want.

var numMap = map[int32][]string{
	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	6: []string{"m", "n", "o"},
	7: []string{"p", "q", "r", "s"},
	8: []string{"t", "u", "v"},
	9: []string{"w", "x", "y", "z"},
}

func letterCombinations(num string) (ret []string) {
	// 把数字序列转化成n个字符数组序列
	var strInput [][]string
	for _, c := range num {
		strs, ok := numMap[c-'0']
		if ok {
			strInput = append(strInput, strs)
		}
	}

	// 组合
	ret = append(ret, "")
	for _, slc := range strInput {  //[]string{"a", "b", "c"},
		var tmp []string
		for _, c := range slc {  //"a", "b", "c"
			for _, pre := range ret {
				tmp = append(tmp, pre+c)
			}
		}
		ret = tmp
	}
	return 
}

func main()  {
	num := "584"
	fmt.Println(num, "letterCombinations:", letterCombinations(num))
}