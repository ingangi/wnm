package main

import "fmt"

//
//Given inorder and postorder traversal of a tree, construct the binary tree.
//
//Note:
//You may assume that duplicates do not exist in the tree.
//
//For example, given
//
//inorder = [9,3,15,20,7]
//postorder = [9,15,7,20,3]
//Return the following binary tree:
//
//	 3
//	/ \
// 9  20
//	  /  \
//	15   7


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
func printBtreePost(node *bTreeNode)  {
	if node == nil {
		return
	}
	printBtreePost(node.l)
	printBtreePost(node.r)
	fmt.Println(node.v)
}

func constructBTreeFromInAndPostOrder(inorder []int, postorder []int) *bTreeNode {
	l := len(postorder)
	if l != len(inorder) || l == 0 {
		return nil
	}

	root := &bTreeNode{postorder[l-1],nil,nil}

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
		root.l = constructBTreeFromInAndPostOrder(inorder[:index], postorder[:index])
	}

	if index < l-1 {
		root.r = constructBTreeFromInAndPostOrder(inorder[index+1:], postorder[index:l-1])
	}

	return root
}

func main()  {
	postorder := []int{9,15,7,20,3}
	inorder := []int{9,3,15,20,7}
	r := constructBTreeFromInAndPostOrder(inorder, postorder)
	printBtreeMiddle(r)
	fmt.Println("======")
	printBtreePost(r)
}