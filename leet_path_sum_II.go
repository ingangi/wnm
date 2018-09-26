package main

import (
	"fmt"
)

//Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given the below binary tree and sum = 22,
//
//		   5
//		  / \
//		 4   8
//		/   / \
//	   11  13  4
//	  /  \    / \
//	7    2  5   1
//Return:
//
//[
//[5,4,11,2],
//[5,8,4,5]
//]

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func pathSum(root *bTreeNode, path []int, presum int, target int) {
	if root == nil {
		return
	}

	path = append(path, root.v)
	presum += root.v

	if root.l == nil && root.r == nil {
		if presum == target {
			fmt.Println(path)
		}
		return
	}

	pathSum(root.l, path, presum, target)
	pathSum(root.r, path, presum, target)
}

func main()  {
	root := &bTreeNode{5, nil, nil}
	root.l = &bTreeNode{4, nil, nil}
	root.r = &bTreeNode{8, nil, nil}
	root.l.l = &bTreeNode{11, nil, nil}
	root.l.l.l = &bTreeNode{7, nil, nil}
	root.l.l.r = &bTreeNode{2, nil, nil}
	root.r.l = &bTreeNode{13, nil, nil}
	root.r.r = &bTreeNode{4, nil, nil}
	root.r.r.l = &bTreeNode{5, nil, nil}
	root.r.r.r = &bTreeNode{1, nil, nil}
	path := []int{}
	pathSum(root, path, 0, 22)
}
