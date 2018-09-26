package main

import "fmt"

//Given a binary tree, flatten it to a linked list in-place.
//
//For example, given the following tree:
//
//	   1
//	  / \
//	 2   5
//	/ \   \
// 3   4   6
//The flattened tree should look like:
//
//	1
//	 \
//	  2
//	   \
//	    3
//	     \
//	      4
//	       \
//	        5
//	         \
//	          6

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

var prenode *bTreeNode
func flattenBTreeToLinkedList(root *bTreeNode) {
	// 使用 右-左-中 的后序遍历
	// 有：前一个遍历到的节点  是后一个遍历到的节点的右孩子
	if root == nil {
		return
	}
	flattenBTreeToLinkedList(root.r)
	flattenBTreeToLinkedList(root.l)
	root.l = nil
	root.r = prenode
	prenode = root
	fmt.Printf("%d<-", root.v)
}

func main()  {
	root := &bTreeNode{1, nil, nil}
	root.l = &bTreeNode{2, nil, nil}
	root.r = &bTreeNode{5, nil, nil}
	root.l.l = &bTreeNode{3, nil, nil}
	root.l.r = &bTreeNode{4, nil, nil}
	root.r.r = &bTreeNode{6, nil, nil}
	flattenBTreeToLinkedList(root)
}
