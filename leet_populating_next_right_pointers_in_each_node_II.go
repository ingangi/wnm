package main

import "fmt"

//Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
//
//Initially, all next pointers are set to NULL.
//
//Note:
//
//You may only use constant extra space.
//Recursive approach is fine, implicit stack space does not count as extra space for this problem.
//Example:
//
//Given the following binary tree,
//
//	      1
//	    /  \
//	   2    3
//	  / \    \
//	 4   5    7
//After calling your function, the tree should look like:
//
//	        1 -> NULL
//	      /  \
//	     2 -> 3 -> NULL
//     / \    \
//	 4-> 5 -> 7 -> NULL


type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
	n *bTreeNode	//pointer to next right node
}

func populatingNextRightPointerInEachNodeII(root *bTreeNode)  {
	head := &bTreeNode{0, nil, nil, nil}  //会指向每一层最左边的节点
	pre := head
	for root != nil {
		if root.l != nil {
			pre.n = root.l
			pre = pre.n
			fmt.Printf("%d->", pre.v)
		}

		if root.r != nil {
			pre.n = root.r
			pre = pre.n
			fmt.Printf("%d->", pre.v)
		}

		root = root.n
		if root == nil {  //一层结束了
			pre = head
			root = head.n
			head.n = nil
		}
	}
}

func main()  {
	root := &bTreeNode{1, nil, nil, nil}
	root.l = &bTreeNode{2, nil, nil, nil}
	root.r = &bTreeNode{3, nil, nil, nil}
	root.l.l = &bTreeNode{4, nil, nil, nil}
	root.l.r = &bTreeNode{5, nil, nil, nil}
	root.r.r = &bTreeNode{7, nil, nil, nil}
	populatingNextRightPointerInEachNodeII(root)
}