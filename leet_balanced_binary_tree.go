package main

import (
	"fmt"
	"math"
)

//Given a binary tree, determine if it is height-balanced.
//
//For this problem, a height-balanced binary tree is defined as:
//
//a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example 1:
//
//Given the following tree [3,9,20,null,null,15,7]:
//
//3
/// \
//9  20
///  \
//15   7
//Return true.
//
//Example 2:
//
//Given the following tree [1,2,2,3,3,null,null,4,4]:
//
//		 1
//		/ \
//	   2   2
//    / \
//   3   3
//  / \
// 4   4
//Return false.

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func depth(root *bTreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(depth(root.l)), float64(depth(root.r))))
}

func isBalancedBTree(root *bTreeNode) bool {
	if root == nil {
		return true
	}

	if math.Abs(float64(depth(root.l) - depth(root.r))) > 1 {
		return false
	}

	return isBalancedBTree(root.l) && isBalancedBTree(root.r)
}

func main()  {
	root := &bTreeNode{1, nil, nil}
	root.l = &bTreeNode{2, nil, nil}
	root.r = &bTreeNode{2, nil, nil}
	root.l.l = &bTreeNode{3, nil, nil}
	root.l.r = &bTreeNode{3, nil, nil}
	root.l.l.l = &bTreeNode{4, nil, nil}
	root.l.l.r = &bTreeNode{4, nil, nil}

	fmt.Println(isBalancedBTree(root))
	
}