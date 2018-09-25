package main

import "fmt"

//Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example:
//
//Given the sorted array: [-10,-3,0,5,9],
//
//One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
//	   0
//	  / \
//	-3   9
//	/   /
//  -10  5

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func printBtreeMiddle(node *bTreeNode)  {
	if node == nil {
		return
	}
	printBtreeMiddle(node.l)
	fmt.Println(node.v)
	printBtreeMiddle(node.r)
}

func convSortedArrToBSTree(arr []int) *bTreeNode {

	l := len(arr)
	if l == 0 {
		return nil
	}

	index := l / 2
	root := &bTreeNode{arr[index], nil, nil}
	if index > 0 {
		root.l = convSortedArrToBSTree(arr[:index])
	}

	if index < l - 1 {
		root.r = convSortedArrToBSTree(arr[index+1:])
	}

	return root
}

func main()  {
	arr := []int{-10,-3,0,5,9}
	printBtreeMiddle(convSortedArrToBSTree(arr))
}

