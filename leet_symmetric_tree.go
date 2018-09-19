package main

import "fmt"

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//	For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//	   1
//	  / \
//	 2   2
//	/ \ / \
//	3  4 4  3
//	But the following [1,2,2,null,3,null,3] is not:
//	   1
//	  / \
//	 2   2
//	 \   \
//   3    3

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func symmetricTree(l *bTreeNode, r *bTreeNode) bool {

	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.v != r.v {
		return false
	}
	if !symmetricTree(l.l, r.r) {
		return false
	}
	if !symmetricTree(r.l, l.r) {
		return false
	}

	return true
}

func main()  {
	root := &bTreeNode{2, nil, nil}
	root.l = &bTreeNode{1, nil, nil}
	root.r = &bTreeNode{1, nil, nil}
	root.r.l = &bTreeNode{4, nil, nil}
	root.r.r = &bTreeNode{3, nil, nil}
	root.l.l = &bTreeNode{3, nil, nil}
	root.l.r = &bTreeNode{4, nil, nil}
	fmt.Println(symmetricTree(root.l, root.r))
}
