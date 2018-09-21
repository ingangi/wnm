package main

import "fmt"

//Given preorder and inorder traversal of a tree, construct the binary tree.
//
//Note:
//You may assume that duplicates do not exist in the tree.
//
//For example, given
//
//preorder = [3,9,20,15,7]
//inorder = [9,3,15,20,7]
//Return the following binary tree:
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7


type bTreeNode struct {
	v int
	l *bTreeNode
	r *bTreeNode
}

func printBtreeMiddle(node *bTreeNode)  {
	if node == nil {
		return
	}
	printBtreeMiddle(node.l)
	fmt.Println(node.v)
	printBtreeMiddle(node.r)
}
func printBtreePre(node *bTreeNode)  {
	if node == nil {
		return
	}
	fmt.Println(node.v)
	printBtreePre(node.l)
	printBtreePre(node.r)
}

func constructBTreeFromPreAndInOrder(preorder []int, inorder []int) *bTreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	root := &bTreeNode{preorder[0],nil,nil}

	//找到中序遍历的根节点位置
	index := -1
	for i,v := range inorder {
		if v == root.v {
			index = i
			break
		}
	}

	if index == -1 {
		return nil
	}

	// index左右两边  分别为root左右子数的 前序和中序
	if index > 0 {
		root.l = constructBTreeFromPreAndInOrder(preorder[1:index+1], inorder[:index])
	}

	if index < len(preorder)-1 {
		root.r = constructBTreeFromPreAndInOrder(preorder[index+1:], inorder[index+1:])
	}

	return root
}

func main()  {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	r := constructBTreeFromPreAndInOrder(preorder, inorder)
	printBtreeMiddle(r)
	fmt.Println("======")
	printBtreePre(r)
}