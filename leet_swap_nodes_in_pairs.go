package main

import (
	"fmt"
)

//Given a linked list, swap every two adjacent nodes and return its head.
//
//Example:
//
//Given 1->2->3->4, you should return the list as 2->1->4->3.
//Note:
//
//Your algorithm should use only constant extra space.
//You may not modify the values in the list's nodes, only nodes itself may be changed.

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

func swapNodes(l *Node) (res *Node) {

	if l != nil && l.next != nil {
		res = l.next
		p1 := l
		p2 := l.next
		p1.next = p2.next
		p2.next = p1
		p1.next = swapNodes(p1.next)
	} else {
		return l
	}

	return
}

func main() {

	var l1 *Node
	insertNode(0, &l1)
	insertNode(1, &l1)
	insertNode(3, &l1)
	insertNode(4, &l1)
	insertNode(5, &l1)
	insertNode(6, &l1)
	//printList(l1)

	printList(swapNodes(l1))
}