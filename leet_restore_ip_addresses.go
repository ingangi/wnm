package main

import "fmt"

//Given a string containing only digits, restore it by returning all possible valid IP address combinations.
//
//Example:
//
//Input: "25525511135"
//Output: ["255.255.11.135", "255.255.111.35"]

func dfsHelper(pre string, left string, index int, prenum int) bool {
	if index < 0 || index > 2 || index > len(left)-1 {
		return false
	}

	if left[0] == '0' {
		return false
	}

	n := 0
	for i := 0; i <= index; i++ {
		n = n*10 + int(left[i]-'0')
	}
	if n < 0 || n > 255 {
		return false
	}

	prestring := pre
	if prestring != "" {
		prestring += "." + left[:index+1]
	} else {
		prestring += left[:index+1]
	}

	if prenum == 3 {
		if index != len(left) - 1 {
			return false
		}
		fmt.Println("IP:", prestring)
		return true
	} else {
		prenum++
		dfsHelper(prestring, left[index+1:], 0, prenum)
		dfsHelper(prestring, left[index+1:], 1, prenum)
		dfsHelper(prestring, left[index+1:], 2, prenum)
	}
	return true
}

func restorIPAdress(input string) {
	dfsHelper("", input, 0, 0)
	dfsHelper("", input, 1, 0)
	dfsHelper("", input, 2, 0)
}

func main()  {
	restorIPAdress("12345678")
}
