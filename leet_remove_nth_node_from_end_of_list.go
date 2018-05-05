package main

import "fmt"

//Given a linked list, remove the n-th node from the end of list and return its head.
//
//Example:
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//
//Given n will always be valid.
//
//Follow up:
//
//Could you do this in one pass?


type Node struct {
	data int
	next *Node
}

func printList(head *Node)  {
	p := head
	i := 0
	for {
		if p == nil {
			return
		}
		fmt.Println("Node", i, " =", p.data)
		p = p.next
		i++
	}
}

func insertNode(d int, head **Node) {
	n := &Node{
		data: d,
		next: nil,
	}

	p := *head
	if p == nil {
		*head = n
		return
	}

	for {
		if p.next == nil {
			break
		}
		p = p.next
	}

	p.next = n
}

func removeNthNode(head *Node, N int) (h *Node)  {

	// 用两个指针start end指向一个长度为N的区间
	// 一开始start end都指向头部
	start, end := head, head
	// end先向前走N步，题目已假设 n不会超出范围
	for ;N>0;N-- {
		end = end.next
	}

	// 此时如果start为nil，说明list长度正好是N，那么要删除的正是head
	if end == nil {
		h = head.next
		return
	}

	h = head

	// 不是nil,说明长度大于N，此时end要继续走，直到走到nil，start跟着走，保持两者间隔N，则start为要删除的节点
	for ; end.next != nil; end = end.next {
		start = start.next
	}
	// start的next是要删除的节点
	start.next = start.next.next

	return
}

func main()  {
	var head *Node
	insertNode(0, &head)
	insertNode(10, &head)
	insertNode(3, &head)
	insertNode(230, &head)
	insertNode(11, &head)
	printList(head)

	head = removeNthNode(head, 1)
	printList(head)
}
