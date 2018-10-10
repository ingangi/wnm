package main

import "fmt"

//Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
//
//An example is the root-to-leaf path 1->2->3 which represents the number 123.
//
//Find the total sum of all root-to-leaf numbers.
//
//Note: A leaf is a node with no children.
//
//
//Input: [4,9,0,5,1]
//	   4
//	  / \
//	 9   0
//	/ \
// 5   1
//Output: 1026
//Explanation:
//The root-to-leaf path 4->9->5 represents the number 495.
//The root-to-leaf path 4->9->1 represents the number 491.
//The root-to-leaf path 4->0 represents the number 40.
//Therefore, sum = 495 + 491 + 40 = 1026.


type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

var sum int
func sumRoot2LeafNumbers(root *bTreeNode, pre int) {
	if root == nil {
		return
	}
	pre = pre*10 + root.v
	if root.r == nil && root.l == nil {
		sum += pre
		return
	}
	if root.l != nil {
		sumRoot2LeafNumbers(root.l, pre)
	}
	if root.r != nil {
		sumRoot2LeafNumbers(root.r, pre)
	}
}

func main()  {
	root := &bTreeNode{4, nil, nil}
	root.l = &bTreeNode{9, nil, nil}
	root.r = &bTreeNode{0, nil, nil}
	root.l.l = &bTreeNode{5, nil, nil}
	root.l.r = &bTreeNode{1, nil, nil}
	sumRoot2LeafNumbers(root, 0)
	fmt.Println(sum)
}

