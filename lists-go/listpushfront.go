package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNodeL := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNodeL
		l.Tail = newNodeL
	} else {
		newNodeL.Next = l.Head
		l.Head = newNodeL
	}
}

func main() {
	link := &List{}
	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")
	current := link.Head
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
	fmt.Println()
}
