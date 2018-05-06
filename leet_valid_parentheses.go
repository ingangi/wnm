package main

import "fmt"

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.
//
//Example 1:
//
//Input: "()"
//Output: true
//Example 2:
//
//Input: "()[]{}"
//Output: true
//Example 3:
//
//Input: "(]"
//Output: false
//Example 4:
//
//Input: "([)]"
//Output: false
//Example 5:
//
//Input: "{[]}"
//Output: true

// 左括号入栈  右括号出栈
func isValid(s string) bool {
	if len(s) == 0 || len(s) % 2 != 0 {
		return false
	}

	var stack []string
	for _,c := range s {
		switch string(c) {
		case "(":
			fallthrough
		case "[":
			fallthrough
		case "{":
			stack = append(stack, string(c))

		case ")":
			if stack[len(stack) - 1] != "(" {
				return false
			}
			stack = stack[0:len(stack)-1]
		case "]":
			if stack[len(stack) - 1] != "[" {
				return false
			}
			stack = stack[0:len(stack)-1]
		case "}":
			if stack[len(stack) - 1] != "{" {
				return false
			}
			stack = stack[0:len(stack)-1]
		default:
			return false

		}
	}
	return true
}

func main()  {
	s := "{[{(())}]}"
	fmt.Println(s, "isValid:", isValid(s))
}