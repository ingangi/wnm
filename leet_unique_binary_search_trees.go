package main

import "fmt"

//Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
//
//Example:
//
//Input: 3
//Output: 5
//Explanation:
//Given n = 3, there are a total of 5 unique BST's:
//
//	1         3     3      2      1
//	\       /     /      / \      \
//	3     2     1      1   3      2
//	/     /       \                 \
//	2     1         2                 3

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

func generateTrees(n int) []*bTreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(left int, right int) []*bTreeNode {
	result := []*bTreeNode{}
	if left > right {
		result = append(result, nil)
		return result
	}

	for i := left; i <= right; i++ {
		leftChild := helper(left, i-1)
		rightChild := helper(i+1, right)
		for _, lnode := range  leftChild {
			for _, rnode := range rightChild {
				newnode := &bTreeNode{i,nil,nil}
				newnode.l = lnode
				newnode.r = rnode
				result = append(result, newnode)
			}
		}
	}

	return result
}

func main()  {
	r := generateTrees(3)
	fmt.Println("total", len(r))
	for _, n := range r {
		fmt.Println("+++++++++++++++")
		printBtreeMiddle(n)
		fmt.Println("+++++++++++++++")
	}
}