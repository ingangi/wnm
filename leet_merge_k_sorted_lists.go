package main

import "fmt"

//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
//Example:
//
//Input:
//[
//1->4->5,
//1->3->4,
//2->6
//]
//Output: 1->1->2->3->4->4->5->6


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

func mergeKSortedList(ls []*Node) (l *Node) {
	length := len(ls)
	if length == 0 {
		return
	}

	if length == 1 {
		return ls[0]
	} else if length == 2 {
		return mergeSortedLists(ls[0], ls[1])
	} else {
		// 分治法
		return mergeSortedLists(mergeKSortedList(ls[:length/2]), mergeKSortedList(ls[length/2:]))
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

	var l3 *Node
	insertNode(11, &l3)
	insertNode(21, &l3)
	insertNode(61, &l3)

	var l4 *Node
	insertNode(12, &l4)
	insertNode(51, &l4)
	insertNode(64, &l4)

	var l5 *Node
	insertNode(6, &l5)
	insertNode(8, &l5)
	insertNode(46, &l5)

	ls := []*Node{l1,l2,l3,l4,l5}

	l := mergeKSortedList(ls)
	printList(l)
}