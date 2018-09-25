package main

import (
	"fmt"
	"math"
)

//Given a binary tree, find its minimum depth.
//
//The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
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
//return its minimum depth = 2.


type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func depthMin(root *bTreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Min(float64(depthMin(root.l)), float64(depthMin(root.r))))
}

func main()  {
	root := &bTreeNode{1, nil, nil}
	root.l = &bTreeNode{2, nil, nil}
	root.r = &bTreeNode{2, nil, nil}
	root.l.l = &bTreeNode{3, nil, nil}
	root.l.r = &bTreeNode{3, nil, nil}
	root.l.l.l = &bTreeNode{4, nil, nil}
	root.l.l.r = &bTreeNode{4, nil, nil}
	fmt.Println(depthMin(root))
}
