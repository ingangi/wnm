package main

import "fmt"

//Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
//Example 1:
//
//Input: 1->2->3->4->5->NULL, k = 2
//Output: 4->5->1->2->3->NULL
//Explanation:
//rotate 1 steps to the right: 5->1->2->3->4->NULL
//rotate 2 steps to the right: 4->5->1->2->3->NULL


type Node struct {
	data int
	next *Node
}

func printList(head *Node)  {
	fmt.Println("+++++++++++++++++")
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

func rotateList(l *Node, k int) *Node {

	// 先遍历一遍
	// 1. 获取长度
	// 2. 将链表变成一个环
	iter := l
	len := 0
	if iter == nil {
		return nil
	}

	for {
		len++
		if iter.next == nil {
			iter.next = l
			break
		} else {
			iter = iter.next
		}
	}

	walk := len - k % len
	// iter 再往前走 len - k步就是新链表的尾巴
	for walk>0 {
		iter = iter.next
		walk--
	}

	r := iter.next
	iter.next = nil
	return r
}

func main() {

	var l1 *Node
	insertNode(1, &l1)
	insertNode(2, &l1)
	insertNode(3, &l1)
	insertNode(4, &l1)
	insertNode(5, &l1)
	printList(l1)

	l := rotateList(l1, 23)
	printList(l)
}