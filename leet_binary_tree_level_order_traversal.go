package main

import "fmt"

//Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//3
/// \
//9  20
///  \
//15   7
//return its level order traversal as:
//[
//[3],
//[9,20],
//[15,7]
//]

type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func btreeLevelOrderTraversal(nodes []*bTreeNode)  {
	if len(nodes) == 0 {
		return
	}

	var values []int
	var nextlev []*bTreeNode
	for _, node := range nodes {
		if node == nil {
			continue
		}

		values = append(values, node.v)
		if node.l != nil {
			nextlev = append(nextlev, node.l)
		}
		if node.r != nil {
			nextlev = append(nextlev, node.r)
		}
	}

	fmt.Println(values)
	btreeLevelOrderTraversal(nextlev)
}

func main()  {
	root := &bTreeNode{2, nil, nil}
	root.l = &bTreeNode{1, nil, nil}
	root.r = &bTreeNode{1, nil, nil}
	root.r.l = &bTreeNode{4, nil, nil}
	root.r.r = &bTreeNode{3, nil, nil}
	root.l.l = &bTreeNode{3, nil, nil}
	root.l.r = &bTreeNode{4, nil, nil}
	l := []*bTreeNode{root}
	btreeLevelOrderTraversal(l)
}