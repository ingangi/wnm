package main

import "fmt"

//Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
//
//Example 1:
//
//Input: 1->2->3->3->4->4->5
//Output: 1->2->5
//Example 2:
//
//Input: 1->1->1->2->3
//Output: 2->3


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

func removeDuplicatesFromSortedList2(list **Node)  {

	head := *list

	// 在头前面加一个头  这样方便处理第一个就是重复的情况
	pre := &Node{
		data: -999,		//我们假设这个值不会在链表中用到
		next: head,
	}
	cur := head

	for ; cur != nil && cur.next != nil; {
		if cur.next.data == cur.data {

			v := cur.data
			for {
				cur = cur.next
				if cur == nil || v != cur.data {
					break
				}
			}

			pre.next = cur
		} else {
			pre = pre.next
			cur = cur.next
		}
	}

	*list = pre
}

func main() {

	var l1 *Node
	insertNode(0, &l1)
	insertNode(0, &l1)
	insertNode(2, &l1)
	insertNode(2, &l1)
	insertNode(5, &l1)
	insertNode(6, &l1)
	insertNode(6, &l1)
	removeDuplicatesFromSortedList2(&l1)
	printList(l1)
}