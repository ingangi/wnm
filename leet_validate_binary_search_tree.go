package main

import "fmt"

//Given a binary tree, determine if it is a valid binary search tree (BST).
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.
//		Example 1:
//
//		Input:
//		 2
//		/ \
//	   1   3
//		Output: true
//		Example 2:
//
//		 5
//		/ \
//	  1   4
//		 / \
//	    3   6
//		Output: false

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}
var startCheck bool
var preV int

func checkData(v int) bool {
	if startCheck {
		if preV >= v {
			return false
		}
	} else {
		startCheck = true
	}

	preV = v
	return true
}

// 中序遍历  应该是递增的
func iterBStreeInorder(node *bTreeNode) bool {
	if node == nil {
		return true
	}
	if !iterBStreeInorder(node.l) {
		return false
	}
	if !checkData(node.v) {
		return false
	}
	if !iterBStreeInorder(node.r) {
		return false
	}
	return true
}

func main()  {
	root := &bTreeNode{2, nil, nil}
	root.l = &bTreeNode{1, nil, nil}
	root.r = &bTreeNode{4, nil, nil}
	root.r.l = &bTreeNode{3, nil, nil}
	root.r.r = &bTreeNode{6, nil, nil}
	fmt.Println(iterBStreeInorder(root))
}

