package main

import (
	"fmt"
)

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//For example, given n = 3, a solution set is:
//
//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]
//
//vector<string> generateParenthesis(int n) {
//vector<string> vec;
//std::function<void(string, int, int)> addParenthesis = [&](string str, int ln, int rn) {
//if (ln == 0 && rn == 0) vec.push_back(str);
//if (ln > 0) addParenthesis(str+"(", ln-1, rn+1);
//if (rn > 0) addParenthesis(str+")", ln, rn-1);
//};
//addParenthesis("", n, 0);
//return vec;
//}

// ln rn分别为剩余的左右括号个数
func addParenthesis(str string, ln int, rn int, slc *[]string) {

	// 重点：任一时刻 左括号个数应该大于等于右括号
	// 当出现非法时  马上停止
	if ln > rn {
		return
	}

	if ln == 0 && rn == 0 {
		*slc = append(*slc, str)
		return
	}

	// 通过递归  穷举所有组合  通过前面的非法判断 剔除了错误的组合
	if ln > 0 {
		addParenthesis(str + "(", ln-1, rn, slc)
	}

	if rn > 0 {
		addParenthesis(str + ")", ln, rn-1, slc)
	}
}

func generateParenthesis(n int) ([]string){

	res := make([]string, 0)
	addParenthesis("", n, n, &res)
	return res
}

func main()  {

	n := 3
	fmt.Println(n, "generateParenthesis:", generateParenthesis(3))
}
