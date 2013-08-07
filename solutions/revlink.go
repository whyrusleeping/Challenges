package main

import (
	"fmt"
)

type linkNode struct {
	val int
	next *linkNode
}

func (l *linkNode) AddToEnd(n *linkNode) {
	if l.next == nil {
		l.next = n
	} else {
		cur := l.next
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = n
	}
}

func NewNode(n int) *linkNode {
	nn := new(linkNode)
	nn.val = n
	nn.next = nil
	return nn
}

func reverseList(a **linkNode) {
	cur := *a
	var prev *linkNode
	var temp *linkNode
	for cur != nil {
		temp = cur.next
		cur.next = prev
		prev,cur = cur, temp
	}
	*a = prev
}

func printList(n *linkNode) {
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
}

func main() {
	a := NewNode(1)
	a.AddToEnd(NewNode(2))
	a.AddToEnd(NewNode(3))
	a.AddToEnd(NewNode(4))
	a.AddToEnd(NewNode(5))
	a.AddToEnd(NewNode(6))
	a.AddToEnd(NewNode(7))
	a.AddToEnd(NewNode(8))
	reverseList(&a)
	printList(a)
}
