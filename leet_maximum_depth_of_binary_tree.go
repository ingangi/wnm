package main

import "fmt"

//Given a binary tree, find its maximum depth.
//
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given binary tree [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//return its depth = 3.

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func maxDepthOfBTree(root *bTreeNode) int {
	if root == nil {
		return 0
	}

	ldepth := maxDepthOfBTree(root.l) + 1
	rdepth := maxDepthOfBTree(root.r) + 1

	if ldepth > rdepth {
		return ldepth
	}

	return rdepth
}

func main()  {
	root := &bTreeNode{2, nil, nil}
	root.l = &bTreeNode{1, nil, nil}
	root.r = &bTreeNode{1, nil, nil}
	root.r.l = &bTreeNode{4, nil, nil}
	root.r.r = &bTreeNode{3, nil, nil}
	root.l.l = &bTreeNode{3, nil, nil}
	root.l.l.r = &bTreeNode{4, nil, nil}
	root.l.l.r.l = &bTreeNode{4, nil, nil}
	fmt.Println(maxDepthOfBTree(root))
}