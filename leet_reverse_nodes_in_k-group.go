package main

import "fmt"

//Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//Example:
//
//Given this linked list: 1->2->3->4->5
//
//For k = 2, you should return: 2->1->4->3->5
//
//For k = 3, you should return: 3->2->1->4->5

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

func reverse(head *Node) (rev *Node) {
	rev = nil
	iter := head
	for ;iter != nil ;  {
		tmp := iter.next
		iter.next = rev
		rev = iter
		iter = tmp
	}
	return
}

func reverseK(head *Node, k int) (rev *Node) {
	rev = nil
	tail := rev
	iter := head
	i := k
	for ;iter != nil;  {
		tmp := iter.next
		iter.next = rev
		if rev == nil {
			tail = iter
		}
		rev = iter
		iter = tmp
		i--

		if i == 0 && iter != nil {
			tail.next = reverseK(iter, k)
			return
		}
	}

	if i > 0 {
		rev = reverse(rev)
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

	printList(reverseK(l1, 4))

}