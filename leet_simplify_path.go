package main

import "fmt"

//Given an absolute path for a file (Unix-style), simplify it.
//
//For example,
//path = "/home/", => "/home"
//path = "/a/./b/../../c/", => "/c"
//
//Corner Cases:
//
//Did you consider the case where path = "/../"?
//In this case, you should return "/".
//Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
//In this case, you should ignore redundant slashes and return "/home/foo".

func simplifyPath(p string) string {

	result := "/"
	ps := make([]string, 10)
	wIndex := 0
	rStart := 0
	rIndex := 0
	isPrePonit := 0  //前一个是否是.

	for i,c := range p {
		switch c {
		case '/':
			isPrePonit = 0
			if rIndex > 0 && rStart > 0 {
				ps[wIndex] = p[rStart : rIndex+1]
				wIndex++
				rIndex = 0
				rStart = 0
			}
		case '.':
			isPrePonit++
			if isPrePonit == 2 {
				// 往上一层
				wIndex--
			}
		default:
			if rStart == 0 {
				rStart = i
			}
			isPrePonit = 0
			rIndex = i
		}
	}

	for i:=0; i < wIndex; i++ {
		if i > 0 {
			result += "/"
		}
		result += ps[i]
	}

	return result
}

func main()  {
	p := "/a/./b/../../c/"
	fmt.Println(p, "simplifyPath", simplifyPath(p))
}