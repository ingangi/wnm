package main

import (
	"fmt"
)

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

func rmNodeByValue(d int, head **Node) {
	if head == nil || *head == nil {
		return
	}

	h := *head
	if h.data == d {
		*head = h.next
		return
	} else {
		for {
			if h.next == nil {
				break
			}

			if h.next.data == d {
				h.next = h.next.next
				break
			}
			h = h.next
		}
	}
}

func main()  {
	var head *Node
	insertNode(0, &head)
	insertNode(10, &head)
	insertNode(3, &head)
	insertNode(230, &head)
	insertNode(11, &head)
	printList(head)

	fmt.Println("======================")
	rmNodeByValue(12, &head)
	printList(head)
}
