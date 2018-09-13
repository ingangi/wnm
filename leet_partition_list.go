package main

import "fmt"

//
//Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
//
//You should preserve the original relative order of the nodes in each of the two partitions.
//
//Example:
//
//Input: head = 1->4->3->2->5->2, x = 3
//Output: 1->2->2->4->3->5

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

func partitionList(list **Node, x int) {
	// 思路是遍历一遍list，生成2个链表，最后将链表连起来
	var less *Node
	var lessT *Node		//tail
	var greater *Node
	var greaterT *Node		//tail

	iter := *list
	for {
		if iter == nil {
			break
		}

		if iter.data < x {
			if less == nil {
				less = iter
			} else {
				lessT.next = iter
			}
			lessT = iter
		} else {
			if greater == nil {
				greater = iter
			} else {
				greaterT.next = iter
			}
			greaterT = iter
		}

		iter = iter.next
	}

	if less == nil {
		greaterT.next = nil
		*list = greater
	} else if greater == nil {
		lessT.next = nil
		*list = less
	} else {
		// 拼接
		lessT.next = greater
		greaterT.next = nil
		*list = less
	}
}

func main()  {
	var l1 *Node
	insertNode(1, &l1)
	insertNode(4, &l1)
	insertNode(3, &l1)
	insertNode(2, &l1)
	insertNode(5, &l1)
	insertNode(2, &l1)
	partitionList(&l1, 3)
	printList(l1)
}