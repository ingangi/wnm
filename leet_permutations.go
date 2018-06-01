package main

import "fmt"

//Given a collection of distinct integers, return all possible permutations.
//
//Example:
//
//Input: [1,2,3]
//Output:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

func permutations(left []int, result []int)  {
	if len(left) <= 0 {
		fmt.Println("result:", result)
		return
	}

	for i,v := range left {
		// 用copy防止对slice递归时造成污染
		rc := make([]int, len(result))
		copy(rc, result)
		rc = append(rc, v)

		// 用copy防止对slice递归时造成污染
		lefttmp := make([]int, len(left[:i]))
		copy(lefttmp, left[:i])
		if i < len(left) - 1 {
			lefttmp = append(lefttmp, left[i+1:]...)
		}

		permutations(lefttmp, rc)
	}
}

func main()  {
	left := []int{1,2,3,4,5}
	result := make([]int, 0)

	permutations(left, result)
}