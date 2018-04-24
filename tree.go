package main

import "fmt"

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

// 根据前序和中序遍历重建二叉树
func buildBtreeMiddle(pre []int, mid []int, node **bTreeNode, s string) *bTreeNode {

	if len(pre) == 0 || len(mid) == 0 {
		return nil
	}

	//fmt.Println("buildBtreeMiddle", pre, mid, s)

	// 从前序中找出根节点
	rootv := pre[0]

	// 从中序中找到根节点位置
	index := 0
	found := false
	for i,v := range mid {
		if v == rootv {
			index = i
			found = true
		}
	}
	if found {

		left := mid[:index]
		right := mid[index+1:]

		// 构造一个节点
		n := *node
		if n == nil {
			n = &bTreeNode{}
		}
		n.v = rootv

		hasleft := index > 0
		hasright := index < len(mid) - 1

		// 递归
		if hasleft {
			n.l = buildBtreeMiddle(pre[1:index+1], left, &(n.l), "l")
		}

		if hasright {
			n.r = buildBtreeMiddle(pre[index+1:], right, &(n.r), "r")
		}

		return n
	}

	return  nil
}

func main() {

	pre := []int{3,5,7,8,9,6,1,4}
	mid := []int{7,5,9,8,3,1,6,4}
	tree := &bTreeNode{}

	buildBtreeMiddle(pre, mid, &tree, "root")
	printBtreeMiddle(tree)
}
