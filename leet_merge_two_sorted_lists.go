package main

import "fmt"

//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

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

func mergeSortedLists(l1 *Node, l2 *Node) (l *Node) {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 选头
	iter1 := l1
	iter2 := l2
	if l1.data > l2.data {
		l = l2
		iter2 = iter2.next
	} else {
		l = l1
		iter1 = iter1.next
	}
	iter3 := l

	for ;iter1 != nil && iter2 != nil ;  {
		if iter1.data > iter2.data {
			iter3.next = iter2
			iter2 = iter2.next
		} else {
			iter3.next = iter1
			iter1 = iter1.next
		}
		iter3 = iter3.next
	}

	for ;iter1 != nil ;  {
		iter3.next = iter1
		iter1 = iter1.next
		iter3 = iter3.next
	}
	for ;iter2 != nil ;  {
		iter3.next = iter2
		iter2 = iter2.next
		iter3 = iter3.next
	}


	return
}

func main() {

	var l1 *Node
	insertNode(0, &l1)
	insertNode(1, &l1)
	insertNode(3, &l1)

	var l2 *Node
	insertNode(1, &l2)
	insertNode(2, &l2)
	insertNode(6, &l2)

	l3 := mergeSortedLists(l1, l2)
	printList(l3)
}