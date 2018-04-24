package main

import "fmt"
//You are given two non-empty linked lists representing two non-negative integers.
//The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

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

func addTwoNum(l1 *Node, l2 *Node) (result *Node)  {
	if l1 == nil && l2 == nil {
		insertNode(0, &result)
		return
	}

	iter1 := l1
	iter2 := l2

	c := 0	//进位
	for {
		if iter1 == nil && iter2 == nil {
			if c > 0 {
				// 进位不能丢  比如999+999=1998 会多一个节点
				insertNode(c, &result)
			}
			break
		}
		d1, d2 := 0, 0
		if iter1 != nil {
			d1 = iter1.data
			iter1 = iter1.next
		}
		if iter2 != nil {
			d2 = iter2.data
			iter2 = iter2.next
		}
		d := c + d1 + d2
		c = d / 10
		r := d % 10
		insertNode(r, &result)
	}

	return
}

func main() {

	var l1 *Node
	insertNode(9, &l1)
	insertNode(9, &l1)
	insertNode(9, &l1)

	var l2 *Node
	insertNode(9, &l2)
	insertNode(9, &l2)
	insertNode(9, &l2)

	res := addTwoNum(l1, l2)
	printList(res)
}